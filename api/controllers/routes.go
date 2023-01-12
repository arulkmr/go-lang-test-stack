package controllers

import "go-lang-test-stack/api/middlewares"

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	//Location Routes
	s.Router.HandleFunc("/location", middlewares.SetMiddlewareJSON(s.CreateLocation)).Methods("POST")
	s.Router.HandleFunc("/location", middlewares.SetMiddlewareJSON(s.GetLocations)).Methods("GET")
	s.Router.HandleFunc("/location/{id}", middlewares.SetMiddlewareJSON(s.GetLocation)).Methods("GET")
	s.Router.HandleFunc("/location/{id}", middlewares.SetMiddlewareJSON(s.UpdateLocation)).Methods("PUT")
	s.Router.HandleFunc("/location/{id}", middlewares.SetMiddlewareJSON(s.DeleteLocation)).Methods("DELETE")

}
