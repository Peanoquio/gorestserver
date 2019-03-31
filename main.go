package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Peanoquio/gorestserver/router"
	"github.com/Peanoquio/gorestserver/router/routes"
)

func main() {
	const PORT = 8080
	// map that will contain the API routes
	var apiRoutes = make(map[string]routes.APIRouteInterface)
	apiRoutes["/api/test"] = &routes.APIRoute{}

	routeMgr := router.NewRouterManager("/v1", apiRoutes)
	routeMgr.ShowRoutes()

	// HTTP server
	if err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), routeMgr.GetRouter()); err != nil {
		log.Fatalf("HTTP server fatal error:%s \n", err.Error())
	}
}
