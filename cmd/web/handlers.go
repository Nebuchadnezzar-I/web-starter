package main

import (
	"net/http"
	ui "web-starter/visual"
)

func (app *application) handlerIndex(w http.ResponseWriter, r *http.Request) {
	component := ui.Index("Home")

	w.Header().Set("Content-Type", "text/html")
	component.Render(r.Context(), w)
}

func (app *application) handlerDashboard(w http.ResponseWriter, r *http.Request) {
	component := ui.Index("Dashboard")

	w.Header().Set("Content-Type", "text/html")
	component.Render(r.Context(), w)
}

func (app *application) handlerWarehouse(w http.ResponseWriter, r *http.Request) {
	component := ui.Index("Warehouse")

	w.Header().Set("Content-Type", "text/html")
	component.Render(r.Context(), w)
}

func (app *application) handlerUser(w http.ResponseWriter, r *http.Request) {
	component := ui.Index("User")

	w.Header().Set("Content-Type", "text/html")
	component.Render(r.Context(), w)
}
