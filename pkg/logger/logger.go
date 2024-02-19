package logger

import (
	"go-skeleton/internal/variable"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"log/slog"
	"os"
)

// New
// @Description: 初始化日志
// @param dirPath
// @param fileName
// @param level 设置日志为Error级别 debug: -4, info: 0, warn: 4, error: 8
// @return *slog.Logger
// @author cx
func New(dirPath, fileName string, level slog.Level) *slog.Logger {
	r := &lumberjack.Logger{
		Filename:   dirPath + "/" + fileName + ".log",        // 文件名称
		MaxSize:    variable.Config.GetInt("log.maxSize"),    // 文件最大大小 1M
		MaxAge:     variable.Config.GetInt("log.maxAge"),     // 最大保留时间 1天
		MaxBackups: variable.Config.GetInt("log.maxBackups"), // 最大保留文件数 3个
		LocalTime:  variable.Config.GetBool("log.localTime"), // 是否用本机时间
		Compress:   variable.Config.GetBool("log.compress"),  // 是否压缩归档日志
	}
	var dlvl slog.LevelVar
	dlvl.Set(level)

	// 同时输出到控制台和日志
	writer := io.MultiWriter(os.Stdout, r)
	log4 := slog.New(slog.NewJSONHandler(writer, &slog.HandlerOptions{
		AddSource: false, // true: 日志在源文件中的位置 是go文件
		Level:     &dlvl,
	}))
	slog.SetDefault(log4)
	return log4
}
