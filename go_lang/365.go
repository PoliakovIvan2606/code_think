package main

import "fmt"

func main() {
	fmt.Println(canMeasureWater(1, 2, 3))
}

func canMeasureWater(x int, y int, target int) int {
    // bal := [2]int{}
	s, b := reverse(x, y)
	fmt.Println(s, b)
	return b - (s - (b - s))	
}

func reverse(x int, y int) (int, int) {
	if x > y {
		y, x = x, y
	}
	return x, y
}