package app

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

// Переменные хранения фалов для обработки
var (
	folderFiles     string
	pathProvider    string
	pathCountryCode string
)

// Переменные хранения фалов Skillbox
var (
	folderSkillbox  string
	pathSkillboxSms string
)

// Переменные хранения обработанных фалов
var (
	folderData  string
	pathDataSms string
)

// Переменные API MMS
var (
	server  string
	portMms string
	urlMms  string
)

// InitEnv Иницизизация переменных ENV
func InitEnv() {
	if err := godotenv.Load(); err != nil {
		panic("No .env file found")
	}

	folderFiles = getEnv("FOLDER_FILES")
	pathProvider = getEnv("PATH_PROVIDER")
	pathCountryCode = getEnv("PATH_COUNTRY_CODE")

	folderSkillbox = getEnv("FOLDER_SKILLBOX")
	pathSkillboxSms = getEnv("PATH_SKILLBOX_SMS")

	folderData = getEnv("FOLDER_DATA")
	pathDataSms = getEnv("PATH_DATA_SMS")

	server = getEnv("SERVER")
	portMms = getEnv("PORT_MMS")
	urlMms = getEnv("URL_MMS")
}

// Проверяет и возвращает ENV переменную
func getEnv(envName string) (env string) {

	env, exists := os.LookupEnv(envName)
	if !exists {
		panic(fmt.Sprintf("ENV NOR FOUND: %s", envName))
	}

	return
}

// ProviderPath Путь до файла разрешенных провайдеров
func ProviderPath() string {
	return folderFiles + pathProvider
}

// CountryCodePath Путь до файла кодов стран
func CountryCodePath() string {
	return folderFiles + pathCountryCode
}

// SkillboxSmsPath Путь до файла skillbox SMS данных
func SkillboxSmsPath() string {
	return folderSkillbox + pathSkillboxSms
}

// DataSmsPath Путь до файла обработанных данных SMS
func DataSmsPath() string {
	return folderData + pathDataSms
}

// PathMms Ссылка API получения данных MMS
func PathMms() string {
	return fmt.Sprintf("%s:%s%s", server, portMms, urlMms)
}
