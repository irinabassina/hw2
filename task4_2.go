// Задание 4.1. Пакет ioutil
//Перепишите задачу 2, используя пакет ioutil.

// Задание 2. Интерфейс io.Reader
// Напишите программу, которая читает и выводит в консоль строки из файла, созданного в предыдущей практике, без использования ioutil.
// Если файл отсутствует или пуст, выведите в консоль соответствующее сообщение.
// Рекомендация:
// Для получения размера файла воспользуйтесь методом Stat(), который возвращает информацию о файле и ошибку.

package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	fileinfo, err := os.Stat("file.txt")
	if err != nil || fileinfo.Size() == 0 {
		fmt.Println("Файл отсутствует или пуст,", err)
		return
	}

	dat, err := ioutil.ReadFile("file.txt")
	if err != nil {
		fmt.Println("Ошибка чтения файла", err)
		return
	}
	fmt.Print(string(dat))
}
