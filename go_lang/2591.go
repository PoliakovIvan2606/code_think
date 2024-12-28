package main

import "fmt"

func main() {
	n := 20
	count := float32(n / 7)
	long := float32(n) - float32(count * 7)
	fmt.Println((28 * count) + ((7 * (count - 1)) / 2) * count + (((count + 1) + long + count) / 2) * long)

}

func totalMoney(n int) int {
	count := float32(n / 7)
	long := float32(n) - float32(count * 7)
	return int((28 * count) + ((7 * (count - 1)) / 2) * count + (((count + 1) + long + count) / 2) * long)

}
