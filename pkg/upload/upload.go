package upload

import (
	"errors"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"strconv"
	"time"

	"github.com/MQEnergy/go-skeleton/pkg/config"
	"github.com/MQEnergy/go-skeleton/pkg/helper"
	"github.com/MQEnergy/go-skeleton/pkg/oss"
)

const (
	// MaxUploadSize 默认最大上传资源大小是10M
	MaxUploadSize = 10 * 1024 * 1024
)

// AllowTypes 默认允许上传的文件类型
var AllowTypes = map[string]string{
	"jpg":  "image/jpg",
	"jpeg": "image/jpeg",
	"png":  "image/png",
	"svg":  "image/svg",
	"gif":  "image/gif",
	"bmp":  "image/bmp",
	"mp3":  "audio/mpeg",
	"mp4":  "video/mp4",
	"avi":  "video/x-msvideo",
	"rmvb": "video/vnd.rn-realmedia-vbr",
	"pdf":  "application/pdf",
	"xls":  "application/vnd.ms-excel",
	"xlsx": "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
	"ppt":  "application/vnd.ms-powerpoint",
	"doc":  "application/msword",
	"docx": "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
}

type File struct {
	FileName *multipart.FileHeader `form:"file"`
}

type Upload struct {
	MaxUploadSize int
	AllowTypes    map[string]string
}

// FileHeader 文件参数
type FileHeader struct {
	Filename   string `json:"file_name"`   // 图片新名称
	FileSize   int64  `json:"file_size"`   // 图片大小
	FilePath   string `json:"file_path"`   // 相对路径地址
	OriginName string `json:"origin_name"` // 图片原名称
	MimeType   string `json:"mime_type"`   // 附件mime类型
	Extension  string `json:"extension"`   // 附件后缀名
}

// New
// @Description: 携带上传参数 并且实例化
// @param max
// @param allowTypes 允许的文件类型 如：[]string{"jpg"", "png", "gif"}
// @return *Upload
func New(maxSize int, allowTypes []string) *Upload {
	if maxSize == 0 {
		maxSize = MaxUploadSize
	}
	allowMaps := make(map[string]string)
	if len(allowTypes) == 0 {
		allowMaps = AllowTypes
	} else {
		for _, allowType := range allowTypes {
			if allowMap, ok := AllowTypes[allowType]; ok {
				allowMaps[allowType] = allowMap
			}
		}
	}
	return &Upload{
		maxSize,
		allowMaps,
	}
}

// UploadToLocal
// @Description: 上传图片
// @receiver u
// @param file 请求的文件
// @param path 子目录名称
// @return *FileHeader
// @return error
func (u *Upload) UploadToLocal(config *config.Config, file *multipart.FileHeader, path string) (*FileHeader, error) {
	header, err := u.validate(file)
	if err != nil {
		return nil, err
	}
	// 创建时间目录 返回：./{server.fileUploadPath}/{path}/YYYY-MM-DD
	filePath, err := helper.MakeTimeFormatDir(config.GetString("server.fileUploadPath"), path, time.DateOnly)
	if err != nil {
		return nil, err
	}
	// 写入文件 返回：./{server.fileUploadPath}/{path}/YYYY-MM-DD/{fileName}
	if err := helper.WriteContentToFile(file, filePath+"/"+header.Filename); err != nil {
		return nil, err
	}
	header.FilePath = filePath + header.Filename
	return header, nil
}

// UploadToOss
// @Description: 上传资源到oss上
// @receiver u
// @param config
// @param file
// @param object 上传到oss的目录结构 支持子目录 如：foo foo/bar foo/bar/baz
// @return *FileHeader
// @return error
// @author cx
func (u Upload) UploadToOss(config *config.Config, file *multipart.FileHeader, object string) (*FileHeader, error) {
	header, err := u.validate(file)
	if err != nil {
		return nil, err
	}
	// 打开文件
	f, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer f.Close()
	fileBytes, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	// 实例化oss
	o, err := oss.New(&oss.Config{
		EndPoint:     config.GetString("oss.endPoint"),
		AccessId:     config.GetString("oss.accessKeyId"),
		AccessSecret: config.GetString("oss.accessKeySecret"),
		BucketName:   config.GetString("oss.bucketName"),
	})
	if err != nil {
		return nil, err
	}
	// 存储到oss
	filePath := config.GetString("server.fileUploadPath") + "/" + object + "/" + time.Now().Format(time.DateOnly) + "/" + header.Filename
	if err := o.PutObject(filePath, fileBytes); err != nil {
		return nil, err
	}
	header.FilePath = filePath
	return header, nil
}

// validate 验证请求的文件 并返回content-type
// validate
// @Description: 验证请求的文件 并返回值
// @receiver u
// @param file
// @return string 文件后缀  jpg, png等
// @return string 文件名称
// @return error
// @author cx
func (u *Upload) validate(file *multipart.FileHeader) (*FileHeader, error) {
	contentType := file.Header.Get("Content-Type")

	if int64(u.MaxUploadSize) < file.Size {
		return nil, errors.New("超过最大上传大小 不能超过" + strconv.Itoa(u.MaxUploadSize/(1000*1000)) + "M")
	}
	if !helper.InAnyMap[string](u.AllowTypes, contentType) {
		return nil, errors.New("上传文件格式错误")
	}
	filePrefix := helper.GetKeyByMap[string](u.AllowTypes, contentType)
	fileName := fmt.Sprintf("file-%s.%s", helper.GenerateUuid(32), filePrefix)
	return &FileHeader{
		Filename:   fileName,
		FileSize:   file.Size,
		FilePath:   "",
		OriginName: file.Filename,
		MimeType:   contentType,
		Extension:  filePrefix,
	}, nil
}
