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

	data := st.Data{Sms: []st.SMSData{}}

	content := service.GetContent(app.SkillboxSmsPath())
	service.ParsingSms(&data, content)
	service.CreateFile(&data)

	content = service.GetContent(app.DataSmsPath())
	service.ParsingSms(&data, content)

	fmt.Println(data)

}
