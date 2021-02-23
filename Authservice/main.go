package main

import (
	"fmt"
	"modules/repo"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	//Connect to database
	repository := repo.ConnectDB()
	fmt.Println(repository)
	r := chi.NewRouter()
	//Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(60 * time.Second))

	//Routes
	test :=  func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "HEllo")
	}
	r.Route("/auth", func(r chi.Router) {
		r.MethodFunc("GET","/", test)
	})



	http.ListenAndServe(":8090", r)
}