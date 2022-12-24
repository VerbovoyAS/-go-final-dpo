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

	data := st.Data{}

	//SMS
	content := service.GetContent(app.SkillboxSmsPath())
	service.ParsingSms(&data, content)
	service.CreateFile(data.VoiceSMSContent(), app.DataSmsPath())

	content = service.GetContent(app.DataSmsPath())
	service.ParsingSms(&data, content)

	// MMS
	body, err := app.Request(app.PathMms())
	if err != nil {
		panic(err)
	}
	service.ParsingMms(body, &data)

	//VoiceCall
	content = service.GetContent(app.SkillboxVoiceCallPath())
	service.ParsingVoiceCall(&data, content)
	service.CreateFile(data.VoiceCallContent(), app.DataVoicePath())

	// Email
	content = service.GetContent(app.SkillboxEmailPath())
	service.ParsingEmail(&data, content)
	service.CreateFile(data.EmailContent(), app.DataEmailPath())
	fmt.Println(data)
}
