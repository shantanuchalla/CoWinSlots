package health

import "net/http"

func GetHealthChecker() func(http.ResponseWriter, *http.Request) {
	return func(http.ResponseWriter, *http.Request) {

	}
}
