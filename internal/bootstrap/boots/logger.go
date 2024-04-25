package boots

import (
	"github.com/MQEnergy/go-skeleton/internal/vars"
	logger2 "github.com/MQEnergy/go-skeleton/pkg/logger"
)

// InitLogger ...
func InitLogger() {
	logger2.New(
		vars.Config.GetString("log.fileName"),
		vars.Config,
	)
}
