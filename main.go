package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	mux "github.com/gorilla/mux"
)

type coffee struct {
	Name      string  `json: "name"`
	Price     float64 `json: "price"`
	Available bool    `json: "available"`
}

func main() {

	log.Println("Application started")

	r := mux.NewRouter()
	r.HandleFunc("/", CoffeeHandler).Methods(http.MethodGet)

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	http.Handle("/", r)

	sigChannel := make(chan os.Signal)
	signal.Notify(sigChannel, os.Interrupt)
	signal.Notify(sigChannel, os.Kill)

	sig := <-sigChannel
	log.Panicln("Receieved termination call", sig)
	//if service has exited log it
	log.Println("Service cosmic-coffee-api has exited")
	tc, _ := context.WithTimeout(context.Background(), 3*time.Second)
	srv.Shutdown(tc)
}

func CoffeeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	cl := CoffeeList

	if len(cl) < 1 {
		//make sure you use the correct header for your response
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	json.NewEncoder(w).Encode(cl)
}

var CoffeeList = []*coffee{
	{
		Name:      "Established Belfast",
		Price:     13.99,
		Available: true,
	},
	{
		Name:      "Established Belfast",
		Price:     9.99,
		Available: false,
	},
}
