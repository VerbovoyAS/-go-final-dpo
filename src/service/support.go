package service

import (
	"encoding/json"
	st "go-final-dpo/src/structure"
)

func ParsingSupport(d *st.Data, body []byte) {

	var supportData []st.SupportData
	if err := json.Unmarshal(body, &supportData); err != nil {
		return
	}

	for _, row := range supportData {
		d.Support = append(d.Support, row)
	}
}
