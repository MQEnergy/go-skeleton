package boots

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/MQEnergy/go-skeleton/internal/vars"
	logger2 "github.com/MQEnergy/go-skeleton/pkg/logger"
)

// InitLogger ...
func InitLogger() {
	vars.Once.Do(func() {
		s := logger2.New(
			vars.Config.GetString("log.fileName"),
			&vars.Config,
		)
		fileName := fmt.Sprintf("%s/%s.log", vars.Config.Get("log.dirPath"), vars.Config.GetString("log.fileName"))
		s.Info("Loading Logger service successfully")
		_ = os.Chmod(fileName, 0o644)
	})
}
