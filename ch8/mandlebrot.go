package main

import (
	"fmt"
	"image"
	"image/color"
	"math/cmplx"
	"runtime"
	"sync"
	"time"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 1024, 1024
	routs                  = 10
)

func main() {
	var total time.Duration
	for i := 0; i < 10; i++ {
		total += perform1()
	}
	fmt.Printf("only one goroutine:%.2f\n", total.Seconds()/10)
	total = 0
	for i := 0; i < 10; i++ {
		total += perform2()
	}
	fmt.Printf("%d goroutines:%.2f\n", routs, total.Seconds()/10)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func perform1() time.Duration {
	start := time.Now()
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	return time.Since(start)
}

func perform2() time.Duration {
	start := time.Now()
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	numWorkers := runtime.NumCPU()
	//fmt.Println(numWorkers)
	numWorkers = 2
	var wg sync.WaitGroup
	taskChan := make(chan [2]int)
	for i := 0; i < numWorkers; i++ {
		go func() {
			for pixel := range taskChan {
				px, py := pixel[0], pixel[1]
				x := float64(px)/width*(xmax-xmin) + xmin
				y := float64(py)/height*(ymax-ymin) + ymin
				z := complex(x, y)
				img.Set(px, py, mandelbrot(z))
				wg.Done()
			}
		}()
	}
	for py := 0; py < height; py++ {
		for px := 0; px < width; px++ {
			wg.Add(1)
			taskChan <- [2]int{px, py}
		}
	}
	close(taskChan)
	wg.Wait()
	return time.Since(start)
}
