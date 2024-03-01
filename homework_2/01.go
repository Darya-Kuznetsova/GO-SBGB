package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	// Напишите программу, которая на вход получала бы строку, введённую пользователем, а в файл писала № строки, дату и сообщение
	// в формате: 2020-02-10 15:00:00 продам гараж.
	// При вводе слова exit программа завершает работу.

	file, err := os.Create("log.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	t := time.Now()
	now := t.Format("2006-01-02 03:04:05")
	for {
		fmt.Println("Введите текст:")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		if scanner.Text() == "exit" {
			return
		} else {
			file.WriteString(now + " " + scanner.Text())
			file.WriteString("\n")
		}
	}

}
