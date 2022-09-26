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
)

// Creates and listens an http server.
func startServer() {
	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", "localhost", 8888),
		Handler: router,
	}

	go func() {
		// Listen server with the specified settings.
		log.Printf("Listenig server %s\n", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error listening on %s: %v\n", server.Addr, err)
		}
	}()

	// Gracefully stop.
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Printf("Shutting down server...\n")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalln("Server forced to shutdown:", err)
	}

	log.Printf("Server closed\n")
}
