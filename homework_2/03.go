package main

import (
	"fmt"
	"os"
)

func main() {
	// Напишите программу, создающую текстовый файл только для чтения, и проверьте, что в него нельзя записать данные.
	// Рекомендация:
	// Для проверки создайте файл, установите режим только для чтения, закройте его, а затем, открыв, попытайтесь прочесть из него данные.
	f, err := os.Create("03.txt")
	if err != nil {
		fmt.Println("CAN'T CREATE", err)
	}
	defer f.Close()
	f.WriteString("test test test")
	readFile("03.txt")
	f.Chmod(0444)
	readFile("03.txt")

}
func readFile(file string) {
	f, err := os.Open(file)
	if err != nil {
		fmt.Println("CAN'T OPEN", err)
	}
	read, err := os.ReadFile(f.Name())
	if err != nil {
		fmt.Println("CAN'T READ", err)
	}
	fmt.Printf("%s \n", read)

}
