package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	rules "github.com/slf-aobrien/hackday2023rules"
)

type Message struct {
	Message string      `json:"message"`
	Code    string      `json:"status"`
	Extra   interface{} `json:"extra"`
}

func main() {
	fmt.Println("Starting the Rules Service")
	os.Setenv("TZ", "UTC")
	r := chi.NewRouter()

	// Basic CORS
	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

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

	r.Options("/v1/validatewithcode", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("success"))
	})

	r.Options("/v1/validatewithgorule", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("success"))
	})

	r.Options("/v1/validatewithgrule", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("success"))
	})

	r.Post("/v1/validatewithcode", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Validation User")

		var testUser rules.Users
		err := json.NewDecoder(r.Body).Decode(&testUser)
		check(err)
		fmt.Println("found a user named: " + testUser.FirstName + " " + testUser.LastName)

		validation := basicengine.validate(testUser)
		w.Write(marshall(validation))
	})

	r.Post("/v1/validatewithgorule", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Validation User")

		var testUser rules.Users
		err := json.NewDecoder(r.Body).Decode(&testUser)
		check(err)
		fmt.Println("found a user named: " + testUser.FirstName + " " + testUser.LastName)
		validation := validateWithGoRule(testUser)

		w.Write(marshall(validation))
	})

	r.Post("/v1/validatewithgrule", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Validation User")

		var testUser rules.Users
		err := json.NewDecoder(r.Body).Decode(&testUser)
		check(err)
		fmt.Println("found a user named: " + testUser.FirstName + " " + testUser.LastName)
		validation := validateWithGrule(testUser)

		w.Write(marshall(validation))
	})

	fmt.Println("Service running at: http://localhost:8889/")

	// a better way to do this: https://github.com/slus-customer-experience/slus-auth/blob/main/cmd/auth/main.go
	http.ListenAndServe(":8889", r)
}

func marshall(data interface{}) []byte {
	json, err := json.Marshal(data)
	check(err)
	return json
}

func check(err error) bool {

	if err != nil {
		log.Println(err)
		return true
	}
	return false
}

func validateWithGoRule(user rules.Users) Message {
	//insert real rule set here
	message := Message{}
	message.Message = "Success"
	message.Code = "OK"
	message.Extra = "No Validation Performed"

	return message
}

func validateWithGrule(user rules.Users) Message {
	//insert real rule set here
	message := Message{}
	message.Message = "Success"
	message.Code = "OK"
	message.Extra = "No Validation Performed"

	return message
}
