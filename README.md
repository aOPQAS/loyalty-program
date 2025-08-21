# Microservice: Loyalty Program API

## Возможности
- Управление программами (`program`)
- Управление услугами (`services`)
- Управление связями программ и услуг (`program_services`)
- Получение данных о подпродуктах из Telebon API
- Защита API через Bearer Token (middleware)
- Эндпоинт для проверки состояния сервиса

## Стек технологий
- **Go** — язык программирования
- **net/http + Fiber** — обработка HTTP-запросов
- **PostgreSQL** — база данных
- **SQL** — работа с БД
- **Zap** — логирование
- **godotenv + env** — конфигурация
- **Telebon API** — интеграция

## Установка и запуск

```bash
git clone https://github.com/aOPQAS/loyalty-program.git
cd microservice
go mod tidy
go run ./cmd/main.go
```
Запускать необходимо из **корневой директории проекта**

## Запуск сервиса через Docker

### Сборка Docker образа
```bash
docker build -t microservice .
```

### Запуск контейнера
```bash
docker run -d --name microservice -p 8080:8080 --env-file .env microservice
```

### Проверка работы сервиса
```bash
curl -H "Authorization: Bearer <TOKEN>" http://localhost:8080/api/program
```

# Проверка работы маршрутов

## Products

### `GET /subproducts`
**Описание:** Получает список всех подпродуктов. \
**Пример curl:**
```bash
curl -X GET http://localhost:8080/api/subproducts
  -H "Authorization: Bearer <TOKEN>"
```

## Programs

### `GET /program`
**Описание:** Получает список всех подпродуктов. \
**Пример curl:**
```bash
curl -H "Authorization: Bearer b445e61bcf21fec8e0e3203237dbf26e" http://localhost:8080/api/program
```

### `GET /program/:id`
**Описание:** Получает программу по её ID. \
**Пример curl:**
```bash
curl -H "Authorization: Bearer <TOKEN>" http://localhost:8080/api/program/1
```

### `POST /program/create`
**Описание:** Создает новую программу. \
**Пример curl:**
```bash
curl -X POST http://localhost:8080/api/program/create
  -H "Authorization: Bearer <TOKEN>"
  -H "Content-Type: application/json"
  -d '{
    "type": "абонемент",
    "name": "Пакет \"Мужской уход\"",
    "fixed_price": 1050,
    "total_services_cost": 1500,
    "discount_percent": 30,
    "services": [
      {
        "service_id": "6899d7adbbdb1bad6c15b987",
        "name": "Стрижка Мужская",
        "tarif": 1000,
        "duration": 60
      },
      {
        "service_id": "6899d7aebbdb1bad6c15b989",
        "name": "Стрижка Усов",
        "tarif": 500,
        "duration": 20
      }
    ],
    "valid_until": "2025-12-31"
```

### `PUT /program/update`
**Описание**: Обновляет данные существующей программы. \
**Пример curl**:
```bash
curl -X PUT http://localhost:8080/api/program/update
  -H "Authorization: Bearer <TOKEN>"
  -H "Content-Type: application/json"
  -d '{
    "id": "89b9b702-c003-465a-b9e6-40ed4723f1cf",
    "type": "абонемент",
    "name": "Пакет \"Мужской уход - обновлённый\"",
    "fixed_price": 1200,
    "total_services_cost": 1600,
    "discount_percent": 25,
    "services": [
      {
        "service_id": "6899d7adbbdb1bad6c15b987",
        "name": "Стрижка Мужская",
        "tarif": 1100,
        "duration": 60
      },
      {
        "service_id": "6899d7aebbdb1bad6c15b989",
        "name": "Стрижка Усов",
        "tarif": 500,
        "duration": 20
      }
    ],
    "valid_until": "2026-01-31",
    "active": true
}'
```

### `DELETE /program/:id`
**Описание:** Удаляет связь программы с сервисом. \
**Пример curl:**
```bash
curl -X DELETE http://localhost:8080/api/program/<PROGRAM_ID> \
  -H "Authorization: Bearer <TOKEN>"
```

# services
### `GET /services`
**Описание:** Получает список всех сервисов. \
**Пример curl:**
```bash
curl -X GET http://localhost:8080/api/services
-H "Authorization: Bearer <TOKEN>"
```
### `POST /services/create`
**Описание:** Создает новый сервис. \
**Пример curl:**
```bash
curl -X POST http://localhost:8080/api/services/create
-H "Content-Type: application/json"
-d '{
  "name": "Мой сервис",
  "tarif": 500,
  "duration": 60
}'
```

### `PUT /services/update`
**Описание:** Обновляет существующий сервис. \
**Пример curl:**
```bash
curl -X PUT http://localhost:8080/api/services/update
-H "Content-Type: application/json"
-d '{
  "service_id": "be2ea95b-632e-437f-8373-a4e3f3efda9b",
  "name": "Мой сервис PRO",
  "tarif": 700,
  "duration": 90
}'
```

### `DELETE /services/delete`
**Описание:** Удаляет сервис. \
**Пример curl:**
```bash
 curl -X DELETE http://localhost:8080/api/services/delete
-H "Content-Type: application/json"
-d '{
  "service_id": "fbebe488-5d78-4fd2-9e7d-1853a8ea060c"
}'
```

## Program services

### `GET /program_services`
**Описание:** Получает все связи программ с сервисами. \
**Пример curl:**
```bash
curl -X GET http://localhost:8080/program_services
-H "Authorization: Bearer <TOKEN>"
```

### `POST /program_services/create`
**Описание:** Создает связь между программой и сервисом. \
**Пример curl:**
```bash
curl -X POST http://localhost:8080/api/program_services/create
-H "Content-Type: application/json"
-d '{
  "program_id": "84a49dee-32e3-424a-a7a4-59c93bdadca0",
  "service_id": "7650b2c1-c812-43eb-86c1-ec7bb81d0076"
}'
```

### `DELETE /program_services/delete`
**Описание:** Удаляет связь программы с сервисом. \
**Пример curl:**
```bash
curl -X DELETE http://localhost:8080/api/program_services/delete
-H "Content-Type: application/json"
-d '{
  "program_id": "84a49dee-32e3-424a-a7a4-59c93bdadca0",
  "service_id": "7650b2c1-c812-43eb-86c1-ec7bb81d0076"
}'
```
