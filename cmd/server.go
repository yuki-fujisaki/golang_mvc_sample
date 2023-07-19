package main

import (
	"fmt"
	"golang_mvc_sample/cmd/router"
	"net/http"
)

func main() {
	r := router.Routes()

	fmt.Println("Server running at http://localhost:3000")
	http.ListenAndServe(":3000", r)
}
