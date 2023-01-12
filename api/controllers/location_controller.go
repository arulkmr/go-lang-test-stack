package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"go-lang-test-stack/api/models"
	"go-lang-test-stack/api/responses"

	"github.com/gorilla/mux"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To REST API")

}

func (server *Server) CreateLocation(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	location := models.Location{}
	err = json.Unmarshal(body, &location)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	locationCreated, err := location.SaveLocation(server.DB)

	if err != nil {

		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%s", r.Host, r.RequestURI, locationCreated.LocationId))
	responses.JSON(w, http.StatusCreated, locationCreated)
}

func (server *Server) GetLocations(w http.ResponseWriter, r *http.Request) {

	location := models.Location{}

	locations, err := location.FindAllLocations(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, locations)
}

func (server *Server) GetLocation(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	locId := vars["id"]

	location := models.Location{}
	locationGotten, err := location.FindLocationByID(server.DB, locId)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, locationGotten)
}

func (server *Server) UpdateLocation(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	locId := vars["id"]

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	location := models.Location{}
	err = json.Unmarshal(body, &location)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	updatedLocation, err := location.UpdateALocation(server.DB, locId)
	if err != nil {

		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, updatedLocation)
}

func (server *Server) DeleteLocation(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	location := models.Location{}

	locId := vars["id"]

	_, err := location.DeleteALocation(server.DB, locId)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%s", locId))
	responses.JSON(w, http.StatusOK, "Location Deleted Successfully")
}
