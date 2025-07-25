package worker

import (
	"ametory-pm/config"
	"ametory-pm/objects"
	"ametory-pm/services/app"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"text/template"

	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/AMETORY/ametory-erp-modules/thirdparty/email_api"
)

func SendMail(erpContext *context.ERPContext) {
	fmt.Println("START SEND MAIL")
	appService, ok := erpContext.AppService.(*app.AppService)
	if ok {
		dataSub := appService.Redis.Subscribe(*erpContext.Ctx, "SEND:MAIL")
		for {
			msg, err := dataSub.ReceiveMessage(*erpContext.Ctx)
			if err != nil {
				log.Println(err)
			}
			var emailData objects.EmailData
			err = json.Unmarshal([]byte(msg.Payload), &emailData)
			if err != nil {
				log.Println(err)
				continue
			}

			// fmt.Println("sender", sender)
			subject := "Welcome to AMETORY PROJECT MANAGER"
			if emailData.Subject != "" {
				subject = emailData.Subject
			}

			fmt.Println("SEND MAIL")
			fmt.Println("======================================")
			fmt.Println("FullName", emailData.FullName)
			fmt.Println("Email", emailData.Email)
			fmt.Println("Subject", subject)

			if config.App.Email.UseAPI {
				senderAPI := email_api.NewEmailApiService(config.App.EmailApi.From, config.App.EmailApi.Domain, config.App.EmailApi.ApiKey, email_api.KirimEmail{})
				t := template.Must(template.ParseFiles("../templates/email/layout.html", "../templates/email/body.html"))

				var buf bytes.Buffer
				if err := t.ExecuteTemplate(&buf, "layout", emailData); err != nil {
					return
				}
				err = senderAPI.SendEmail(subject, emailData.Email, buf.String(), []string{})
				if err != nil {
					log.Println(err)
					fmt.Println(err)
					continue
				}

				continue
			} else {

				erpContext.EmailSender.SetAddress(emailData.FullName, emailData.Email)

				if err := erpContext.EmailSender.SendEmail(subject, emailData, []string{}); err != nil {
					log.Println(err)
					fmt.Println(err)
					continue
				}
			}

		}
	}
}

func TestEmail(erpContext context.ERPContext, email string) {
	var emailData objects.EmailData
	emailData.Subject = "Test Email"
	emailData.FullName = "Test Email"
	emailData.Email = email
	erpContext.EmailSender.SetAddress(emailData.FullName, emailData.Email)
	if err := erpContext.EmailSender.SendEmail(emailData.Subject, emailData, []string{}); err != nil {
		log.Println(err)
		fmt.Println(err)
	}
}
