package weather

import (
	"net/http"
)

func Router() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", handle.base)
	return mux
}
