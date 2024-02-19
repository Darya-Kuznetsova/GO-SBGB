package main

import "fmt"

func main() {
	arr1 := [5]int{7, 5, 9, 0, 3}
	arr2 := [4]int{8, 9, 12, 3}
	arrNew := newArray(arr1[:], arr2[:])
	fmt.Println(arrNew)
}

func newArray(arr1, arr2 []int) []int {
	var newArray []int
	for _, v1 := range arr1 {
		newArray = append(newArray, v1)
	}
	for _, v2 := range arr2 {
		newArray = append(newArray, v2)
	}
	return newArray
}
