package main

import (
	"fmt"
	"golang_mvc_sample/pkg/controller"
	"net/http"
)

func main() {
	r := controller.UserRoutes()

	fmt.Println("Server running at http://localhost:3000")
	http.ListenAndServe(":3000", r)
}
