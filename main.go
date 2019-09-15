package main

import (
	_ "jrpc-orm/conf"
	"jrpc-orm/routers"
	"log"
	"net/http"
	"os"
)

func main() {
	router := routers.NewRouter()

	httpServer := http.Server{
		Addr:    os.Getenv("ADDR"),
		Handler: router,
	}

	log.Printf("HTTP2 jrpc-orm Server started on %s", os.Getenv("ADDR"))

	err := httpServer.ListenAndServeTLS(
		os.Getenv("TLS_PEM"), os.Getenv("TLS_KEY"))

	if err != nil {
		log.Fatal(err)
	}
}
