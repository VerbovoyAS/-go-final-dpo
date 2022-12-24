package validate

import (
	"go-final-dpo/src/enum"
	st "go-final-dpo/src/structure"
)

// Sms Проверяет Валидность данных смс
func Sms(str []string, length int, country int, provider int) bool {
	return len(str) == length && checkValueMap(enum.CountryCode, str[country]) && checkValueMap(enum.Provider, str[provider])
}

// Mms Проверяет валидность данных MMS
func Mms(mmsData st.MMSData) bool {
	return checkValueMap(enum.CountryCode, mmsData.Country) && checkValueMap(enum.Provider, mmsData.Provider)
}

// Проверяет наличие ключа в Map
func checkValueMap(m map[string]string, str string) (ok bool) {
	_, ok = m[str]
	return
}

// VoiceCall Проверяет валидность данных Voice Call
func VoiceCall(str []string, length int, country int, voiceProvider int) bool {
	return len(str) == length && checkValueMap(enum.CountryCode, str[country]) && checkValueMap(enum.VoiceProvider, str[voiceProvider])
}

// Email Проверяет валидность данных E-mail
func Email(str []string, length int, country int, emailProvider int) bool {
	return len(str) == length && checkValueMap(enum.CountryCode, str[country]) && checkValueMap(enum.EmailProvider, str[emailProvider])
}

// CheckValueString Проверяет значение в массиве
func CheckValueString(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
