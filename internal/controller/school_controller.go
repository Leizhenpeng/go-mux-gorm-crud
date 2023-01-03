package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAllSchools(w http.ResponseWriter, r *http.Request) {
	var schools []model.School
	err := model.GetAllSchools(&schools)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, schools)
}

func GetSchoolById(w http.ResponseWriter, r *http.Request) {
	var school model.School
	schoolId := mux.Vars(r)["id"]
	err := model.GetSchoolById(&school, schoolId)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, school)
}

func CreateSchool(w http.ResponseWriter, r *http.Request) {
	var school model.School
	err := json.NewDecoder(r.Body).Decode(&school)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	err = model.CreateSchool(&school)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, school)
}

func UpdateSchool(w http.ResponseWriter, r *http.Request) {
	var school model.School
	schoolId := mux.Vars(r)["id"]
	err := model.GetSchoolById(&school, schoolId)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&school)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	err = model.UpdateSchool(&school, schoolId)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, school)
}

func DeleteSchool(w http.ResponseWriter, r *http.Request) {
	var school model.School
	schoolId := mux.Vars(r)["id"]
	err := model.DeleteSchool(&school, schoolId)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, school)
}
