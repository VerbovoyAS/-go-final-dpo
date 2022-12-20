package app

import (
	"go-final-dpo/src/enum"
	"go-final-dpo/src/service"
	"strings"
)

// InitEnum Инициализирует все данные Enum
func InitEnum() {
	initCountryCode()
	initProvider()
}

// initCountryCode Инициализирует данные по кодам стран из файла
func initCountryCode() {
	content := service.GetContent(CountryCodePath())
	parsingCountryCode(content)
}

// initProvider Инициализирует данные по провайдерам из файла
func initProvider() {
	content := service.GetContent(ProviderPath())
	parsingProvider(content)
}

// Собирает map по провайдерам
func parsingProvider(content string) {
	rows := strings.Split(content, "\n")
	provider := enum.Provider

	for _, row := range rows {
		if row == "" {
			continue
		}
		provider[row] = row
	}

	return
}

// Собирает map по кодам стран
func parsingCountryCode(content string) {
	rows := strings.Split(content, "\n")
	code := enum.CountryCode

	for _, row := range rows {
		if row == "" {
			continue
		}

		str := strings.Split(row, ",")
		code[str[1]] = str[0]
	}

	return
}
