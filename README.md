Запрос без проверки кто подписал сертификат сервера
```bash
 curl -k https://localhost:8443/api/data
{"status": "secure", "message": "Данные переданы через TLS"}
```

```bash
go run cmd/server/server.go
🚀 Сервер запущен на https://localhost:8443
```

```bash
go run cmd/client/client.go
✅ Ответ сервера: {"status": "secure", "message": "Данные переданы через TLS"}
```

```bash
 go run cmd/mtls/server/server.go
🔒 Сервер в режиме mTLS запущен на https://localhost:8443
```