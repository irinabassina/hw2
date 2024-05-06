// Задание 3. Уровни доступа

// Напишите программу, создающую текстовый файл только для чтения, и проверьте, что в него нельзя записать данные.
// Рекомендация:
// Для проверки создайте файл, установите режим только для чтения, закройте его, а затем, открыв, попытайтесь
// прочесть из него данные.

package main

import (
	"os"
)

func main() {
	fileName := "Test.txt"
	fileTest, err := os.Create(fileName) //Создаем файл
	if err != nil {
		os.Exit(1)
	}
	defer fileTest.Close()

	if err := os.Chmod(fileName, 0400); err != nil { //Изменяем право доступа
		os.Exit(2)
	}
	fileTest.WriteString("Записываем строку в файл")
	if fileTest, err = os.Open(fileName); err != nil {
		os.Exit(3)
	}

	if _, err = fileTest.WriteString("Записываем вторую строку в файл"); err != nil {
		os.Exit(5) //Добавили обработку ошибки для записи второй строки в файл
	}
}
