package app_utils

import (
	"strings"

	mdl "github.com/AMETORY/ametory-erp-modules/shared/models"
)

func GetThumbnail(files []mdl.FileModel) (*mdl.FileModel, []mdl.FileModel) {
	restFiles := []mdl.FileModel{}
	var thumbnail *mdl.FileModel
	for _, v := range files {
		if strings.HasPrefix(v.MimeType, "image/") && thumbnail == nil {
			thumbnail = &v
		} else {
			restFiles = append(restFiles, v)
		}
	}
	return thumbnail, restFiles
}
