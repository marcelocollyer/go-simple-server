package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

type response struct {
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp,omitempty"`
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := response{Message: "Hello World!"}
	logRequest(r, http.StatusOK)
	jsonResponse(w, resp, http.StatusOK)
}

func currentTimeHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	resp := response{Message: "Hello " + name, Timestamp: time.Now().Unix()}
	logRequest(r, http.StatusOK)
	jsonResponse(w, resp, http.StatusOK)
}

func healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	logRequest(r, http.StatusOK)
	w.WriteHeader(http.StatusOK)
}

func jsonResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func logRequest(r *http.Request, statusCode int) {
	log.Printf(`{"method": "%s", "url": "%s", "status": %d}`, r.Method, r.URL.String(), statusCode)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello_world", helloWorldHandler).Methods("GET")
	r.HandleFunc("/current_time", currentTimeHandler).Methods("GET")
	r.HandleFunc("/healthcheck", healthcheckHandler).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
