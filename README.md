# gRPC мессенджер. Домашка по кафедральному курсу в МФТИ

Приложение представляет из себя консольный мессенджер с возможностью получать сообщения не в real time-е

Использована технология `gRPC` для передачи сообщений и `Redis` для хранения сообщений

Для запуска сервера, нужно запустить `go run ./cmd/grpc_server/main.go` из кроня проекта

Для запуска клиента, нужно запустить `go run ./cmd/grpc_client/main.go` из корня проекта
