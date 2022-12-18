// Package app configures and runs application.
package app

import (
	"os"
	"os/signal"
	"syscall"

	"examples/kahootee/internal/di"
	"examples/kahootee/pkg/logger"
)

func RunHTTPServer() {
	httpServer, cleanup, err := di.InitializeHttpServer()
	if err != nil {
		logger.Fatal("failed to initialize http server: %v", err)
	}

	defer cleanup()

	httpServer.Start()

	waitSignal(httpServer.Notify())

	err = httpServer.Shutdown()
	if err != nil {
		logger.Error("app - Run - httpServer.Shutdown: %v", err)
	}
}

func waitSignal(err <-chan error) {
	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		logger.Info("app - Run - signal: " + s.String())
	case err := <-err:
		logger.Error("app - Run - httpServer.Notify: %v", err)
	}
}
