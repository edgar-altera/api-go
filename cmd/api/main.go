package main

import (
	"fmt"
	"net/http"
	"github.com/edgar-altera/api-go/internal/server"
)

func main() {

	fmt.Println("Init App")
	
	r := server.NewRouter()
 
	http.ListenAndServe(":8080", r)
}