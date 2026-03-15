package handlers

import (
	"net/http"
	"web-starter/visual"
	"web-starter/internal/app"
)

func RobotsTxt() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./visual/static/robots.txt")
	})
}

func Healthcheck(app *app.Application) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("All good!"))
	})
}

func Home(app *app.Application) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		component := visual.Html()
		component.Render(r.Context(), w)
	})
}
