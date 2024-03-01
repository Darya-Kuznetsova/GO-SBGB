package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	// Напишите программу, которая читает и выводит в консоль строки из файла, созданного в предыдущей практике, без использования ioutil. Если файл отсутствует или пуст, выведите в консоль соответствующее сообщение.
	// Рекомендация:
	// Для получения размера файла воспользуйтесь методом Stat(), который возвращает информацию о файле и ошибку.

	f, err := os.Open("log.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	stat, err := f.Stat()
	if err != nil {
		panic(err)
	}
	if stat.Size() == 0 {
		fmt.Println("The file is empty")
	} else {
		readSlice := make([]byte, stat.Size())
		if _, err := io.ReadFull(f, readSlice); err != nil {
			panic(err)
		}
		fmt.Printf("%s \n", readSlice)
	}

}
