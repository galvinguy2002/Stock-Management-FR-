package main

import (
	"fmt"
	"net/http"
)

func main() {

	fmt.Println("Setting Up Routes For Server...")
	database.Connect() //sample connecting dummy variable
	api.SetupRoutes()  //sample connecting dummy variable
	fmt.Println("Starting server at 8080")
	http.ListenAndServe(":8080", nil)
}
