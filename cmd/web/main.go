package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"

	"web-starter/internal/cache"
	"web-starter/internal/compression"
)

type application struct {
	logger *slog.Logger
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	app := application{
		logger: logger,
	}

	server := &http.Server{
		Addr:              ":6969",
		Handler:           app.router(),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      20 * time.Second,
		IdleTimeout:       60 * time.Second,
		MaxHeaderBytes:    1 << 20,
	}

	logger.Info("starting server")

	err := server.ListenAndServe()
	logger.Error(err.Error())
	os.Exit(1)
}

func (app *application) router() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", app.handlerIndex)
	mux.HandleFunc("GET /dashboard", app.handlerDashboard)
	mux.HandleFunc("GET /warehouse", app.handlerWarehouse)
	mux.HandleFunc("GET /user", app.handlerUser)

	static := http.FileServer(http.Dir("./visual/static"))
	manifest := cache.CacheStatic(http.FileServer(
		http.Dir("./visual/static/favicon"),
	))

	mux.Handle("GET /robots.txt", manifest)
	mux.Handle("GET /site.webmanifest", manifest)
	mux.Handle("GET /favicon.ico", manifest)
	mux.Handle("GET /favicon-16x16.png", manifest)
	mux.Handle("GET /favicon-32x32.png", manifest)
	mux.Handle("GET /apple-touch-icon.png", manifest)
	mux.Handle("GET /android-chrome-192x192.png", manifest)
	mux.Handle("GET /android-chrome-512x512.png", manifest)

	mux.Handle("GET /static/",
		http.StripPrefix("/static/",
			compression.GzipMiddleware(
				cache.CacheStatic(static))),
	)

	return mux
}
