package main

import (
	"flag"
	"fmt"

	"github.com/x-junkang/x-url/internal/config"
	"github.com/x-junkang/x-url/internal/model"
	"github.com/x-junkang/x-url/internal/service"
	"github.com/x-junkang/x-url/internal/xlog"
)

var cfgPath = flag.String("c", "etc/test.json", "config file path")

func main() {
	flag.Parse()

	cfg := config.DefaultConfig()
	err := cfg.LoadFromJson(*cfgPath)
	if err != nil {
		fmt.Println(err)
		return
	}

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
	err = model.InitDB(cfg.MySqlDSN)
	if err != nil {
		panic(err)
	}
	svc := service.NewService(cfg)
	svc.Run()
}
