package main

import (
	"khushal/goMux/controller"
	"khushal/goMux/repository"
	"khushal/goMux/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	db := repository.InitDB()

	service := service.New(db)
	userController:= controller.New(service)
	loginController := controller.NewLoginController(service)
	server := mux.NewRouter()
	server.HandleFunc("/login", loginController.Login).Methods("GET")
	
	server.HandleFunc("/users", userController.GetUsers).Methods("GET")
	server.HandleFunc("/users", userController.CreateUser).Methods("POST")
	server.HandleFunc("/users/{id}", userController.GetUser).Methods("GET")
	server.HandleFunc("/users/{id}", userController.UpdateUser).Methods("PUT")
	server.HandleFunc("/users/{id}", userController.DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9000", server))

}
