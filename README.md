Генерация сертификатов
```bash
bash certs/gen.sh
```

Запрос без проверки кто подписал сертификат сервера
```bash
 curl -k https://localhost:8443/api/data
{"status": "secure", "message": "Данные переданы через TLS"}
```
# TLS
```bash
go run cmd/server/server.go
🚀 Сервер запущен на https://localhost:8443
```

```bash
go run cmd/client/client.go
✅ Ответ сервера: {"status": "secure", "message": "Данные переданы через TLS"}
```

# MTLS
```bash
 go run cmd/mtls/server/server.go
🔒 Сервер в режиме mTLS запущен на https://localhost:8443
```


```bash
aleksey@aleksey:~/guard-go$ go run cmd/mtls/client/client.go
✅ Успех mTLS! Ответ сервера: {"status": "mTLS Success", "client": "GuardGo-Client"}
```