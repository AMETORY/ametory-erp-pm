package services

import (
	rapid_api_models "ametory-pm/models/rapid_api"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/AMETORY/ametory-erp-modules/utils"
)

type RapidApiService struct {
	ctx *context.ERPContext
}

func NewRapidAdpiService(erpContext *context.ERPContext) *RapidApiService {
	service := RapidApiService{
		ctx: erpContext,
	}
	if !erpContext.SkipMigration {
		erpContext.DB.AutoMigrate(
			&rapid_api_models.RapidApiPlugin{},
			&rapid_api_models.RapidApiEndpoint{},
			&rapid_api_models.RapidApiData{},
			&rapid_api_models.RapidApiDataHistory{},
			&rapid_api_models.CompanyRapidApiPlugin{},
		)
	}
	return &service
}

func (u *RapidApiService) GetPlugins() ([]rapid_api_models.RapidApiPlugin, error) {
	var plugins []rapid_api_models.RapidApiPlugin
	err := u.ctx.DB.Preload("RapidApiEndpoints").Find(&plugins).Error
	if err != nil {
		return nil, err
	}
	return plugins, nil
}

func (u *RapidApiService) GetData(plugin rapid_api_models.RapidApiPlugin, endpoint rapid_api_models.RapidApiEndpoint, params []map[string]any, companyID string) (interface{}, error) {

	if endpoint.Key == "video_details" && plugin.Key == "yt-api" {
		params = append(params, map[string]interface{}{
			"key":   "extend",
			"value": "1",
		})
	}
	q := endpoint.URL + "?"
	for _, param := range params {
		// for key, val := range param {
		q += fmt.Sprintf("%s=%v&", param["key"], param["value"])
		// }
	}

	var companyRapidApiPlugin rapid_api_models.CompanyRapidApiPlugin
	err := u.ctx.DB.Find(&companyRapidApiPlugin, "company_id = ? AND rapid_api_plugin_id = ?", companyID, endpoint.RapidApiPluginID).Error
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	url := q[:len(q)-1]

	fmt.Println("URL", url)
	req, _ := http.NewRequest(endpoint.Method, url, nil)

	// fmt.Println(u.CompanyID, u.RapidApiEndpoint.RapidApiPluginID)
	fmt.Println("x-rapidapi-key", companyRapidApiPlugin.RapidApiKey)
	fmt.Println("x-rapidapi-host", companyRapidApiPlugin.RapidApiHost)
	req.Header.Add("x-rapidapi-key", companyRapidApiPlugin.RapidApiKey)
	req.Header.Add("x-rapidapi-host", companyRapidApiPlugin.RapidApiHost)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var resBody interface{}
	err = json.Unmarshal(body, &resBody)
	if err != nil {
		return nil, err
	}

	// fmt.Println(res)
	fmt.Println(string(body))

	// util.LogJson(resBody)
	return u.ParseData(plugin, endpoint, resBody, companyID)

	// return &resBody, err
}

