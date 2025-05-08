package worker

import (
	"ametory-pm/services/app"
	"bytes"
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"math/big"
	"net/http"
	"strings"

	"github.com/AMETORY/ametory-erp-modules/contact"
	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/AMETORY/ametory-erp-modules/shared/models"
	"github.com/AMETORY/ametory-erp-modules/utils"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
)

func ImportContact(erpContext *context.ERPContext) {
	fmt.Println("START IMPORT CONTACT")
	appService, ok := erpContext.AppService.(*app.AppService)
	if ok {
		dataSub := appService.Redis.Subscribe(*erpContext.Ctx, "IMPORT:CONTACT")
		for {
			fmt.Println("REQUEST IMPORT START")
			msg, err := dataSub.ReceiveMessage(*erpContext.Ctx)
			if err != nil {
				log.Println(err)
			}

			var data map[string]string
			err = json.Unmarshal([]byte(msg.Payload), &data)
			if err != nil {
				log.Println(err)
				continue
			}

			fmt.Println("GET FILE", data["file_url"])
			resp, err := http.Get(data["file_url"])
			if err != nil {
				log.Println(err)
				continue
			}
			defer resp.Body.Close()

			file, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Println(err)
				continue
			}
			xlsx, err := excelize.OpenReader(bytes.NewReader(file))
			if err != nil {
				log.Println(err)
				continue
			}
			firstSheet := xlsx.GetSheetName(0)
			rows, err := xlsx.GetRows(firstSheet)
			if err != nil {
				log.Println(err)
				continue
			}

			if err != nil {
				log.Println(err)
				continue
			}

			for i, row := range rows {
				if i == 0 {
					continue
				}
				for j, r := range row {
					fmt.Printf("%v.\t %s: %s\n", j+1, rows[0][j], r)
				}
				fmt.Println("------------------------------------------------")
				var phone *string
				if cleanString(rows[i][2]) != "" {
					phoneStr := utils.ParsePhoneNumber(cleanString(rows[i][2]), "")
					phone = &phoneStr
				}
				userID := data["user_id"]
				companyID := data["company_id"]

				var tags []models.TagModel

				if cleanString(rows[i][4]) != "" {
					dataTags := strings.Split(cleanString(rows[i][4]), ",")
					for _, v := range dataTags {
						fmt.Println("TAG", cleanString(v))
						var checkTag models.TagModel
						err := erpContext.DB.Where("name = ?", cleanString(v)).First(&checkTag).Error
						if err != nil {
							if errors.Is(err, gorm.ErrRecordNotFound) {
								tag := models.TagModel{
									Name:      cleanString(v),
									CompanyID: &companyID,
									Color:     randomColor(),
								}
								err := erpContext.DB.Create(&tag).Error
								if err != nil {
									log.Println(err)
									continue
								}
								tags = append(tags, tag)
							}
						} else {
							tags = append(tags, checkTag)
						}
					}
				}
				var products []models.ProductModel
				if cleanString(rows[i][5]) != "" {
					dataProducts := strings.Split(cleanString(rows[i][5]), ",")
					for _, v := range dataProducts {
						fmt.Println("PRODUCT", cleanString(v))
						var checkProduct models.ProductModel
						err := erpContext.DB.Where("name = ? or sku = ?", cleanString(v), cleanString(v)).First(&checkProduct).Error
						if err == nil {
							products = append(products, checkProduct)
						}
					}
				}

				newContact := models.ContactModel{
					Name:       rows[i][0],
					Email:      rows[i][1],
					Phone:      phone,
					UserID:     &userID,
					CompanyID:  &companyID,
					Tags:       tags,
					Products:   products,
					IsCustomer: true,
				}

				err := erpContext.ContactService.(*contact.ContactService).CreateContact(&newContact)
				if err != nil {
					log.Println(err)
					continue
				}
			}
		}
	}
}

func cleanString(s string) string {
	return strings.TrimSpace(s)
}

func randomHex(n int) string {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	return fmt.Sprintf("#%x", b)
}

func randomColor() string {
	colors := []string{
		"#FF5733", "#33FF57", "#3357FF", "#FF33A1", "#A133FF",
		"#33FFF5", "#F5FF33", "#FF8C33", "#FF3380", "#33FFBD",
		"#FFC400", "#00BFFF", "#FF00FF", "#008000", "#4B0082",
		"#9400D3", "#00FF00", "#FFD700", "#C71585", "#00FFFF",
		"#FF1493", "#00BFFF", "#32CD32", "#6495ED", "#DC143C",
		"#006400", "#8B008B", "#B03060", "#FF6347", "#1E90FF",
		"#228B22", "#FF7F24", "#FFC0CB", "#8B0A1A", "#4682B4",
		"#228B22", "#00688B", "#C71585", "#00BFFF", "#9400D3",
		"#FFD700", "#00FF00", "#FF1493", "#4B0082", "#9400D3",
		"#FFC400", "#00BFFF", "#FF00FF", "#008000", "#4B0082",
		"#9400D3", "#00FF00", "#FFD700", "#C71585", "#00FFFF",
		"#FF1493", "#00BFFF", "#32CD32", "#6495ED", "#DC143C",
		"#006400", "#8B008B", "#B03060", "#FF6347", "#1E90FF",
		"#228B22", "#FF7F24", "#FFC0CB", "#8B0A1A", "#4682B4",
	}
	n, err := rand.Int(rand.Reader, big.NewInt(int64(len(colors))))
	if err != nil {
		log.Println(err)
		return "#000000" // default to black if there's an error
	}
	return colors[n.Int64()]
}
