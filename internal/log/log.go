package log

import (
	"log/slog"
	"os"
	"sync"
	"template/internal/lib/logger/handlers/slogpretty"
	"template/internal/lib/logger/writers/jsonwrt"
)


const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

var Logger *slog.Logger
var once sync.Once

func NewLogger(env string) {

	once.Do(func() {
        Logger = setupLogger(env)
	})
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
    if _, err := os.Stat("data"); os.IsNotExist(err) {
        err := os.Mkdir("data", 0755)
        if err != nil{
            panic("Error creating data dir: " + err.Error())
        }
    }

    if _, err := os.Stat("data/logs"); os.IsNotExist(err) {
        err := os.Mkdir("data/logs", 0755)
        if err != nil{
            panic("Error creating data/logs dir: " + err.Error())
        }
    }

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


