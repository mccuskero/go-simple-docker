package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	// Create Server and configure the handler
	r := mux.NewRouter()

	r.HandleFunc("/", handler)

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Configure Logging
	LOG_FILE_LOCATION := os.Getenv("LOG_FILE_LOCATION")
	if LOG_FILE_LOCATION != "" {
		log.SetOutput(&lumberjack.Logger{
			Filename:   LOG_FILE_LOCATION,
			MaxSize:    100, // megabytes
			MaxBackups: 2,
			MaxAge:     12,   //days
			Compress:   true, // disabled by default
		})
	}

	// Start Server
	go func() {
		log.Println("Starting simple go docker server...")
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	// Graceful Shutdown
	waitForShutdownSignal(srv)
}

func handler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "Unknoown"
	}
	log.Printf("Received %s\n", name)
	w.Write([]byte(fmt.Sprintf("Received, %s\n", name)))
}

func waitForShutdownSignal(srv *http.Server) {
	interruptSignal := make(chan os.Signal, 1)
	signal.Notify(interruptSignal, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// block for signal.
	<-interruptSignal

	// Use a context with a timeout to shutdown server 
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	srv.Shutdown(ctx)

	log.Println("Exiting simple go docker example... ")
	os.Exit(0)
}
