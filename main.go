package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Bernar11296/rest/app"
	"github.com/Bernar11296/rest/controllers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.Use(app.JwtAuthentication)
	router.HandleFunc("api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("api/user/login", controllers.Authenticate).Methods("POST")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print(err)
	}
}
