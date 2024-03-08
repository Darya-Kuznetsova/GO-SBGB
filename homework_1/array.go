package main

import "fmt"

func main() {
	arr1 := [4]int{2, 3, 4, 5}
	arr2 := [5]int{6, 7, 8, 9, 10}
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
