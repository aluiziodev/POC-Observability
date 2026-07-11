package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"order-service/internal/config"
	"order-service/internal/observability"
	"order-service/internal/router"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	logger := observability.NewLogger()

	cfg := config.Load()

	router := router.GenerateRouter(logger)

	server := &http.Server{
		Addr:              ":" + cfg.Port,
		Handler:           router,
		ReadHeaderTimeout: cfg.ReadTimeout,
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	serverErr := make(chan error, 1)
	go func() {
		logger.Info("Server iniciando", slog.String("port", cfg.Port))
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			serverErr <- err
			return
		}
		serverErr <- nil
	}()

	select {
	case err := <-serverErr:
		if err != nil {
			logger.Error("Erro no inicio do servidor!", slog.String("error", err.Error()))
			os.Exit(1)
		}
	case <-ctx.Done():
		logger.Info("Sinal de desligamento gracioso recebido...")

		shutdownCtx, cancel := context.WithTimeout(context.Background(), cfg.ShutDownTime)
		defer cancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			logger.Error("Erro ao encerrar o servidor de forma graciosa!!", slog.String("error", err.Error()))
			os.Exit(1)
		}

		logger.Info("Server encerrado de forma graciosa")
	}
}
