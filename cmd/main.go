package main

import (
	"os"
	"time"
	"net/http"
	"log/slog"
	"web-starter/internal/app"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	router := http.NewServeMux()

	app := &app.Application {
		Router: router,
		Logger: logger,
	}

	server := &http.Server {
		Handler: routes(app),
		WriteTimeout: 5 * time.Second,
		ReadTimeout: 5 * time.Second,
		IdleTimeout: 5 * time.Second,
		Addr: ":6969",
	}

	logger.Info("Starting server")

	error := server.ListenAndServeTLS("./tls/localhost+2.pem", "./tls/localhost+2-key.pem")
	logger.Error(error.Error())
	os.Exit(1)
}
