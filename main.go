package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	server := &http.Server{
		Addr:         ":8080",
		Handler:      http.HandlerFunc(CorsExamplesHandler),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("starting on port 8080")

	log.Fatal(server.ListenAndServe())
}

const allowedOrigin string = "https://kolesa.kz"

func CorsExamplesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("new request received:")
	_ = r.Write(os.Stdout)

	switch r.URL.Path {
	case "/without-any-headers":
		_, _ = w.Write([]byte("Hello, this is method without preflight handling"))
	case "/only-simple":
		w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
		_, _ = w.Write([]byte("Hello, you can send only simply requests here"))
	case "/":
		w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)

		// preflight-request
		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, DELETE")
			w.WriteHeader(http.StatusNoContent)
		}

		_, _ = w.Write([]byte(fmt.Sprintf("Hi %s!", allowedOrigin)))
	case "/all-origins":
		w.Header().Set("Access-Control-Allow-Origin", "*")

		// preflight-request
		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, PATCH, DELETE")
			w.WriteHeader(http.StatusNoContent)
		}

		_, _ = w.Write([]byte("Hi there!"))
	default:
		http.NotFound(w, r)
	}
}
