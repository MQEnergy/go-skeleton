package attachment

import "github.com/MQEnergy/go-skeleton/pkg/upload"

type UploadReq struct {
	upload.File
	FilePath string `form:"file_path"`
}
