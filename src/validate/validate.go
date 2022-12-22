package validate

import (
	"go-final-dpo/src/enum"
	st "go-final-dpo/src/structure"
)

// Sms Проверяет Валидность данных смс
func Sms(str []string) bool {
	return len(str) == 4 && checkValueMap(enum.CountryCode, str[0]) && checkValueMap(enum.Provider, str[3])
}

// Mms Проверяет Валидность данных MMS
func Mms(mmsData st.MMSData) bool {
	return checkValueMap(enum.CountryCode, mmsData.Country) && checkValueMap(enum.Provider, mmsData.Provider)
}

// Проверяет наличие ключа в Map
func checkValueMap(m map[string]string, str string) (ok bool) {
	_, ok = m[str]
	return
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
