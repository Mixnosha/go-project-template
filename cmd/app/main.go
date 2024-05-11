package main

import (
	"log/slog"
	"os"

	"template/internal/config"
	"template/internal/lib/logger/handlers/slogpretty"
	"template/internal/lib/logger/writers/jsonwrt"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)


func main(){
    config.InitEnv()

    cfg := config.MustLoad()

    log := setupLogger(cfg.Env)

    log.Info("Program started with: ", slog.String("env", cfg.Env))

}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = setupPrettySlog()
	case envDev:
		log = setupFileOutputSlog(&slog.HandlerOptions{Level: slog.LevelDebug})
	case envProd:
		log = setupFileOutputSlog(&slog.HandlerOptions{Level: slog.LevelInfo})
	}

	return log
}

func setupFileOutputSlog(
	hadlerOptions *slog.HandlerOptions,
) *slog.Logger {
	writer := jsonwrt.NewJsonWriter("data/logs/")
	log := slog.New(
		slog.NewJSONHandler(writer, hadlerOptions),
	)
	return log
}

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}


