package server

import "net/http"

func InitHandler() *http.ServeMux {
	handler := http.NewServeMux()
	addRoutes(handler)

	return handler
}
