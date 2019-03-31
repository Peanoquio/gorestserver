package routes

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// APIRoute contains the API routes
type APIRoute struct {
	router *chi.Mux
}

// GenerateRoutes generates the API routes based on the path then maps it to the handler functions
func (apiRoute *APIRoute) GenerateRoutes() *chi.Mux {
	apiRoute.router = chi.NewRouter()
	apiRoute.router.Get("/{testParam}", requestTest)
	apiRoute.router.Post("/{testParam}", requestTest)
	apiRoute.router.Put("/{testParam}", requestTest)
	apiRoute.router.Patch("/{testParam}", requestTest)
	apiRoute.router.Delete("/{testParam}", requestTest)
	return apiRoute.router
}

// RequestTest is handler for the /{testParam} API endpoint
func requestTest(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "testParam")
	// the JSON payload for the request
	jsonPayload := map[string]interface{}{
		"message":   "Hello world! The server got your message.",
		"testParam": param,
	}
	render.JSON(w, r, jsonPayload)
}
