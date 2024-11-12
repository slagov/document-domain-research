package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

const port = 80
const domainVictim = "victim-test.ru"
const domainVictim3rdParty = "3rd-party." + domainVictim

type TokenResponse struct {
	Token string `json:"token"`
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s [%s] %s%s\n", r.Method, time.Now().Format(time.RFC3339), r.Host, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func SubdomainVictimGetStealerHTMLHandler(writer http.ResponseWriter, request *http.Request) {
	http.ServeFile(writer, request, "try-steal.html")
}

func SubdomainVictimGetStealerHTMLImpersonateHandler(writer http.ResponseWriter, request *http.Request) {
	http.ServeFile(writer, request, "try-steal-impersonate.html")
}

func VictimGetTokenHandler(writer http.ResponseWriter, request *http.Request) {
	token := "5up32-53c2e7-70k3n-52"
	data, _ := json.Marshal(TokenResponse{Token: token})

	writer.Header().Set("Content-Type", "application/json")
	writer.Write(data)
	writer.WriteHeader(http.StatusOK)
}

func VictimOptionsTokenHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", fmt.Sprintf("http://%s:%d", domainVictim, port))
	writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	writer.Header().Set("Access-Control-Max-Age", "0")
	writer.WriteHeader(http.StatusOK)
}

func main() {
	r := mux.NewRouter()
	r.Use(loggingMiddleware)

	// victim-test.ru setup
	r.HandleFunc("/", VictimGetTokenHandler).Host(domainVictim).Methods(http.MethodGet)
	r.HandleFunc("/", VictimOptionsTokenHandler).Host(domainVictim).Methods(http.MethodOptions)

	// 3rd-party.victim-test.ru setup
	r.HandleFunc("/try-steal", SubdomainVictimGetStealerHTMLHandler).Host(domainVictim3rdParty).Methods(http.MethodGet)
	r.HandleFunc("/try-steal-impersonate", SubdomainVictimGetStealerHTMLImpersonateHandler).Host(domainVictim3rdParty).Methods(http.MethodGet)

	fmt.Printf("Starting server, listening on port %d\n", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}
