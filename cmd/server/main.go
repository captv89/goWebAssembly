package main

import (
	"fmt"
	"net/http"
)

func main() {
	err := http.ListenAndServe(":8081", http.FileServer(http.Dir("/Users/Scorpion/Go-Projects/goWebAssembly/assets")))
	if err != nil {
		fmt.Println("Error starting server: ", err)
		return
	}
}
