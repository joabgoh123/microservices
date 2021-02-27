package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/joabgoh123/common/customerrors"
)

func main() {
	//Connect to database
	// repository := repo.ConnectDB()
	// fmt.Println(repository)
	r := chi.NewRouter()
	//Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(60 * time.Second))

	//Routes
	test :=  func(w http.ResponseWriter, r *http.Request) error{
		fmt.Fprint(w, "HEllo")
		return errors.New("AN error occured")
	}
	//Append the auth route
	r.Route("/auth", func(r chi.Router) {
		r.Method("GET","/test", customerrors.BaseErrorHandler(test))
	})


	http.ListenAndServe(":8090", r)
}