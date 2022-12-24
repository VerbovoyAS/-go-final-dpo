package service

import (
	st "go-final-dpo/src/structure"
	"go-final-dpo/src/validate"
	"strings"
)

const (
	EMAIL_COUNTRY       = 0
	EMAIL_PROVIDER      = 1
	EMAIL_DELIVERY_TIME = 2
	EMAIL_LENGTH        = 3
)

func ParsingEmail(d *st.Data, content string) {
	rows := strings.Split(content, "\n")

	for _, row := range rows {
		if row == "" {
			continue
		}

		str := strings.Split(row, ";")
		var voiceCall st.EmailData

		if validate.Email(str, EMAIL_LENGTH, EMAIL_COUNTRY, EMAIL_PROVIDER) {

			deliveryTime, err := convertToInt(str[EMAIL_DELIVERY_TIME])

			if err != nil {
				continue
			}

			voiceCall.Country = str[EMAIL_COUNTRY]
			voiceCall.Provider = str[EMAIL_PROVIDER]
			voiceCall.DeliveryTime = deliveryTime

			d.Email = append(d.Email, voiceCall)
		}

	}

	return
}
