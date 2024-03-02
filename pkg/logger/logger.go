package logger

import (
	"fmt"
	"io"
	"log/slog"
	"os"

	"github.com/MQEnergy/go-skeleton/pkg/config"

	"gopkg.in/natefinch/lumberjack.v2"
)

var dlvl slog.LevelVar

// New
// @Description: init
// @param dirPath
// @param fileName
// @param level Set the log level to Error debug: -4, info: 0, warn: 4, error: 8
// @return *slog.Logger
// @author cx
func New(fileName string, c *config.Config) *slog.Logger {
	writer := ApplyWriter(fileName, c)
	log4 := slog.New(slog.NewJSONHandler(writer, &slog.HandlerOptions{
		AddSource: false, // true: The output is printed at the location of the go source file
		Level:     &dlvl,
	}))
	slog.SetDefault(log4)
	return log4
}

// ApplyWriter ...
func ApplyWriter(fileName string, c *config.Config) io.Writer {
	r := &lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/%s.log", c.Get("log.dirPath"), fileName), // file name
		MaxSize:    c.GetInt("log.maxSize"),                                  // Maximum file size default 1M
		MaxAge:     c.GetInt("log.maxAge"),                                   // Maximum retention time 1天
		MaxBackups: c.GetInt("log.maxBackups"),                               // Maximum number of reserved files 3个
		LocalTime:  c.GetBool("log.localTime"),                               // Whether to use local time
		Compress:   c.GetBool("log.compress"),                                // Whether to compress archive logs
	}
	dlvl.Set(slog.Level(c.GetInt("log.level")))
	writer := io.MultiWriter(os.Stdout, r)
	if c.GetString("server.mode") == "production" {
		writer = r
	}
	return writer
}
