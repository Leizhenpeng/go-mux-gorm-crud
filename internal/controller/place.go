package controller

import (
	"encoding/json"
	"fmt"
	"leizhenpeng/go-mux-gorm-crud/internal/model"
	"leizhenpeng/go-mux-gorm-crud/internal/util"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAllPlaces(w http.ResponseWriter, r *http.Request) {
	places, err := model.GetAllPlaces()
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, places)
	}
	respondWithJSON(w, http.StatusOK, places)
}
func GetPlacesById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, err)
		return
	}
	place, err := model.GetPlaceByID(idInt)
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, err)
		return
	}
	respondWithJSON(w, http.StatusOK, place)
}

func CreatePlace(w http.ResponseWriter, r *http.Request) {
	newPlace := model.Place{}
	fmt.Println(r.Body)
	err := util.ParseBody(r, &newPlace)
	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, err)
		return
	}
	fmt.Println(newPlace)
	place, err := model.CreatePlace(newPlace)
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, err)
		return
	}
	respondWithJSON(w, http.StatusCreated, place)
}

func DeletePlace(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, err)
		return
	}
	err2 := model.DeletePlace(idInt)
	if err2 != nil {
		respondWithJSON(w, http.StatusInternalServerError, err)
		return
	}
	respondWithJSON(w, http.StatusOK, "success delete")

}

func UpdatePlace(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, err)
		return
	}
	newPlace := model.Place{}
	err = util.ParseBody(r, &newPlace)
	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, err)
		return
	}
	place, err := model.UpdatePlace(idInt, newPlace)
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, err)
		return
	}
	respondWithJSON(w, http.StatusOK, place)
}

func respondWithJSON(w http.ResponseWriter, ok int, res interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(ok)
	json.NewEncoder(w).Encode(res)

}
