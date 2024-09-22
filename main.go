package main

import (
	"log"
	"net/http"
	"os"
)

func main() {

	port := os.Getenv("PORT")
	host := os.Getenv("HOST")
	log.Printf("Listening on %s:%s", host, port)
	http.HandleFunc("/gorik", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello GORIK"))
	})
	panic(http.ListenAndServe(host+":"+port, nil))

}
