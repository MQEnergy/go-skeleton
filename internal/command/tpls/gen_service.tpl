package {{.PkgName}}

{{if .CmdDir -}}
import (
	"{{.ImportPackage}}/internal/app/service"
)
{{end -}}

type {{.ServiceName}}Service struct{
    {{- if .CmdDir}}
	service.Service
	{{- else}}
	Service
	{{- end}}
}
var {{.ServiceName}} = &{{.ServiceName}}Service{}

// Index ...
func (s *{{.ServiceName}}Service) Index() error {
    // Todo implement ...
    return nil
}