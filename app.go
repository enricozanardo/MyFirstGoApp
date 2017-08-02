package main

import (
	"log"
	"net/http"
	"os"
	"github.com/onezerobinary/MyGoApp/handler"
	//for extracting service credentials from VCAP_SERVICES
	//"github.com/cloudfoundry-community/go-cfenv"
)

const (
	DEFAULT_PORT = "8080"
)

func hill(val1 int, val2 int) (int, int){
	vet1 := []int{13, 3}
	vet2 := []int{17, 4}

	res1 :=  (vet1[0] * val1 + vet1[1] * val2) % 30
	res2 :=  (vet2[0] * val1 + vet2[1] * val2) % 30

	println("res1: ", res1)
	println("res2: ", res2)

	return res1, res2
}


func main() {
	var port string
	if port = os.Getenv("PORT"); len(port) == 0 {
		port = DEFAULT_PORT
	}

	val1 , val2 := hill(19,14)

	println("Hill [12 4]:", val1, val2)

	// Think about that declaration
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", handler.HomeHandler)


	log.Printf("Starting app on port %+v\n", port)
	http.ListenAndServe(":"+port, nil)
}
