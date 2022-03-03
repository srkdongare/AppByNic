package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/srkdongare/AppByNic/handlers"
)

func main() {

	//Gorilla NOT implemented yet
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewProducts(l)

	//Adding Handler
	sm := http.NewServeMux()
	sm.Handle("/", hh)

	//webserver configuration
	s := &http.Server{
		Addr:         ":33333",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	//start the server
	//starting the server
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	//wait for resource removal
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)
	sig := <-sigChan
	l.Println("Interupted...", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)

}
