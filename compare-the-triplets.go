package main

import (
	"fmt"
)

func test(data []int) []int {
	result := []int{0, 0}
	var a = []int{
		data[0], data[1], data[2],
	}
	var b = []int{
		data[3], data[4], data[5],
	}

	for i := range a {
		if a[i] > b[i] {
			result[0] += 1
		} else if a[i] < b[i] {
			result[1] += 1
		}
	}

	return result

}

func main() {
	var a0, a1, a2, b0, b1, b2 int
	fmt.Scan(&a0, &a1, &a2)
	fmt.Scan(&b0, &b1, &b2)

	var data = []int{a0, a1, a2, b0, b1, b2}
	var result = test(data)
	fmt.Print(result[0])
	fmt.Print(" ")
	fmt.Print(result[1])

	//Enter your code here. Read input from STDIN. Print output to STDOUT
}
