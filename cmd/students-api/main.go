package main

import (
	"fmt"
	"log"
	"net/http"
	// "os"
	// "os/signal"
	// "syscall"

	"github.com/zuhaib042/students-api/internal/config"
)

func main() {
	// load Config
	cfg := config.MustLoad()
	// database setup


	// Setup Router
	router := http.NewServeMux()
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to students api"))
	})

	// Setup Server
	server := http.Server {
		Addr: cfg.HTTPServer.Addr,
		Handler: router,
	}

	fmt.Println("server started", cfg.HTTPServer.Addr)


		
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("failed to start server")
	}
}