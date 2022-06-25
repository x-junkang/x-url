package main

import (
	"github.com/x-junkang/x-url/internal/config"
	"github.com/x-junkang/x-url/internal/service"
	"github.com/x-junkang/x-url/internal/xlog"
)

func main() {
	cfg := config.DefaultConfig()
	logConf := xlog.Config{
		ConsoleLoggingEnabled: false,
		EncodeLogsAsJson:      true,
		FileLoggingEnabled:    true,
		Directory:             cfg.LogDir,
		Filename:              cfg.LogFile,
		MaxSize:               5,
		MaxBackups:            10,
		MaxAge:                7,
		Level:                 cfg.LogLevel,
	}
	xlog.Configure(logConf)

	svc := service.NewService(cfg)
	svc.Run()
}
