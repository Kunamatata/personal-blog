package users

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	types "github.com/kunamatata/blog/types"
	"golang.org/x/crypto/bcrypt"
)

type UserController interface {
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
}

//Controller is the controller for posts
type Controller struct {
	userRepository *UserRepo
}

func NewController(userRepository *UserRepo) *Controller {
	return &Controller{
		userRepository: userRepository,
	}
}

func (c *Controller) Create(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		fmt.Println("Could not decode request body into post")
	}

	createdUser, err := c.userRepository.Create(&user)
	if err != nil {
		log.Println("Could not create user")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(types.APIResponse{
			Status: "error",
			Data:   nil,
		})
		return
	}

	json.NewEncoder(w).Encode(types.APIResponse{
		Status: "success",
		Data:   createdUser,
	})
}

func (c *Controller) GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := c.userRepository.GetAll()
	if err != nil {
		log.Println("Could not get all users")
	}

	json.NewEncoder(w).Encode(types.APIResponse{
		Status: "success",
		Data:   users,
	})
}

func (c *Controller) Get(w http.ResponseWriter, r *http.Request) {
	slug := mux.Vars(r)["slug"]
	user, err := c.userRepository.Get(slug)
	if err != nil {
		log.Println("Could not get user")
	}

	json.NewEncoder(w).Encode(types.APIResponse{
		Status: "success",
		Data:   user,
	})
}

func (c *Controller) Login(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		fmt.Println("Could not decode request body into post")
	}

	gotUser, err := c.userRepository.GetByEmail(user.Email)

	if gotUser == nil || err != nil {
		log.Println("Could not get user")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(types.APIResponse{
			Status: "error",
			Data:   "",
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(gotUser.Password), []byte(user.Password))
	if err != nil {
		log.Println("Passwords do not match")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(types.APIResponse{
		Status: "success",
		Data:   nil,
	})
}
