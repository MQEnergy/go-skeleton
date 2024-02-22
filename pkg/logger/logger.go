package logger

import (
	"fmt"
	"go-skeleton/internal/variable"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"log/slog"
	"os"
)

// New
// @Description: init
// @param dirPath
// @param fileName
// @param level Set the log level to Error debug: -4, info: 0, warn: 4, error: 8
// @return *slog.Logger
// @author cx
func New(dirPath, fileName string, level slog.Level) *slog.Logger {
	r := &lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/%s.log", dirPath, fileName), // file name
		MaxSize:    variable.Config.GetInt("log.maxSize"),       // Maximum file size default 1M
		MaxAge:     variable.Config.GetInt("log.maxAge"),        // Maximum retention time 1天
		MaxBackups: variable.Config.GetInt("log.maxBackups"),    // Maximum number of reserved files 3个
		LocalTime:  variable.Config.GetBool("log.localTime"),    // Whether to use local time
		Compress:   variable.Config.GetBool("log.compress"),     // Whether to compress archive logs
	}
	var dlvl slog.LevelVar
	dlvl.Set(level)
	writer := io.MultiWriter(os.Stdout, r)
	if variable.Config.GetString("server.mode") == "production" {
		writer = r
	}
	log4 := slog.New(slog.NewJSONHandler(writer, &slog.HandlerOptions{
		AddSource: false, // true: The output is printed at the location of the go source file
		Level:     &dlvl,
	}))
	slog.SetDefault(log4)
	return log4
}
