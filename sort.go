package main

import "fmt"

func main() {
	var arr = []int8{2, 1, 2, 4, 10, 0}
	arr = bubbleSort(arr)
	for _, v := range arr {
		fmt.Printf("%d\n", v)
	}
}
