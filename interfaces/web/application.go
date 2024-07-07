package web

import (
	"errors"
	"log"
	"net/http"
)

func RunApplication() {

	clients := GetClients()
	NewWSTradingListener(clients)

	http.HandleFunc("/", GetTransactionHandler)
	http.HandleFunc("/ws", WSHandler)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: http.DefaultServeMux,
	}

	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("ListenAndServe(): %s", err)
	}
}
