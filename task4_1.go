// Задание 4.1. Пакет ioutil
//Перепишите задачу 1, используя пакет ioutil.

// Напишите программу, которая на вход получала бы строку, введённую пользователем, а в файл писала № строки, дату и сообщение в формате:
// 2020-02-10 15:00:00 продам гараж.
// При вводе слова exit программа завершает работу.

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func main() {
	var b bytes.Buffer

	i := 1
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Введите текст: ")
		if !scanner.Scan() {
			break
		}
		text := scanner.Text()

		if strings.ToLower(text) == "exit" || text == "" {
			fmt.Println("Завершение работы программы")
			break
		} else {
			date := time.Now().Format(time.DateTime)

			if _, err := b.WriteString(fmt.Sprintf("%d %s %s\n", i, date, text)); err != nil {
				panic(err)
			}

		}
		i++
	}

	fileName := "log.txt"
	if err := ioutil.WriteFile(fileName, b.Bytes(), 0666); err != nil {
		panic(err)
	}
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	resultBytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	fmt.Println("Сохраненный лог: ")
	fmt.Println(string(resultBytes))
}
