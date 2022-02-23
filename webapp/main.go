package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MijPeter/gotutorial/tree/main/webapp/handlers"
)

func main() {
	http.HandleFunc("/view/", handlers.MakeHandler(handlers.ViewHandler))
	http.HandleFunc("/edit/", handlers.MakeHandler(handlers.EditHandler))
	http.HandleFunc("/save/", handlers.MakeHandler(handlers.SaveHandler))

	fmt.Println("Starting server on port 8080.")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
