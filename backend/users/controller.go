package users

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kunamatata/blog/cookiesession"
	types "github.com/kunamatata/blog/types"
	"golang.org/x/crypto/bcrypt"
)

type UserController interface {
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
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
	session, _ := cookiesession.Store.Get(r, "session")
	userID := session.Values["userID"]
	log.Println("UserID: ", userID)

	slug := mux.Vars(r)["userID"]
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
	//TODO: implement cookie session here
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		fmt.Println("Could not decode request body into post")
		JSONError(w, types.APIResponse{
			Status: "error",
			Data:   nil,
		}, http.StatusInternalServerError)
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
		message := "Username or password do not match"
		JSONError(w, types.APIResponse{
			Status:  "error",
			Message: &message,
		}, http.StatusUnauthorized)
		return
	}

	session, _ := cookiesession.Store.Get(r, "session")
	session.Values["userID"] = gotUser.ID
	session.Save(r, w)
	log.Println("Successfully Logged In")

	json.NewEncoder(w).Encode(types.APIResponse{
		Status: "success",
		Data:   nil,
	})
}

func (c *Controller) Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := cookiesession.Store.Get(r, "session")
	session.Values["userID"] = ""
	session.Options.MaxAge = -1 // delete cookie
	session.Save(r, w)
	log.Println("Successfully Logged Out")
}

func JSONError(w http.ResponseWriter, err interface{}, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(err)
}
