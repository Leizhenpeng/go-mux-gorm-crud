package main

import (
	"fmt"
	"leizhenpeng/go-mux-gorm-crud/internal/model"
	"leizhenpeng/go-mux-gorm-crud/internal/router"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	//load env from ./deploy/.env file
	err := godotenv.Load("./deploy/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
func main() {

	//log os.getenv("MYSQL_URL")
	model.InitDB()
	r := mux.NewRouter()
	router.RegisterPlaceRoutes(r)
	http.Handle("/", r)
	fmt.Println("Listening on port 8080", "http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
