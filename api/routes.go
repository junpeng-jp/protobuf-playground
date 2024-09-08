package api

import (
	"net/http"

	"github.com/junpeng.ong/protobuf-playground/api/handlers"
)

func addRoutes(mux *http.ServeMux) *http.ServeMux {

	// Routes that demonstrate static unmarshaling of nested protobuf
	mux.HandleFunc("POST /static/location", handlers.HandleLocation)
	mux.HandleFunc("POST /static/error", handlers.HandleError)
	mux.HandleFunc("POST /static/container", handlers.HandleContainer)

	// Routes that demonstrate dynamic unmarshaling of nested protobuf
	mux.HandleFunc("POST /dynamic/container", handlers.HandleMessageDynamically)

	return mux
}
