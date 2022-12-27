package service

import (
	"encoding/json"
	st "go-final-dpo/src/structure"
)

func ParsingSupport(body []byte, d *st.Data) {

	var supportData []st.SupportData
	if err := json.Unmarshal(body, &supportData); err != nil {
		return
	}

	for _, row := range supportData {
		d.Support = append(d.Support, row)
	}
}
