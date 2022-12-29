package app

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

// Переменные хранения фалов для ENUM
var (
	folderFiles       string
	pathProvider      string
	pathCountryCode   string
	pathVoiceProvider string
	pathEmailProvider string
)

// Переменные хранения фалов SKILLBOX
var (
	folderSkillbox        string
	pathSkillboxSms       string
	pathSkillboxVoiceCall string
	pathSkillboxEmail     string
	pathSkillboxBilling   string
)

// Переменные API SKILLBOX
var (
	server      string
	portServer  string
	urlMms      string
	urlSupport  string
	urlIncident string
)

// Переменные хранения обработанных фалов
var (
	folderData    string
	pathDataSms   string
	pathDataVoice string
	pathDataEmail string
)

// Переменные MY SERVER
var (
	myServer     string
	myPortServer string
	myMainUrl    string
)

// InitEnv Инициализация переменных ENV
func InitEnv() {
	if err := godotenv.Load(); err != nil {
		panic("No .env file found")
	}

	folderFiles = getEnv("FOLDER_FILES")
	pathProvider = getEnv("PATH_PROVIDER")
	pathCountryCode = getEnv("PATH_COUNTRY_CODE")
	pathVoiceProvider = getEnv("PATH_VOICE_PROVIDER")
	pathEmailProvider = getEnv("PATH_EMAIL_PROVIDER")

	folderSkillbox = getEnv("FOLDER_SKILLBOX")
	pathSkillboxSms = getEnv("PATH_SKILLBOX_SMS")
	pathSkillboxVoiceCall = getEnv("PATH_SKILLBOX_VOICE_CALL")
	pathSkillboxEmail = getEnv("PATH_SKILLBOX_EMAIL")
	pathSkillboxBilling = getEnv("PATH_SKILLBOX_BILLING")

	folderData = getEnv("FOLDER_DATA")
	pathDataSms = getEnv("PATH_DATA_SMS")
	pathDataVoice = getEnv("PATH_DATA_VOICE")
	pathDataEmail = getEnv("PATH_DATA_EMAIL")

	server = getEnv("SERVER")
	portServer = getEnv("PORT_SERVER")
	urlMms = getEnv("URL_MMS")
	urlSupport = getEnv("URL_SUPPORT")
	urlIncident = getEnv("URL_INCIDENT")

	myServer = getEnv("MY_SERVER")
	myPortServer = getEnv("MY_PORT_SERVER")
	myMainUrl = getEnv("MY_MAIN_URL")
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

// VoiceProviderPath Путь до файла провайдеров голосовых звонков
func VoiceProviderPath() string {
	return folderFiles + pathVoiceProvider
}

// EmailProviderPath Путь до файла провайдеров E-mael
func EmailProviderPath() string {
	return folderFiles + pathEmailProvider
}

// SkillboxSmsPath Путь до файла skillbox SMS данных
func SkillboxSmsPath() string {
	return folderSkillbox + pathSkillboxSms
}

// DataSmsPath Путь до файла обработанных данных SMS
func DataSmsPath() string {
	return folderData + pathDataSms
}

// DataVoicePath Путь до файла обработанных данных Voice Call
func DataVoicePath() string {
	return folderData + pathDataVoice
}

// DataEmailPath Путь до файла обработанных данных E-mail
func DataEmailPath() string {
	return folderData + pathDataEmail
}

// PathMms Ссылка API получения данных MMS
func PathMms() string {
	return fmt.Sprintf("%s:%s%s", server, portServer, urlMms)
}

// PathSupport Ссылка API получения данных Support
func PathSupport() string {
	return fmt.Sprintf("%s:%s%s", server, portServer, urlSupport)
}

// PathIncident Ссылка API получения данных Incident
func PathIncident() string {
	return fmt.Sprintf("%s:%s%s", server, portServer, urlIncident)
}

// PathMyServer Host и port API получения данных о системе
func PathMyServer() string {
	return fmt.Sprintf("%s:%s", myServer, myPortServer)
}

// PathMyServerUrl Url API получения данных о системе
func PathMyServerUrl() string {
	return myMainUrl
}

// SkillboxVoiceCallPath Путь до файла skillbox VoiceCall данных
func SkillboxVoiceCallPath() string {
	return folderSkillbox + pathSkillboxVoiceCall
}

// SkillboxEmailPath Путь до файла skillbox email данных
func SkillboxEmailPath() string {
	return folderSkillbox + pathSkillboxEmail
}

// SkillboxBillingPath Путь до файла skillbox Billing данных
func SkillboxBillingPath() string {
	return folderSkillbox + pathSkillboxBilling
}
