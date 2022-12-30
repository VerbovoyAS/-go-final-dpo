package main

import (
	"github.com/gorilla/mux"
	"go-final-dpo/src/app"
	"go-final-dpo/src/service"
	st "go-final-dpo/src/structure"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	app.InitEnv()
	app.InitEnum()
}

func main() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	router := mux.NewRouter()
	router.HandleFunc(app.PathMyServerUrl(), func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("Не верно указан метод, используйте метод GET"))
			return
		}

		data := st.Data{}
		getContentSkillbox(&data)
		res := service.GetResultData(data)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(res.ToJson())
	})
	go func() {
		if err := http.ListenAndServe(app.PathMyServer(), router); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Print("Server Started")

	<-done
	log.Print("Server Stopped")
}

// getContentSkillbox Получение информации от внешнего сервиса
func getContentSkillbox(data *st.Data) {
	//SMS
	content := service.GetContent(app.SkillboxSmsPath())
	service.ParsingSms(data, content)
	service.CreateFile(data.VoiceSMSContent(), app.DataSmsPath())

	// MMS
	body := app.Request(app.PathMms())
	service.ParsingMms(body, data)

	//VoiceCall
	content = service.GetContent(app.SkillboxVoiceCallPath())
	service.ParsingVoiceCall(data, content)
	service.CreateFile(data.VoiceCallContent(), app.DataVoicePath())

	// Email
	content = service.GetContent(app.SkillboxEmailPath())
	service.ParsingEmail(data, content)
	service.CreateFile(data.EmailContent(), app.DataEmailPath())

	// Billing
	content = service.GetContent(app.SkillboxBillingPath())
	service.ParsingBilling(data, content)

	// Support
	body = app.Request(app.PathSupport())
	service.ParsingSupport(data, body)

	// Incident
	body = app.Request(app.PathIncident())
	service.ParsingIncident(data, body)
}
