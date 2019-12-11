package main

import (
	"fmt"
	"htttpServer/models"
	"htttpServer/routes"
	"htttpServer/utils"
	"log"
	"net/http"
)

const PORT = ":8080"

func main() {
	models.TestConnection()
	fmt.Printf("Listening Port %s\n", PORT)
	utils.LoadTemplates("views/*.html")
	r := routes.NewRouter()
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(PORT, nil))
}
