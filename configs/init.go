package configs

import "embed"

//go:embed config.*.yaml
var TemplateFs embed.FS

//go:embed rbac_model.conf
var RbacModelConf string
