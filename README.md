# gRPC Orders

gRPC-микросервис на Go для управления заказами, использующий Protocol Buffers для строго типизированного API с unary и streaming RPC.
Сервис построен на слоистой архитектуре (handlers, services, repositories) и включает interceptor’ы для логирования и обработки запросов.

---

## Возможности

- Создания заказа 
- Получения заказа по ID  
- Получения списка заказа  
- Хранилище In-Memory

---

## Быстрый старт

Запуск с in-memory хранилищем:

```bash
git clone https://github.com/ultrabor/grpc-orders.git
cd go-user-api
go run ./...
```

Микросервис будет доступен по адресу:  
`http://localhost:50051`

---


