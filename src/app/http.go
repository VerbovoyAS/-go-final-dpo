package app

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Request Отправляет запрос и возвращает тело ответа
func Request(url string) (body []byte, err error) {
	res, err := http.Get(url)

	if err != nil {
		return
	}

	if res.StatusCode != 200 {
		err = fmt.Errorf("status code not 200. Current code: %d", res.StatusCode)
	}

	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	return
}
