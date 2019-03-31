package routes

import "github.com/go-chi/chi"

// APIRouteInterface serves the interface that needs to be implemented
type APIRouteInterface interface {
	GenerateRoutes() *chi.Mux
}
