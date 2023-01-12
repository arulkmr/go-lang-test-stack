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

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("TEST - 5  CONT HOME.GO")

	responses.JSON(w, http.StatusOK, "Welcome To REST API")

}

func CreateLocation(w http.ResponseWriter, r *http.Request) {

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
	locationCreated, err := service.SaveLocation(&location)

	if err != nil {

		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%s", r.Host, r.RequestURI, locationCreated.LocationId))
	responses.JSON(w, http.StatusCreated, locationCreated)
}

func GetLocations(w http.ResponseWriter, r *http.Request) {

	fmt.Println("TEST - 5  GETLOC CONT.GO")

	locations, err := service.FindAllLocations()
	fmt.Println("TEST - 5  GETLOC CONT.GO", locations)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, locations)
}

func GetLocation(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	locId := vars["id"]

	locationGotten, err := service.FindLocationByID(locId)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, locationGotten)
}

func UpdateLocation(w http.ResponseWriter, r *http.Request) {

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

	updatedLocation, err := service.UpdateALocation(locId, &location)
	if err != nil {

		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, updatedLocation)
}

func DeleteLocation(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	locId := vars["id"]

	_, err := service.DeleteALocation(locId)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%s", locId))
	responses.JSON(w, http.StatusOK, "Location Deleted Successfully")
}
