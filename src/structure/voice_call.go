package st

import "fmt"

type VoiceCallData struct {
	Country             string
	Bandwidth           string
	ResponseTime        string
	Provider            string
	ConnectionStability float32
	TTFB                int
	VoicePurity         int
	MedianOfCallsTime   int
}

// ToString преобразует структуру в строку
func (s *VoiceCallData) ToString() string {
	return fmt.Sprintf(
		"%s;%s;%s;%s;%f;%d;%d;%d\n",
		s.Country, s.Bandwidth, s.ResponseTime, s.Provider, s.ConnectionStability, s.TTFB, s.VoicePurity, s.MedianOfCallsTime)
}
