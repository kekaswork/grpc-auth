package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/kekaswork/grpc-auth/internal/app"
	"github.com/kekaswork/grpc-auth/internal/config"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	// init app configuration
	cfg := config.MustLoad()
	_ = cfg

	// init logger
	log := setupLogger(cfg.Env)
	_ = log

	// init main application
	mainApp := app.New(log, cfg.GRPC.Port, cfg.StoragePath, cfg.TokenTTL)

	// Run gRPC server
	go mainApp.GRPCSrv.MustRun()

	// gracefull shotdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	signal := <-stop
	log.Info("application stopped", slog.String("signal", signal.String()))
	mainApp.GRPCSrv.Stop()
}

func setupLogger(env string) (log *slog.Logger) {
	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return
}
