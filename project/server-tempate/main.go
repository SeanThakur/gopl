package main

import (
	"fmt"
	"net/http"

	"github.com/SeanThakur/gopl/project/server-tempate/handlers"
)

const PORT = ":8080"

func main() {
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/about", handlers.AboutHandler)

	fmt.Println("Starting port at ", PORT)
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		fmt.Println("Server crashed", err.Error())
		return
	}
}
