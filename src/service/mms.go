package service

import (
	"encoding/json"
	"fmt"
	st "go-final-dpo/src/structure"
	"go-final-dpo/src/validate"
)

func ParsingMms(body []byte, d *st.Data) {

	var mmsData []st.MMSData
	if err := json.Unmarshal(body, &mmsData); err != nil {
		fmt.Println(body)
		return
	}

	for _, row := range mmsData {
		if validate.Mms(row) {
			d.Mms = append(d.Mms, row)
		}
	}
}
