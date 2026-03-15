package main

import (
	"net/http"
	"web-starter/internal/app"
	"web-starter/internal/headers"
	"web-starter/internal/handlers"
)

func routes(app *app.Application) http.Handler {
	fs := http.FileServer(http.Dir("./visual/static"))

	muxStatic := http.NewServeMux()
	muxStatic.Handle("/static/", http.StripPrefix("/static", fs))

	muxPages := http.NewServeMux()
	muxPages.Handle("GET /", handlers.Home(app))
	muxPages.Handle("GET /healthcheck", handlers.Healthcheck(app))
	muxPages.Handle("GET /robots.txt", handlers.RobotsTxt())

	app.Router.Handle("/static/", headers.CacheStatic(muxStatic))
	app.Router.Handle("/", headers.CachePages(muxPages))

	return headers.SecureHeaders(app.Router)
}
