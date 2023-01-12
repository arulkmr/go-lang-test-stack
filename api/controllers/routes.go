package controllers

import (
	"fmt"
	"go-lang-test-stack/api/middlewares"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) InitializeRoutes(port string) {
	routes := mux.NewRouter()

	// Home Route
	routes.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	//Location Routes
	routes.HandleFunc("/location", middlewares.SetMiddlewareJSON(s.CreateLocation)).Methods("POST")
	routes.HandleFunc("/location", middlewares.SetMiddlewareJSON(s.GetLocations)).Methods("GET")
	routes.HandleFunc("/location/{id}", middlewares.SetMiddlewareJSON(s.GetLocation)).Methods("GET")
	routes.HandleFunc("/location/{id}", middlewares.SetMiddlewareJSON(s.UpdateLocation)).Methods("PUT")
	routes.HandleFunc("/location/{id}", middlewares.SetMiddlewareJSON(s.DeleteLocation)).Methods("DELETE")

	fmt.Println("Listening to port 7000")
	log.Fatal(http.ListenAndServe(port, routes))
}
