package {{.PkgName}}

import (
    {{if .CmdDir -}}
	"{{.ImportPackage}}/internal/app/controller"
    {{end -}}
	"{{.ImportPackage}}/pkg/response"

	"github.com/gofiber/fiber/v2"
)

type {{.CtlName}}Controller struct{
    {{- if .CmdDir}}
	controller.Controller
	{{- else}}
	Controller
	{{- end}}
}

var {{.CtlName}} = &{{.CtlName}}Controller{}

// Index ...
func (s *{{.CtlName}}Controller) Index(ctx *fiber.Ctx) error {
    // Todo implement ...
	return response.SuccessJSON(ctx, "", "")
}