package service

import (
	"bufio"
	"fmt"
	st "go-final-dpo/src/structure"
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

func CreateFile(d *st.Data) {
	file, err := os.Create("data/sms.csv")

	if err := os.Chmod("data/sms.csv", 0664); err != nil {
		fmt.Println(err)
	}

	w := bufio.NewWriter(file)
	if err != nil {
		fmt.Println(err)
		panic("Ошибка открытия файла")
	}
	defer file.Close()

	for _, str := range d.Sms {
		s := str.ToString()
		w.WriteString(s)
	}

	w.Flush()
}
