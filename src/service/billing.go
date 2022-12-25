package service

import (
	st "go-final-dpo/src/structure"
	"strings"
)

const (
	B_CREATE_CUSTOMER = 0
	B_PURCHASE        = 1
	B_PAYOUT          = 2
	B_RECURRING       = 3
	B_FRAUD_CONTROL   = 4
	B_CHECKOUT_PAGE   = 5
	B_LENGTH          = 6
)

func ParsingBilling(d *st.Data, content string) {
	rows := strings.Split(content, "\n")

	for _, row := range rows {
		if row == "" {
			continue
		}

		b := []byte(row)
		mapBilling := make([]bool, B_LENGTH)
		length := len(b) - 1
		// Читаем бит с право на лево
		for i := length; i >= 0; i-- {
			// Если длина текущиго бита привышает длину billing - пропускаем
			if (len(b) - i) > len(mapBilling) {
				continue
			}
			// Собираем map по billing
			mapBilling[length-i] = (int(b[i]) & 1) == 1
		}

		var billingData st.BillingData
		billingData.CreateCustomer = mapBilling[B_CREATE_CUSTOMER]
		billingData.Purchase = mapBilling[B_PURCHASE]
		billingData.Payout = mapBilling[B_PAYOUT]
		billingData.Recurring = mapBilling[B_RECURRING]
		billingData.FraudControl = mapBilling[B_FRAUD_CONTROL]
		billingData.CheckoutPage = mapBilling[B_CHECKOUT_PAGE]

		d.Billing = append(d.Billing, billingData)

	}

	return
}
