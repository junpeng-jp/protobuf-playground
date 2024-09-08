package api

import (
	"net/http"

	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

type Service struct {
	registry  map[string][]byte
	typeCache protoregistry.Types
	kvstore   map[string]protoreflect.Message
	mux       *http.ServeMux
}

func NewService() http.Handler {
	mux := http.NewServeMux()
	mux = addRoutes(mux)

	registry := NewRegistry()
	typeCache := protoregistry.Types{}
	kvstore := make(map[string]protoreflect.Message)

	service := &Service{
		registry,
		typeCache,
		kvstore,
		mux,
	}

	return service.mux
}

func NewRegistry() map[string][]byte {
	registry := make(map[string][]byte)

	return registry
}
