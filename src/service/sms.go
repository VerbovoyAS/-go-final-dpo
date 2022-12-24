package service

import (
	st "go-final-dpo/src/structure"
	"go-final-dpo/src/validate"
	"strings"
)

const (
	SMS_COUNTRY       = 0
	SMS_BAND_WIDTH    = 1
	SMS_RESPONSE_TIME = 2
	SMS_PROVIDER      = 3
	SMS_LENGTH        = 4
)

func ParsingSms(d *st.Data, content string) {
	rows := strings.Split(content, "\n")

	for _, row := range rows {
		if row == "" {
			continue
		}

		str := strings.Split(row, ";")
		var sms st.SMSData

		if validate.Sms(str, SMS_LENGTH, SMS_COUNTRY, SMS_PROVIDER) {
			sms.Country = str[SMS_COUNTRY]
			sms.Bandwidth = str[SMS_BAND_WIDTH]
			sms.ResponseTime = str[SMS_RESPONSE_TIME]
			sms.Provider = str[SMS_PROVIDER]

			d.Sms = append(d.Sms, sms)
		}

	}

	return
}
