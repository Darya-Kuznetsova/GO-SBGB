package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var str string
	for {
		fmt.Println("Введите число:")
		fmt.Scan(&str)
		if strings.ToLower(str) == "стоп" {
			return
		} else {
			num, err := strconv.Atoi(str)
			if err != nil {
				fmt.Println("Введено не число", err)
			}
			v1 := Square(num)
			v2 := Product(v1)
			fmt.Println("Произведение", <-v2)
		}
	}

}

func Square(number int) chan int {
	fmt.Println("Число:", number)
	c1 := make(chan int)
	go func() {
		c1 <- number * number
	}()
	return c1
}

func Product(first chan int) chan int {
	square := <-first
	fmt.Println("Квадрат", square)
	c2 := make(chan int)
	go func() {
		c2 <- square * 2
	}()
	return c2
}
