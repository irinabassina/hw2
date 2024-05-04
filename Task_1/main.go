// Напишите программу, которая на вход получала бы строку, введённую пользователем, а в файл писала № строки, дату и сообщение в формате:
// 2020-02-10 15:00:00 продам гараж.
// При вводе слова exit программа завершает работу.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	f, err := os.Create("file.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

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
			date := time.Now().Format("2020-02-10 15:00:00")
			if _, err = f.WriteString(fmt.Sprintf("%d %s %s\n", i, date, text)); err != nil {
				panic(err)
			}

		}
		i++
	}
}
