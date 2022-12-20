package service

import (
	st "go-final-dpo/src/structure"
	"go-final-dpo/src/validate"
	"strings"
)

func ParsingSms(d *st.Data, content string) {
	rows := strings.Split(content, "\n")

	for _, row := range rows {
		if row == "" {
			continue
		}

		str := strings.Split(row, ";")
		var sms st.SMSData

		if validate.Sms(str) {
			sms.Country = str[0]
			sms.Bandwidth = str[1]
			sms.ResponseTime = str[2]
			sms.Provider = str[3]

			d.Sms = append(d.Sms, sms)
		}

	}

	return
}
