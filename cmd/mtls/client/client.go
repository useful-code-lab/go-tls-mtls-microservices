package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	// 1. Загружаем сертификат нашего CA (чтобы проверять сервер)
	caCert, err := os.ReadFile("certs/ca.crt")
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// 2. Загружаем личный СЕРТИФИКАТ и КЛЮЧ клиента (наш "паспорт")
	// Это те самые файлы, которые потребует сервер в режиме mTLS
	clientCert, err := tls.LoadX509KeyPair("certs/client.crt", "certs/client.key")
	if err != nil {
		log.Fatal(err)
	}

	// 3. Настраиваем TLS-конфигурацию клиента
	tlsConfig := &tls.Config{
		RootCAs:      caCertPool,               // Кому доверяем мы (серверу)
		Certificates: []tls.Certificate{clientCert}, // Что предъявляем мы (наш паспорт)
	}

	// 4. Создаем HTTP-клиент с этой конфигурацией
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}

	// 5. Делаем запрос
	resp, err := client.Get("https://localhost:8443/api/data")
	if err != nil {
		log.Fatalf("❌ Сервер отверг нас: %v", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("✅ Успех mTLS! Ответ сервера: %s\n", string(body))
}
