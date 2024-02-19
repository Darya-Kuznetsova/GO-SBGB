package main

import "fmt"

func main() {
	arr := [8]int{0, 12, 9, 8, 7, 3, 2, 0}
	sortArr := bubbleSort(arr[:])
	fmt.Println(sortArr)
}

func bubbleSort(array []int) []int {
	for i := 0; i <= len(array); i++ {
		for j := 0; j < len(array)-1; j++ {
			/* len(array) - 1 т.к. array[len(array) - 2] сравнивается с array[len(array) - 1],
			а array[len(array) - 1] уже ни с чем не сравнивается */
			if array[j] > array[j+1] {
				memory := array[j]
				array[j] = array[j+1]
				array[j+1] = memory
			}
		}
	}
	return array

}
