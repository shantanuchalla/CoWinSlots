package server

import (
	"net/http"

	"github.com/shantanuchalla/awesomeProject/pkg/health"
)

func addRoutes(handler *http.ServeMux) {
	handler.HandleFunc("/health-check", health.GetHealthChecker())
}
