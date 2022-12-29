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
	initVoiceProvider()
	initEmailProvider()
}

// initCountryCode Инициализирует данные по кодам стран из файла
func initCountryCode() {
	content := service.GetContent(CountryCodePath())
	collectMapCountryCode(content)
}

// initProvider Инициализирует данные по провайдерам из файла
func initProvider() {
	content := service.GetContent(ProviderPath())
	collectMap(content, enum.Provider)
}

// initVoiceProvider Инициализирует данные по провайдерам голосовых звонков из файла
func initVoiceProvider() {
	content := service.GetContent(VoiceProviderPath())
	collectMap(content, enum.VoiceProvider)
}

// initEmailProvider Инициализирует данные по провайдерам E-mail из файла
func initEmailProvider() {
	content := service.GetContent(EmailProviderPath())
	collectMap(content, enum.EmailProvider)
}

// collectMap Собирает map
func collectMap(content string, m map[string]string) {
	rows := strings.Split(content, "\n")

	for _, row := range rows {
		if row == "" {
			continue
		}
		m[row] = row
	}

	return
}

// collectMapCountryCode Собирает map по кодам стран
func collectMapCountryCode(content string) {
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
