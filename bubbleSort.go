package main

func bubbleSort(arr []int8) []int8 {
	count := len(arr)
	for i := 0; i < count-1; i++ {
		for j := 0; j < count-i-1; j++ {
			if arr[j] > arr[j+1] {
				var temp = arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
	return arr
}
