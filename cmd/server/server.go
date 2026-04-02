package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/data", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"status": "secure", "message": "Данные переданы через TLS"}`)
	})

	fmt.Println("🚀 Сервер запущен на https://localhost:8443")

	certFile := "certs/server.crt"
	keyFile := "certs/server.key"

	err := http.ListenAndServeTLS(":8443", certFile, keyFile, nil)
	if err != nil {
		log.Fatalf("❌ Ошибка запуска сервера: %v", err)
	}
}
