package service

import "strconv"

// convertToInt Конвертирует из строки в int
func convertToInt(s string) (i int, err error) {
	i, err = strconv.Atoi(s)
	return
}

// convertToFloat Конвертирует из строки в float32
func convertToFloat(s string) (f float32, err error) {
	n, err := strconv.ParseFloat(s, 32)
	f = float32(n)
	return
}
