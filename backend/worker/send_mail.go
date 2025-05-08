package worker

import (
	"ametory-pm/objects"
	"ametory-pm/services/app"
	"encoding/json"
	"fmt"
	"log"

	"github.com/AMETORY/ametory-erp-modules/context"
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
			// fmt.Println("FullName", emailData.FullName)
			// fmt.Println("Email", emailData.Email)
			// fmt.Println("sender", sender)
			subject := "Welcome to AMETORY PROJECT MANAGER"
			if emailData.Subject != "" {
				subject = emailData.Subject
			}
			erpContext.EmailSender.SetAddress(emailData.FullName, emailData.Email)

			if err := erpContext.EmailSender.SendEmail(subject, emailData, []string{}); err != nil {
				log.Println(err)
				fmt.Println(err)
				continue
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
