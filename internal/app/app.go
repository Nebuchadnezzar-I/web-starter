package app

import (
	"net/http"
	"log/slog"
)

type Application struct {
	Router *http.ServeMux
	Logger *slog.Logger
}
