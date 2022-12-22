package main

import (
	"fmt"
	"go-final-dpo/src/app"
	"go-final-dpo/src/service"
	st "go-final-dpo/src/structure"
)

func init() {
	app.InitEnv()
	app.InitEnum()
}

func main() {
	fmt.Println("")

	data := st.Data{Sms: []st.SMSData{}, Mms: []st.MMSData{}}

	content := service.GetContent(app.SkillboxSmsPath())
	service.ParsingSms(&data, content)
	service.CreateFile(&data, app.DataSmsPath())

	content = service.GetContent(app.DataSmsPath())
	service.ParsingSms(&data, content)

	body, err := app.Request(app.PathMms())
	if err != nil {
		panic(err)
	}

	service.ParsingMms(body, &data)

	fmt.Println(data)
}
