package main

import (
	"log"
	"net/http"
	"os"

	"github.com/izaakdale/ittp"
	"github.com/rs/cors"
)

func main() {
	mux := ittp.NewServeMux()

	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("200 OK"))
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	})
	mux.AddMiddleware(c.Handler)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
