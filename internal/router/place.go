package router

import (
	"leizhenpeng/go-mux-gorm-crud/internal/controller"

	"github.com/gorilla/mux"
)

func RegisterPlaceRoutes(r *mux.Router) {
	r.HandleFunc("/place", controller.CreatePlace).Methods("POST")
	r.HandleFunc("/places", controller.GetAllPlaces).Methods("GET")
	r.HandleFunc("/place/{id}", controller.UpdatePlace).Methods("PUT")
	r.HandleFunc("/place/{id}", controller.DeletePlace).Methods("DELETE")
}
