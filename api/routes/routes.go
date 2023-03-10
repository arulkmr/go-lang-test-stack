package routes

import (
	"fmt"
	c "go-lang-test-stack/api/controllers"
	"go-lang-test-stack/api/middlewares"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func InitializeRoutes(port string) {
	routes := mux.NewRouter()

	// Home Route
	routes.HandleFunc("/", middlewares.SetMiddlewareJSON(c.Home)).Methods("GET")

	//Location Routes
	routes.HandleFunc("/location", middlewares.SetMiddlewareJSON(c.CreateLocation)).Methods("POST")
	routes.HandleFunc("/location", middlewares.SetMiddlewareJSON(c.GetLocations)).Methods("GET")
	routes.HandleFunc("/location/{id}", middlewares.SetMiddlewareJSON(c.GetLocation)).Methods("GET")
	routes.HandleFunc("/location/{id}", middlewares.SetMiddlewareJSON(c.UpdateLocation)).Methods("PUT")
	routes.HandleFunc("/location/{id}", middlewares.SetMiddlewareJSON(c.DeleteLocation)).Methods("DELETE")

	//Loncation Query Handling
	routes.HandleFunc("/location/search", middlewares.SetMiddlewareJSON(c.LocationQuery)).Methods("POST")

	fmt.Println("Listening to port 7000")
	log.Fatal(http.ListenAndServe(port, routes))
}
