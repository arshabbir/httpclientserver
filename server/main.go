package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func main() {

	os.Setenv("PORT", ":8080")

	port := os.Getenv("PORT")
	mux := http.NewServeMux()
	mux.HandleFunc("/request", handRequest)
	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      mux,
		Addr:         port,
	}

	log.Println("Waiting for connections.....")
	srv.ListenAndServe()
}

func handRequest(w http.ResponseWriter, r *http.Request) {
	resp := rand.Intn(1000)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("%d", resp)))
}
