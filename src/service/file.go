package service

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// GetContent Открывает файл и возвращает его содержимое
func GetContent(fileName string) string {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	result, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Не смогли прочитать файл.", err)
		panic("Не смогли прочитать файл")
	}

	return string(result)
}

// CreateFile Создает файл
func CreateFile(content string, fileName string) {
	file, err := os.Create(fileName)

	if err := os.Chmod(fileName, 0664); err != nil {
		fmt.Println(err)
	}

	w := bufio.NewWriter(file)
	if err != nil {
		fmt.Println(err)
		panic("Ошибка открытия файла")
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	_, _ = w.WriteString(content)
	_ = w.Flush()
}
