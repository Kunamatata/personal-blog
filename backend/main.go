package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kunamatata/blog/database"
	"github.com/kunamatata/blog/posts"
	"github.com/kunamatata/blog/users"
)

func corsHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}

func main() {

	db := database.NewConnection()

	postRepo := posts.NewPostRepo(db)
	postController := posts.NewController(postRepo)
	userRepo := users.NewUserRepo(db)
	userController := users.NewController(userRepo)

	router := mux.NewRouter()

	router.Use(corsHandler)
	router.HandleFunc("/posts/{postID}", postController.Delete).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/posts/{postID}", postController.Get).Methods("GET", "OPTIONS")
	router.HandleFunc("/posts", postController.Create).Methods("POST", "OPTIONS")
	router.HandleFunc("/posts", postController.GetAll).Methods("GET", "OPTIONS")

	router.HandleFunc("/users", userController.GetAll).Methods("GET", "OPTIONS")
	router.HandleFunc("/users/{userID}", userController.Get).Methods("GET", "OPTIONS")
	router.HandleFunc("/users", userController.Create).Methods("POST", "OPTIONS")
	router.HandleFunc("/login", userController.Login).Methods("POST", "OPTIONS")
	router.HandleFunc("/logout", userController.Logout).Methods("POST", "OPTIONS")

	router.Use(mux.CORSMethodMiddleware(router))

	log.Fatal(http.ListenAndServe(":8080", router))
}
