package st

type Data struct {
	Sms       []SMSData
	Mms       []MMSData
	VoiceCall []VoiceCallData
}

func (d *Data) VoiceSMSContent() (content string) {
	for _, v := range d.Sms {
		content += v.ToString()
	}
	return
}

func (d *Data) VoiceCallContent() (content string) {
	for _, v := range d.VoiceCall {
		content += v.ToString()
	}
	return
}
