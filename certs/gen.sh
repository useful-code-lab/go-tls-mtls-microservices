#!/bin/bash

# 1. Создаем Корневой Сертификат (CA) — это наш "главный документ"
openssl genrsa -out ca.key 2048
openssl req -x509 -new -nodes -key ca.key -sha256 -days 3650 -out ca.crt \
  -subj "/CN=GuardGo-CA/O=Development"

# 2. Создаем сертификат для СЕРВЕРА
openssl genrsa -out server.key 2048
# Запрос на подпись (CSR)
openssl req -new -key server.key -out server.csr -subj "/CN=localhost"
# Подписываем сертификат сервера нашим CA
openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial \
  -out server.crt -days 365 -sha256

# 3. Создаем сертификат для КЛИЕНТА (понадобится для mTLS позже)
openssl genrsa -out client.key 2048
openssl req -new -key client.key -out client.csr -subj "/CN=GuardGo-Client"
openssl x509 -req -in client.csr -CA ca.crt -CAkey ca.key -CAcreateserial \
  -out client.crt -days 365 -sha256

# Удаляем временные файлы запросов
rm *.csr *.srl
