// Задание 2. Интерфейс io.Reader
// Напишите программу, которая читает и выводит в консоль строки из файла, созданного в предыдущей практике, без использования ioutil.
// Если файл отсутствует или пуст, выведите в консоль соответствующее сообщение.
// Рекомендация:
// Для получения размера файла воспользуйтесь методом Stat(), который возвращает информацию о файле и ошибку.

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
			date := time.Now().Format(time.DateTime)
			if _, err = f.WriteString(fmt.Sprintf("%d %s %s\n", i, date, text)); err != nil {
				panic(err)
			}

		}
		i++
	}
	f, err = os.Open("file.txt")
	if err != nil {
		fmt.Println("Не смогли открыть файл,", err)
		return
	}
	defer f.Close()

	buffer := make([]byte, 155)
	_, err = f.Read(buffer)
	if err != nil {
		fmt.Println("Не смогли прочитать последовательность байт из файла,", err)
		return
	}
	fmt.Println(string(buffer))
}
