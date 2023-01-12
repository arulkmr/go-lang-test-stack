package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"go-lang-test-stack/api/db"
	"go-lang-test-stack/api/models"
	"go-lang-test-stack/api/responses"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
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

	locationCreated, err := location.SaveLocation(db.DB.Db)

	if err != nil {

		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%s", r.Host, r.RequestURI, locationCreated.LocationId))
	responses.JSON(w, http.StatusCreated, locationCreated)
}

func GetLocations(w http.ResponseWriter, r *http.Request) {

	location := models.Location{}

	locations, err := location.FindAllLocations(db.DB.Db)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, locations)
}

func GetLocation(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	locId := vars["id"]

	location := models.Location{}
	locationGotten, err := location.FindLocationByID(db.DB.Db, locId)
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

	updatedLocation, err := location.UpdateALocation(db.DB.Db, locId)
	if err != nil {

		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, updatedLocation)
}

func DeleteLocation(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	location := models.Location{}

	locId := vars["id"]

	_, err := location.DeleteALocation(db.DB.Db, locId)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%s", locId))
	responses.JSON(w, http.StatusOK, "Location Deleted Successfully")
}
