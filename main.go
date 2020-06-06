package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"./handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	// Create the handlers
	ph := handlers.NewProducts(l)

	// Create a new serve mux and register the handlers
	sm := http.NewServeMux()
	sm.Handle("/", ph)

	// Create a new server
	s := &http.Server{
		Addr:         ":9090",           // Configure the bind address
		Handler:      sm,                // Set the default handler
		ErrorLog:     l,                 // Set the logger for the server
		IdleTimeout:  120 * time.Second, // Max time for connections using TCP Keep-Alive
		ReadTimeout:  5 * time.Second,   // Max time to read request from the client
		WriteTimeout: 10 * time.Second,  // Max time to write response to the client
	}

	// Start the server
	go func() {
		l.Println("Starting server on port 9090...")

		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	// Trap the sigterm or interrupt and gracefully shutdown the server
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Recieved terminated, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
