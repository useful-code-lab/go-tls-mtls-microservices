package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	// Создаем клиента, который игнорирует проверку сертификатов (аналог curl -k)
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true, // ВОТ ОН: аналог флага -k
			},
		},
	}

	// Делаем запрос к вашему серверу
	resp, err := client.Get("https://localhost:8443/api/data")
	if err != nil {
		log.Fatalf("❌ Ошибка запроса: %v", err)
	}
	defer resp.Body.Close()

	// Читаем ответ
	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("✅ Ответ сервера: %s\n", string(body))
}
