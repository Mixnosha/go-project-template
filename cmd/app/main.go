package main

import (
	"log/slog"

	"template/internal/config"
	"template/internal/log"
)



func main(){
    config.InitEnv()

    cfg := config.MustLoad()

    log.NewLogger(cfg.Env)

    log.Logger.Info("Program started with: ", slog.String("env", cfg.Env))
}

