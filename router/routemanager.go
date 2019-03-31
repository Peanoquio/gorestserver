package router

import (
	"log"
	"net/http"
	"time"

	"github.com/Peanoquio/gorestserver/router/routes"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

// RouteManagerInterface is the interface that needs to be implemented
type RouteManagerInterface interface {
	initRoutes(mainPath string, apiRoutes map[string]routes.APIRouteInterface) *chi.Mux
	ShowRoutes()
	GetRouter() *chi.Mux
}

// RouteManager is the struct class for using middlewares and handling request routes
type RouteManager struct {
	router *chi.Mux
}

// initRoutes creates the router that will use middlewares and handle API calls
func (routeMgr *RouteManager) initRoutes(mainPath string, apiRoutes map[string]routes.APIRouteInterface) *chi.Mux {
	routeMgr.router = chi.NewRouter()
	// the middleware for the HTTP requests
	routeMgr.router.Use(
		// set content-type to be JSON
		render.SetContentType(render.ContentTypeJSON),
		// Injects a request ID into the context of each request
		middleware.RequestID,
		// Sets a http.Request's RemoteAddr to either X-Forwarded-For or X-Real-IP
		middleware.RealIP,
		// Logs the start and end of each request with the elapsed processing time
		middleware.Logger,
		// Redirect slashes on routing paths
		middleware.RedirectSlashes,
		// Gracefully absorb panics and prints the stack trace
		middleware.Recoverer,
	)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	routeMgr.router.Use(middleware.Timeout(60 * time.Second))

	routeMgr.router.Route(mainPath, func(r chi.Router) {
		for apiPath, apiRoute := range apiRoutes {
			router := apiRoute.GenerateRoutes()
			// mount the api path and route handler
			r.Mount(apiPath, router)
		} // end loop
	})

	return routeMgr.router
}

// ShowRoutes displays all the generated routes
func (routeMgr *RouteManager) ShowRoutes() {
	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("ShowRoutes method:%s route:%s \n", method, route)
		return nil
	}
	if err := chi.Walk(routeMgr.router, walkFunc); err != nil {
		log.Panicf("ShowRoutes error:%s \n", err.Error())
	}
}

// GetRouter returns the router instance
func (routeMgr *RouteManager) GetRouter() *chi.Mux {
	return routeMgr.router
}

// NewRouterManager creates a new instance of the RouteManager class (like a factory method)
func NewRouterManager(mainPath string, apiRoutes map[string]routes.APIRouteInterface) RouteManagerInterface {
	var routeMgrInterface RouteManagerInterface
	routeMgr := &RouteManager{router: nil}
	routeMgr.initRoutes(mainPath, apiRoutes)
	routeMgrInterface = routeMgr
	return routeMgrInterface
}
