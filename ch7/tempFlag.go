package main

import (
	"flag"
	"fmt"
)

type Celsius float64    // 摄氏温度
type Fahrenheit float64 // 华氏温度

const (
	AbsoluteZeroC Celsius = -273.15 // 绝对零度
	FreezingC     Celsius = 0       // 结冰点温度
	BoilingC      Celsius = 100     // 沸水温度
)

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }

type fahrenheitFlag struct{ Fahrenheit }

func (f *fahrenheitFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit) // no error check needed
	switch unit {
	case "F", "°F":
		f.Fahrenheit = Fahrenheit(value)
		return nil
	case "C", "°C":
		f.Fahrenheit = CToF(Celsius(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

func FahrenheitFlag(name string, value Fahrenheit, usage string) *Fahrenheit {
	f := &fahrenheitFlag{value}
	flag.CommandLine.Var(f, name, usage)
	return &f.Fahrenheit
}

var temp = FahrenheitFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
