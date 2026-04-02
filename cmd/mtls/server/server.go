package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// 1. Загружаем наш CA (чтобы сервер знал, каким сертификатам КЛИЕНТОВ доверять)
	caCert, err := os.ReadFile("certs/ca.crt")
	if err != nil {
		log.Fatalf("Ошибка чтения ca.crt: %v", err)
	}

	// 2. Создаем пул и добавляем туда наш CA
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// 3. Настраиваем TLS-конфигурацию сервера
	tlsConfig := &tls.Config{
		ClientCAs:  caCertPool,               // Кому доверяем (нашему CA)
		ClientAuth: tls.RequireAndVerifyClientCert, // ТРЕБУЕМ сертификат от клиента!
	}

	// 4. Создаем HTTP-сервер с этой конфигурацией
	server := &http.Server{
		Addr:      ":8443",
		TLSConfig: tlsConfig,
	}

	// Наш защищенный эндпоинт
	http.HandleFunc("/api/data", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		// Теперь мы можем даже узнать имя клиента из его сертификата!
		clientName := r.TLS.PeerCertificates[0].Subject.CommonName
		fmt.Fprintf(w, `{"status": "mTLS Success", "client": "%s"}`, clientName)
	})

	fmt.Println("🔒 Сервер в режиме mTLS запущен на https://localhost:8443")

	// Запускаем (сертификат сервера по-прежнему нужен для шифрования)
	err = server.ListenAndServeTLS("certs/server.crt", "certs/server.key")
	if err != nil {
		log.Fatalf("❌ Ошибка: %v", err)
	}
}
