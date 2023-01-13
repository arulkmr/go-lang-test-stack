package routes

import (
	"fmt"
	"go-lang-test-stack/api/controllers"
	"go-lang-test-stack/api/middlewares"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type RoutesInterface interface {
	InitializeRoutes()
}
type Routes struct {
	controllers.ControllerInterface
	router *mux.Router
}

func (r *Routes) InitializeRoutes() {

	r.router = mux.NewRouter()

	// Home Route
	r.router.HandleFunc("/", middlewares.SetMiddlewareJSON(r.Home)).Methods("GET")

	//Location Routes
	r.router.HandleFunc("/location", middlewares.SetMiddlewareJSON(r.CreateLocation)).Methods("POST")
	r.router.HandleFunc("/location", middlewares.SetMiddlewareJSON(r.GetLocations)).Methods("GET")
	r.router.HandleFunc("/location/{id}", middlewares.SetMiddlewareJSON(r.GetLocation)).Methods("GET")
	r.router.HandleFunc("/location/{id}", middlewares.SetMiddlewareJSON(r.UpdateLocation)).Methods("PUT")
	r.router.HandleFunc("/location/{id}", middlewares.SetMiddlewareJSON(r.DeleteLocation)).Methods("DELETE")

	fmt.Println("Listening to port 7000")
	log.Fatal(http.ListenAndServe(os.Getenv("SERVER_PORT"), r.router))
}
