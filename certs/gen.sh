#!/bin/bash
cd certs

# 1. Генерируем CA (если уже есть, можно не трогать, но лучше обновить всё сразу)
openssl genrsa -out ca.key 2048
openssl req -x509 -new -nodes -key ca.key -sha256 -days 3650 -out ca.crt \
  -subj "/CN=GuardGo-CA"

# 2. Генерируем ключ сервера
openssl genrsa -out server.key 2048

# 3. Генерируем запрос (CSR) и САМОЕ ВАЖНОЕ — добавляем SAN для localhost
openssl req -new -key server.key -out server.csr -subj "/CN=localhost" \
  -addext "subjectAltName = DNS:localhost, IP:127.0.0.1"

# 4. Подписываем сертификат сервера с учетом расширений SAN
openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial \
  -out server.crt -days 365 -sha256 \
  -copy_extensions copyall

# 5. Генерируем клиентский сертификат (для клиента SAN обычно не критичен, но сделаем для порядка)
openssl genrsa -out client.key 2048
openssl req -new -key client.key -out client.csr -subj "/CN=GuardGo-Client"
openssl x509 -req -in client.csr -CA ca.crt -CAkey ca.key -CAcreateserial \
  -out client.crt -days 365 -sha256

rm *.csr *.srl
cd ..
