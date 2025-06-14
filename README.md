# RZD Sales Backend

REST API для поиска железнодорожных билетов, использующее API РЖД.

## Требования

- Go 1.21 или выше

## Установка

1. Перейдите в директорию backend:
```bash
cd backend
```

2. Установите зависимости:
```bash
go mod download
```

## Запуск

### Запуск rzd-api

```bash
cd rzd-api
docker run -p 8000:8000 --rm --name rzd-api -v ${PWD}:/app-it pavelsr/rzd-api
```

### Запуск нашего сервиса
```bash
cd backend
go run main.go
```

API будет доступно по адресу http://localhost:8080

## API Endpoints

### Проверка здоровья сервиса
```
GET /api/v1/health
```
Ответ:
```json
{
    "status": "ok"
}
```

### Поиск станций
```
GET /api/v1/stations?query={query}
```
Параметры:
- `query` - часть названия станции (минимум 2 символа)

Ответ:
```json
[
    {
        "code": "2004000",
        "name": "САНКТ-ПЕТЕРБУРГ"
    }
]
```

### Поиск поездов
```
GET /api/v1/trains?fromCode={fromCode}&toCode={toCode}&date={date}
```
Параметры:
- `fromCode` - код станции отправления
- `toCode` - код станции прибытия
- `date` - дата в формате YYYY-MM-DD

Ответ:
```json
{
    "trains": [
        {
            "number": "012А",
            "fromStation": "САНКТ-ПЕТЕРБУРГ",
            "toStation": "МОСКВА",
            "departureTime": "2024-01-01T12:00:00Z",
            "arrivalTime": "2024-01-02T08:00:00Z",
            "duration": "20:00",
            "price": 1500.00
        }
    ]
}
```

## Структура проекта

```
.
├── internal/
│   ├── config/      # Конфигурация приложения
│   ├── handlers/    # HTTP обработчики
│   ├── models/      # Модели данных
│   ├── rzd/         # Клиент API РЖД
│   └── server/      # HTTP сервер
├── main.go         # Точка входа
└── README.md
```

## Конфигурация

Основные настройки приложения через переменные окружения:
- `PORT` - порт для запуска сервера (по умолчанию 8080)
- `RZD_API_URL` - URL API РЖД
- `RZD_TIMEOUT` - таймаут для запросов к API РЖД в секундах (по умолчанию 10)

## Обработка ошибок

Все ошибки возвращаются в формате:
```json
{
    "error": "Описание ошибки"
}
```

## Лицензия

MIT 