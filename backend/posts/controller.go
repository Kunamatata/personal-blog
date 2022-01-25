package posts

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kunamatata/blog/types"
)

//PostController represents the interface for the post controller
type PostController interface {
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
}

//Controller is the controller for posts
type Controller struct {
	postRepository *PostRepo
}

//NewController creates a new instance of the post controller
func NewController(postRepository *PostRepo) *Controller {
	return &Controller{
		postRepository: postRepository,
	}
}

//Create handles the request to create a new post
func (c *Controller) Create(w http.ResponseWriter, r *http.Request) {
	var post Post
	err := json.NewDecoder(r.Body).Decode(&post)

	if err != nil {
		fmt.Println("Could not decode request body into post")
	}

	createdPost, err := c.postRepository.Create(&post)
	if err != nil {
		log.Println("Could not create post")
	}

	json.NewEncoder(w).Encode(types.APIResponse{
		Status: "success",
		Data:   createdPost,
	})
}

//Get handles the request to get a post
func (c *Controller) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	slug := vars["postID"]
	apiResponse := types.APIResponse{}

	post, err := c.postRepository.Get(slug)

	if err != nil {
		log.Println("Could not retrieve post")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if post != nil {
		apiResponse = types.APIResponse{
			Status: "success",
			Data:   post,
		}
		w.WriteHeader(http.StatusOK)
	} else {
		var message = "Post not found"
		apiResponse = types.APIResponse{
			Status:  "error",
			Message: &message,
		}
		w.WriteHeader(http.StatusNotFound)
	}
	json.NewEncoder(w).Encode(apiResponse)
}

//Delete handles the request to delete a post
func (c *Controller) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	slug := vars["postID"]

	err := c.postRepository.Delete(slug)
	if err != nil {
		log.Println("Could not delete post")
		json.NewEncoder(w).Encode(
			types.APIResponse{
				Status: "error",
			},
		)
	}

	json.NewEncoder(w).Encode(
		types.APIResponse{
			Status: "success",
		},
	)
}

//GetAll handles the request to get all posts
func (c *Controller) GetAll(w http.ResponseWriter, r *http.Request) {
	posts, err := c.postRepository.GetAll()
	if err != nil {
		message := "Could not retrieve posts"
		json.NewEncoder(w).Encode(types.APIResponse{
			Status:  "error",
			Message: &message,
		})
	}

	json.NewEncoder(w).Encode(types.APIResponse{
		Status: "success",
		Data:   posts,
	})
}

//Update handles the request to update a post
func (c *Controller) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	slug := vars["postID"]

	var post Post
	err := json.NewDecoder(r.Body).Decode(&post)

	if err != nil {
		fmt.Println("Could not decode request body into post")
	}

	post.Slug = slug
	updatedPost, err := c.postRepository.Update(&post)
	if err != nil {
		log.Println("Could not update post")
	}

	json.NewEncoder(w).Encode(types.APIResponse{
		Status: "success",
		Data:   updatedPost,
	})

}
