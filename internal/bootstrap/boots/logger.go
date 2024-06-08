package boots

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/MQEnergy/go-skeleton/internal/vars"
	logger2 "github.com/MQEnergy/go-skeleton/pkg/logger"
)

// InitLogger ...
func InitLogger() error {
	vars.Once.Do(func() {
		slog.Info("Loading Logger service successfully")
		logger2.New(
			vars.Config.GetString("log.fileName"),
			&vars.Config,
		)
		fileName := fmt.Sprintf("%s/%s.log", vars.Config.Get("log.dirPath"), vars.Config.GetString("log.fileName"))
		os.Chmod(fileName, 0o644)
	})
	return nil
}
