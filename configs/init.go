package configs

import "embed"

//go:embed config.*.yaml
var TemplateFs embed.FS
