package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	fmt.Println("Starting the Rules Service")
	os.Setenv("TZ", "UTC")
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Returning welcome message")
		w.Write([]byte("Welcome"))
	})
	fmt.Println("Service running at: http://localhost:8889/")

	// a better way to do this: https://github.com/slus-customer-experience/slus-auth/blob/main/cmd/auth/main.go
	http.ListenAndServe(":8889", r)
}
