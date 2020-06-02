package main

import (
	"log"
	"net/http"

	"github.com/gozaddy/AuthProject/database"

	"github.com/gorilla/mux"
	"github.com/gozaddy/AuthProject/controllers"
)

func init() {

	database.Connect()
	database.InitCache()

}
func main() {
	router := mux.NewRouter()

	defer database.Cache.Close()
	router.HandleFunc("/", controllers.IndexPage).Methods("GET")
	router.HandleFunc("/login", controllers.LoginPage).Methods("GET")
	router.HandleFunc("/register", controllers.SignInPage).Methods("GET")

	router.HandleFunc("/login", controllers.Login).Methods("POST")
	router.HandleFunc("/register", controllers.Register).Methods("POST")
	router.HandleFunc("/logout", controllers.LogOut)

	log.Fatal(http.ListenAndServe(":8080", router))
}
