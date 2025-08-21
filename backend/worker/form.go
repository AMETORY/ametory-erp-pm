package worker

import (
	"ametory-pm/services/app"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/AMETORY/ametory-erp-modules/customer_relationship"
	"github.com/AMETORY/ametory-erp-modules/utils"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"gopkg.in/olahol/melody.v1"
)

func FormDownload(erpContext *context.ERPContext) {
	log.Println("START FORM DOWNLOAD WORKER")
	appService, ok := erpContext.AppService.(*app.AppService)

	csrSevice, ok2 := erpContext.CustomerRelationshipService.(*customer_relationship.CustomerRelationshipService)
	if !ok {
		panic("CustomerRelationshipService is not found")
	}
	if ok && ok2 {
		dataSub := appService.Redis.Subscribe(*erpContext.Ctx, "FORM:DOWNLOAD")
		for {
			msg, err := dataSub.ReceiveMessage(*erpContext.Ctx)
			if err != nil {
				log.Println(err)
				continue
			}
			var formDownloadData map[string]any
			err = json.Unmarshal([]byte(msg.Payload), &formDownloadData)
			if err != nil {
				log.Println(err)
				continue
			}
			// log.Println("FORM DOWNLOAD", reflect.TypeOf(formDownloadData["user"]))
			// utils.LogJson(formDownloadData)

			time.Sleep(time.Second * 1)
			user, ok := formDownloadData["user"].(map[string]any)
			if !ok {
				log.Println("ERROR DOWNLOAD FORM", "user not found")
				return
			}
			form, err := csrSevice.FormService.GetForm(formDownloadData["id"].(string))
			if err != nil {
				log.Println("ERROR DOWNLOAD FORM", err)
				return
			}

			row := 1

			file := excelize.NewFile()
			sheet1 := file.GetSheetName(0)
			headerStyle, _ := file.NewStyle(&excelize.Style{
				Font: &excelize.Font{
					Bold: true,
					Size: 12,
				},
				Alignment: &excelize.Alignment{
					Horizontal: "center",
					Vertical:   "center",
				},
				Fill: excelize.Fill{
					Type:    "pattern",
					Color:   []string{"#DCE6F1"}, // Soft blue
					Pattern: 1,
				},
			})
			headerDescStyle, _ := file.NewStyle(&excelize.Style{
				Font: &excelize.Font{
					Size: 10,
				},
				Fill: excelize.Fill{
					Type:    "pattern",
					Color:   []string{"#DCE6F1"}, // Soft blue
					Pattern: 1,
				},
				Alignment: &excelize.Alignment{
					Horizontal: "center",
					Vertical:   "center",
				},
			})

			// file.SetColWidth(sheet1, "A", "Z", 20)
			fieldCol := 0
			for _, v := range form.FormTemplate.Sections {
				for _, v := range v.Fields {
					colWidth := max(len(v.Label)+5, 20)
					file.SetColWidth(sheet1, utils.NumToAlphabet(fieldCol+1), utils.NumToAlphabet(fieldCol+1), float64(colWidth))
					fieldCol++
				}
			}
			// create headers
			headerCol := 0
			for _, v := range form.FormTemplate.Sections {
				countFields := len(v.Fields)
				cell := utils.NumToAlphabet(headerCol + 1)
				file.SetCellValue(sheet1, fmt.Sprintf("%s%d", cell, row), v.SectionTitle)
				headerCol += countFields
				cellEnd := utils.NumToAlphabet(headerCol)
				file.MergeCell(sheet1, fmt.Sprintf("%s%d", cell, row), fmt.Sprintf("%s%d", cellEnd, row))
				file.SetCellStyle(sheet1, fmt.Sprintf("%s%d", cell, row), fmt.Sprintf("%s%d", cellEnd, row), headerStyle)
			}
			row++
			headerColDesc := 0
			for _, v := range form.FormTemplate.Sections {
				countFields := len(v.Fields)
				cell := utils.NumToAlphabet(headerColDesc + 1)
				file.SetCellValue(sheet1, fmt.Sprintf("%s%d", cell, row), v.Description)
				headerColDesc += countFields
				cellEnd := utils.NumToAlphabet(headerColDesc)
				file.MergeCell(sheet1, fmt.Sprintf("%s%d", cell, row), fmt.Sprintf("%s%d", cellEnd, row))
				file.SetCellStyle(sheet1, fmt.Sprintf("%s%d", cell, row), fmt.Sprintf("%s%d", cellEnd, row), headerDescStyle)
			}
			row++

			// Create header fields
			fieldCol = 0
			for _, v := range form.FormTemplate.Sections {
				for _, v := range v.Fields {
					file.SetCellValue(sheet1, fmt.Sprintf("%s%d", utils.NumToAlphabet(fieldCol+1), row), v.Label)
					file.SetCellStyle(sheet1, fmt.Sprintf("%s%d", utils.NumToAlphabet(fieldCol+1), row), fmt.Sprintf("%s%d", utils.NumToAlphabet(fieldCol+1), row), headerStyle)
					fieldCol++
				}
			}
			row++

			// WRITE RESPONSE
			respCol := 0
			for _, r := range form.Responses {
				for _, s := range r.Sections {
					for _, field := range s.Fields {
						field.Value = strings.TrimSpace(field.Value.(string))
						file.SetCellValue(sheet1, fmt.Sprintf("%s%d", utils.NumToAlphabet(respCol+1), row), field.Value)
						respCol++
					}
				}
				row++
				respCol = 0
			}

			var buf bytes.Buffer
			if err := file.Write(&buf); err != nil {
				log.Println(err)
				return
			}

			fileName := fmt.Sprintf("form_%v.xlsx", form.ID)
			f, err := os.Create(fileName)
			if err != nil {
				log.Println(err)
				return
			}
			defer f.Close()

			if _, err := f.Write(buf.Bytes()); err != nil {
				log.Println(err)
				return
			}

			// log.Println("END FORM DOWNLOAD WORKER", fileName)

			log.Println("START SEND MAIL to", user["email"].(string))

			erpContext.EmailSender.SetAddress(user["full_name"].(string), user["email"].(string))
			erpContext.EmailSender.SetTemplate("../templates/email/layout.html", "../templates/email/form_response.html")
			if err := erpContext.EmailSender.SendEmail(fmt.Sprintf(`Generate Form %s`, form.Title),
				map[string]any{
					"Form":         form,
					"File":         fileName,
					"UserFullName": user["full_name"].(string),
				},
				[]string{fileName}); err != nil {
				log.Println(err)
				fmt.Println(err)
				continue
			}

			os.Remove(fileName)

			broadcastMsg := gin.H{
				"message":     "form generated, please check your email at " + user["email"].(string),
				"form_id":     form.ID,
				"command":     "FORM_GENERATED",
				"sender_id":   user["id"].(string),
				"sender_name": user["full_name"].(string),
			}
			b, _ := json.Marshal(broadcastMsg)
			appService.Websocket.BroadcastFilter(b, func(q *melody.Session) bool {
				url := fmt.Sprintf("%s/api/v1/ws/%s", appService.Config.Server.BaseURL, formDownloadData["company_id"].(string))
				return fmt.Sprintf("%s%s", appService.Config.Server.BaseURL, q.Request.URL.Path) == url
			})
		}
	}
}
