package controller

import (
	"encoding/json"
	"khushal/goMux/entity"
	"khushal/goMux/service"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type UserController interface {
	GetUsers(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)
	CreateUser(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
}

type controller struct {
	service service.UserService
}

func New(service service.UserService) UserController {
	return &controller{service: service}
}

func (controller *controller) GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(controller.service.GetUsers())
}
func (controller *controller) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err!=nil{
		json.NewEncoder(w).Encode("invalid id")
		return
	}

	json.NewEncoder(w).Encode(controller.service.GetUser(uint(id)))
}
func (controller *controller) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user entity.User
	json.NewDecoder(r.Body).Decode(&user)
	controller.service.CreateUser(user)
	json.NewEncoder(w).Encode(user)
}
func (controller *controller) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err!=nil{
		json.NewEncoder(w).Encode("invalid id")
		return
	}
	controller.service.DeleteUser(uint(id))
	json.NewEncoder(w).Encode("user deleted")
}
func (controller *controller) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user entity.User
	json.NewDecoder(r.Body).Decode(&user)
	controller.service.UpdateUser(user)
	json.NewEncoder(w).Encode(user)
}
