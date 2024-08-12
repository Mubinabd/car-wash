package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	r "github.com/Mubinabd/car-wash/internal/http"
	"github.com/Mubinabd/car-wash/internal/http/handlers"
	"github.com/Mubinabd/car-wash/logger"
)

func main() {
	handlers, err := handlers.NewHandlers()
	if err != nil {
		logger.Fatal("Failed to initialize handlers: ", err)
	}

	engine := r.NewRouter(handlers)

	server := &http.Server{
		Addr:    ":8080",
		Handler: engine,
	}

	go func() {
		logger.Info("API Gateway started successfully on :8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("Failed to start server: ", err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	signalReceived := <-sigChan
	logger.Info("Received signal:", signalReceived)

	shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownRelease()

	if err := server.Shutdown(shutdownCtx); err != nil {
		logger.Fatal("Server shutdown error: ", err)
	}

	logger.Info("Graceful shutdown complete.")
}
