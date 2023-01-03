package router

import (
	"leizhenpeng/go-mux-gorm-crud/internal/controller"

	"github.com/gorilla/mux"
)

var RegisterSchoolRoutes = func(router *mux.Router) {
	router.HandleFunc("/schools", controller.GetAllSchools).Methods("GET")
	router.HandleFunc("/schools/{id}", controller.GetSchoolById).Methods("GET")
	router.HandleFunc("/schools", controller.CreateSchool).Methods("POST")
	router.HandleFunc("/schools/{id}", controller.UpdateSchool).Methods("PUT")
	router.HandleFunc("/schools/{id}", controller.DeleteSchool).Methods("DELETE")
}
