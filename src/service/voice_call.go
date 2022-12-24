package service

import (
	st "go-final-dpo/src/structure"
	"go-final-dpo/src/validate"
	"strconv"
	"strings"
)

const (
	VC_COUNTRY              = 0
	VC_BANDWIDTH            = 1
	VC_RESPONSETIME         = 2
	VC_PROVIDER             = 3
	VC_CONNECTION_STABILITY = 4
	VC_TTFB                 = 5
	VC_VOICE_PURITY         = 6
	VC_MEDIAN_OF_CALLS_TIME = 7
	VC_LENGTH               = 8
)

func ParsingVoiceCall(d *st.Data, content string) {
	rows := strings.Split(content, "\n")

	for _, row := range rows {
		if row == "" {
			continue
		}

		str := strings.Split(row, ";")
		var voiceCall st.VoiceCallData

		if validate.VoiceCall(str, VC_LENGTH, VC_COUNTRY, VC_PROVIDER) {

			connectionStability, err := convertToFloat(str, VC_CONNECTION_STABILITY)
			ttfb, err := convertToInt(str, VC_TTFB)
			voicePurity, err := convertToInt(str, VC_VOICE_PURITY)
			medianOfCallsTime, err := convertToInt(str, VC_MEDIAN_OF_CALLS_TIME)

			if err != nil {
				continue
			}

			voiceCall.Country = str[VC_COUNTRY]
			voiceCall.Bandwidth = str[VC_BANDWIDTH]
			voiceCall.ResponseTime = str[VC_RESPONSETIME]
			voiceCall.Provider = str[VC_PROVIDER]
			voiceCall.ConnectionStability = connectionStability
			voiceCall.TTFB = ttfb
			voiceCall.VoicePurity = voicePurity
			voiceCall.MedianOfCallsTime = medianOfCallsTime

			d.VoiceCall = append(d.VoiceCall, voiceCall)
		}

	}

	return
}

func convertToInt(s []string, indx int) (i int, err error) {
	i, err = strconv.Atoi(s[indx])
	return
}

func convertToFloat(s []string, indx int) (f float32, err error) {
	n, err := strconv.ParseFloat(s[indx], 32)
	f = float32(n)
	return
}
