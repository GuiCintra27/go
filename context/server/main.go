package main

import (
	"log"
	"net/http"
	"time"
)

func main()  {
	http.HandleFunc("/", home)
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log.Println("request started")
	defer log.Println("request ended")

	select{
		case <- time.After(time.Second * 10):
			log.Println("request processed")
			w.Write([]byte("request processed with success"))

		case <- ctx.Done():
			log.Println("request canceled")
			http.Error(w, "request canceled", http.StatusRequestTimeout)
	}
}