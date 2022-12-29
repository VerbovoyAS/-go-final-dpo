package service

import (
	"encoding/json"
	st "go-final-dpo/src/structure"
)

func ParsingIncident(d *st.Data, body []byte) {

	var incidentData []st.IncidentData
	if err := json.Unmarshal(body, &incidentData); err != nil {
		return
	}

	for _, row := range incidentData {
		d.Incident = append(d.Incident, row)
	}
}
