package main

import (
	"blog/Post/controller"
	"blog/Post/repository"
	"blog/Post/usecase"
	"blog/config"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// r := mux.NewRouter()
	// check1 := r.PathPrefix("/").Subrouter()
	// check1.Use(loggingMiddleware)
	// r.Use(loggingMiddleware)
	// check2 := r.PathPrefix("/check").Subrouter()
	// check2.Use(controller.IsAunthenticate)
	// http.Handle("/", r)
	// check2.HandleFunc("/", controller.HandlerIndex)
	// r.HandleFunc("/detail/{id}", controller.HandlerDetail).Methods("GET")
	// check2.HandleFunc("/insert", controller.HandlerInsert).Methods("POST")
	// r.HandleFunc("/update/{id}", controller.HandlerUpdate).Methods("PUT")
	// r.HandleFunc("/delete/{id}", controller.HandlerDelete).Methods("DELETE")
	// r.HandleFunc("/login", controller.HandlerLogin).Methods("POST")
	// r.HandleFunc("/register", controller.HandlerRegister).Methods("POST")
	db := config.Connect()

	postrepo := repository.NewPostRepository(db)
	postusecase := usecase.NewPostUsecase(postrepo)
	controller.PostRouter(postusecase)

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Printf("method=%s uri=%s\n body=%s\n", r.Method, r.RequestURI, r.Body)

		next.ServeHTTP(w, r)
	})
}
