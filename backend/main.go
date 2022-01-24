package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kunamatata/blog/database"
	"github.com/kunamatata/blog/posts"
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

	router := mux.NewRouter()

	router.Use(corsHandler)
	router.HandleFunc("/posts/{postID}", postController.Delete).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/posts/{postID}", postController.Get).Methods("GET", "OPTIONS")
	router.HandleFunc("/posts", postController.Create).Methods("POST", "OPTIONS")
	router.HandleFunc("/posts", postController.GetAll).Methods("GET", "OPTIONS")

	router.Use(mux.CORSMethodMiddleware(router))

	log.Fatal(http.ListenAndServe(":8080", router))
}
