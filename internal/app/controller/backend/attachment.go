package backend

import (
	"github.com/MQEnergy/go-skeleton/internal/app/controller"
	"github.com/MQEnergy/go-skeleton/internal/request/attachment"
	"github.com/MQEnergy/go-skeleton/internal/vars"
	"github.com/MQEnergy/go-skeleton/pkg/response"
	"github.com/MQEnergy/go-skeleton/pkg/upload"
	"github.com/gofiber/fiber/v2"
)

type AttachmentController struct {
	controller.Controller
}

var Attachment = &AttachmentController{}

// Upload 上传资源
func (c *AttachmentController) Upload(ctx *fiber.Ctx) error {
	reqParams := &attachment.UploadReq{}
	if err := c.Validate(ctx, reqParams); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	fileName, err := ctx.FormFile("file")
	if err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	reqParams.FileName = fileName
	fileHeader, err := upload.New(0, []string{}).UploadToLocal(&vars.Config, reqParams.FileName, reqParams.FilePath)
	if err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	return response.SuccessJSON(ctx, "", fileHeader)
}
