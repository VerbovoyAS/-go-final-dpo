package st

type Data struct {
	Sms       []SMSData
	Mms       []MMSData
	VoiceCall []VoiceCallData
	Email     []EmailData
	Billing   BillingData
	Support   []SupportData
	Incident  []IncidentData
}

// VoiceSMSContent Возвращает данные по SMS в виде строки
func (d *Data) VoiceSMSContent() (content string) {
	for _, v := range d.Sms {
		content += v.ToString()
	}
	return
}

// VoiceCallContent Возвращает данные по Voice Call в виде строки
func (d *Data) VoiceCallContent() (content string) {
	for _, v := range d.VoiceCall {
		content += v.ToString()
	}
	return
}

// EmailContent Возвращает данные по Email в виде строки
func (d *Data) EmailContent() (content string) {
	for _, v := range d.Email {
		content += v.ToString()
	}
	return
}