func (u *RapidApiService) ParseData(plugin rapid_api_models.RapidApiPlugin, endpoint rapid_api_models.RapidApiEndpoint, resp interface{}, companyID string) (*map[string]interface{}, error) {
	newData := make(map[string]interface{})

	data, ok := resp.(map[string]interface{})
	if !ok {
		dataSlice, ok := resp.([]interface{})
		if !ok {
			return nil, fmt.Errorf("unknown type of data")
		}
		data = dataSlice[0].(map[string]interface{})
	}
	thumbnails := []string{}
	// util.LogJson(data)
	// fmt.Println("PLUGIN NAME", u.RapidApiEndpoint.Key, u.RapidApiPlugin.Key)
	if endpoint.Key == "video_details" && plugin.Key == "yt-api" {
		newData = utils.ReduceMap(data, []string{
			"id",
			"title",
			"lengthSeconds",
			"channelTitle",
			"channelId",
			"description",
			"viewCount",
			"likeCount",
			"commentCount",
			"category",
			"publishDate",
		})

		// lastIndex := len(data["thumbnail"].([]interface{})) - 1
		for i := range data["thumbnail"].([]interface{}) {
			if i == 0 {
				thumbnail := data["thumbnail"].([]interface{})[i]

				resp, err := http.Get(thumbnail.(map[string]interface{})["url"].(string))
				if err != nil {
					return nil, err
				}
				defer resp.Body.Close()
				img, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, err
				}
				encodedImg := base64.StdEncoding.EncodeToString(img)
				// u.ThumbnailURL = "data:image/png;base64," + encodedImg
				thumbnails = append(thumbnails, "data:image/png;base64,"+encodedImg)
			}
		}

		metaDataUrl := fmt.Sprintf("https://yt-api.p.rapidapi.com/updated_metadata?id=%s", newData["id"])
		metaReq, err := http.NewRequest("GET", metaDataUrl, nil)
		if err != nil {
			return nil, err
		}

		var companyRapidApiPlugin rapid_api_models.CompanyRapidApiPlugin
		err = u.ctx.DB.Find(&companyRapidApiPlugin, "company_id = ? AND rapid_api_plugin_id = ?", companyRapidApiPlugin, endpoint.RapidApiPluginID).Error
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		metaReq.Header.Add("x-rapidapi-key", companyRapidApiPlugin.RapidApiKey)
		metaReq.Header.Add("x-rapidapi-host", companyRapidApiPlugin.RapidApiHost)

		metaRes, err := http.DefaultClient.Do(metaReq)
		if err != nil {
			return nil, err
		}
		defer metaRes.Body.Close()

		metaBody, err := io.ReadAll(metaRes.Body)
		if err != nil {
			return nil, err
		}

		var metaData map[string]interface{}
		err = json.Unmarshal(metaBody, &metaData)
		if err != nil {
			return nil, err
		}

		// Log or process the metadata as needed
		fmt.Println("Meta Data likeCount:", metaData["likeCount"])
		fmt.Println("Meta Data viewCount:", metaData["viewCount"])

		newData["likeCount"] = metaData["likeCount"]

		newData["viewCount"] = metaData["viewCount"]
		viewCount, err := strconv.Atoi(metaData["viewCount"].(string))
		if err == nil {
			newData["viewCount"] = viewCount
		}

	}
	if endpoint.Key == "get_facebook_post_details" && plugin.Key == "facebook-scraper-api4" {
		if details, exists := data["details"]; exists {
			newDetails := utils.ReduceMap(details.(map[string]interface{}), []string{
				"post_id",
				"post_link",
				"share_count",
				"play_count",
				"comments_count",
			})

			for k, v := range newDetails {
				newData[k] = v
			}
		}

		newData["reactions"] = data["reactions"]
		newData["user_details"] = data["user_details"]
		newData["values"] = data["values"]

		if attachments, exists := data["attachments"]; exists {
			for _, attachment := range attachments.([]interface{}) {
				if photo, exists := attachment.(map[string]interface{})["photo_image"]; exists {
					// util.LogJson(photo)
					resp, err := http.Get(photo.(map[string]interface{})["uri"].(string))
					if err != nil {
						return nil, err
					}
					defer resp.Body.Close()
					img, err := io.ReadAll(resp.Body)
					if err != nil {
						return nil, err
					}
					encodedImg := base64.StdEncoding.EncodeToString(img)
					thumbnails = append(thumbnails, "data:image/png;base64,"+encodedImg)
					// u.ThumbnailURL = photo.(map[string]interface{})["url"].(string)
				} else {
					if thumbnailUrl, exists := attachment.(map[string]interface{})["thumbnail_url"]; exists {
						// util.LogJson(photo)
						resp, err := http.Get(thumbnailUrl.(string))
						if err != nil {
							return nil, err
						}
						defer resp.Body.Close()
						img, err := io.ReadAll(resp.Body)
						if err != nil {
							return nil, err
						}
						encodedImg := base64.StdEncoding.EncodeToString(img)
						thumbnails = append(thumbnails, "data:image/png;base64,"+encodedImg)
						// u.ThumbnailURL = photo.(map[string]interface{})["url"].(string)
					}

				}
			}

		}
	}

	if itemInfo, exists := data["itemInfo"]; exists {
		if itemStruct, exists := itemInfo.(map[string]interface{})["itemStruct"]; exists {
			newData = utils.ReduceMap(itemStruct.(map[string]interface{}), []string{
				"id",
				"desc",
				"video",
				"author",
				"stats",
			})

			video, ok := newData["video"]
			if ok {
				resp, err := http.Get(video.(map[string]interface{})["cover"].(string))
				if err != nil {
					return nil, err
				}
				defer resp.Body.Close()
				img, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, err
				}
				encodedImg := base64.StdEncoding.EncodeToString(img)
				thumbnails = append(thumbnails, "data:image/png;base64,"+encodedImg)

			}
		}
	}
	if tweet, exists := data["tweet"]; exists {
		newData = utils.ReduceMap(tweet.(map[string]interface{}), []string{
			"id_str",
			"bookmark_count",
			"favorite_count",
			"quote_count",
			"reply_count",
			"retweet_count",
			"created_at",
			"full_text",
			"entities",
		})
		entities, ok := newData["entities"]
		if ok {
			media, ok2 := entities.(map[string]interface{})["media"]
			if ok2 {
				lastIndex := len(media.([]interface{})) - 1
				for i, v := range media.([]interface{}) {
					if i == lastIndex {
						resp, err := http.Get(v.(map[string]interface{})["media_url_https"].(string))
						if err != nil {
							return nil, err
						}
						defer resp.Body.Close()
						img, err := io.ReadAll(resp.Body)
						if err != nil {
							return nil, err
						}
						encodedImg := base64.StdEncoding.EncodeToString(img)
						thumbnails = append(thumbnails, "data:image/png;base64,"+encodedImg)
					}
				}
			}
		}
	}
	if user, exists := data["user"]; exists {
		if legacy, exists := user.(map[string]interface{})["legacy"]; exists {
			newData["user"] = legacy
		}
	}
	if media, exists := data["media_or_ad"]; exists {
		newData = utils.ReduceMap(media.(map[string]interface{}), []string{
			"pk",
			"id",
			"play_count",
			"user",
			"comment_count",
			"like_count",
			"save_count",
			"reshare_count",
			"caption",
			"product_type",
			"video_versions",
			"image_versions2",
		})
		thumbnail, ok := newData["image_versions2"]
		if ok {
			candidates, ok2 := thumbnail.(map[string]interface{})["candidates"]
			if ok2 {
				imgCandidates := candidates.([]interface{})
				lastIndex := len(imgCandidates) - 1
				fmt.Println("lastIndex", lastIndex)
				fmt.Println("len(imgCandidates)", len(imgCandidates))
				for i, v := range imgCandidates {
					if i == lastIndex {
						resp, err := http.Get(v.(map[string]interface{})["url"].(string))
						if err != nil {
							return nil, err
						}
						defer resp.Body.Close()
						img, err := io.ReadAll(resp.Body)
						if err != nil {
							return nil, err
						}
						encodedImg := base64.StdEncoding.EncodeToString(img)
						thumbnails = append(thumbnails, "data:image/png;base64,"+encodedImg)
					}

				}
			}

		}
	}

	// u.ctx.DB.Save(&u)
	newData["endpoint_key"] = endpoint.Key
	newData["thumbnails"] = thumbnails
	return &newData, nil
}

func (u *RapidApiService) ExtractMap(data map[string]interface{}, key string) (map[string]interface{}, error) {
	if val, exists := data[key]; exists {
		if valMap, ok := val.(map[string]interface{}); ok {
			return valMap, nil
		}
	}

	return nil, fmt.Errorf("key %s not exists or not map", key)
}
