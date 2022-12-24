package service

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func GetContent(fileName string) string {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	result, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Не смогли прочитать файл.", err)
		panic("Не смогли прочитать файл")
	}

	return string(result)
}

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
	defer file.Close()

	w.WriteString(content)
	w.Flush()
}
