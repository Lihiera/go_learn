package main

import "fmt"

func max(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, fmt.Errorf("no value")
	}
	max := vals[0]
	for _, val := range vals {
		if val > max {
			max = val
		}
	}
	return max, nil
}

func min(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, fmt.Errorf("no value")
	}
	min := vals[0]
	for _, val := range vals {
		if val < min {
			min = val
		}
	}
	return min, nil
}

func main() {
	fmt.Println(myJoin(",:", "a", "b", "c"))
	fmt.Println(myJoin(""))
}

func myJoin(sep string, vals ...string) string {
	result := ""
	for i, val := range vals {
		if i == 0 {
			result += val
		} else {
			result += sep + val
		}
	}
	return result
}
