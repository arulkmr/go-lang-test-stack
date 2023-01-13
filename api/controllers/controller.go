package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"go-lang-test-stack/api/models"
	"go-lang-test-stack/api/responses"
	"go-lang-test-stack/api/service"

	"github.com/gorilla/mux"
)

type Service struct {
	service.ServiceLogic
}

type ControllerInterface interface {
	Home(w http.ResponseWriter, r *http.Request)
	CreateLocation(w http.ResponseWriter, r *http.Request)
	GetLocations(w http.ResponseWriter, r *http.Request)
	GetLocation(w http.ResponseWriter, r *http.Request)
	UpdateLocation(w http.ResponseWriter, r *http.Request)
	DeleteLocation(w http.ResponseWriter, r *http.Request)
}

func (s *Service) Home(w http.ResponseWriter, r *http.Request) {

	responses.JSON(w, http.StatusOK, "Welcome To REST API")

}

func (s *Service) CreateLocation(w http.ResponseWriter, r *http.Request) {

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
	locationCreated, err := s.SaveLocation(&location)

	if err != nil {

		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%s", r.Host, r.RequestURI, locationCreated.LocationId))
	responses.JSON(w, http.StatusCreated, locationCreated)
}

func (s *Service) GetLocations(w http.ResponseWriter, r *http.Request) {

	locations, err := s.FindAllLocations()

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, locations)
}

func (s *Service) GetLocation(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	locId := vars["id"]

	locationGotten, err := s.FindLocationByID(locId)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, locationGotten)
}

func (s *Service) UpdateLocation(w http.ResponseWriter, r *http.Request) {

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

	updatedLocation, err := s.UpdateALocation(locId, location)
	if err != nil {

		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, updatedLocation)
}

func (s *Service) DeleteLocation(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	locId := vars["id"]

	_, err := s.DeleteALocation(locId)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", locId)
	responses.JSON(w, http.StatusOK, "Location Deleted Successfully")
}
