package main

import (
	"flag"
	"net/http"
	"os"
	"log/slog"
)

type application struct {
	logger *slog.Logger
}

func main() {
	addr := flag.String("addr", ":8080", "HTTP network address")
	flag.Parse()
	
	// creating a custom logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
	}))

	app := &application{
		logger: logger,
	}

	// Initialize a new server mux
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("../frontend/dist"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("/{$}", app.home)

	// Start the HTTP server
	logger.Info("Starting server", "addr", *addr)
	err := http.ListenAndServe(*addr, mux)
	if err != nil {
		logger.Error("Error starting server", "error", err)
	}
}