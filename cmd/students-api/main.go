package main

import (
	"context"
	// "fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/zuhaib042/students-api/internal/config"
	"github.com/zuhaib042/students-api/internal/http/student"
)

func main() {
	// load Config
	cfg := config.MustLoad()
	// database setup


	// Setup Router
	router := http.NewServeMux()
	router.HandleFunc("POST /api/students", student.New())

	// Setup Server
	server := http.Server {
		Addr: cfg.HTTPServer.Addr,
		Handler: router,
	}

	slog.Info("server started", slog.String("address", cfg.Addr))
	// fmt.Println("server started", cfg.HTTPServer.Addr)

	// GRACEFULLY STOPPING THE SERVER
	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatal("failed to start server")
		}
	}()
	<-done
	
	slog.Info("shutting down the server")
	
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		slog.Error("failed to shutdown server", slog.String("error", err.Error()))
	}

	slog.Info("server shutdown successfully")
}