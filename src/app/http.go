package app

import (
	"fmt"
	"io"
	"net/http"
)

// Request Отправляет запрос и возвращает тело ответа
func Request(url string) (body []byte) {
	res, err := http.Get(url)

	if err != nil {
		return
	}

	if res.StatusCode != 200 {
		fmt.Printf("status code not 200. Current code: %d", res.StatusCode)
		return
	}

	body, err = io.ReadAll(res.Body)
	if err != nil {
		return
	}
	return
}
