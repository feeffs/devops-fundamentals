# 6. DevOps Fundamentals

## Структура репозитория

| Папка | Описание                                                                |
| ----- | ----------------------------------------------------------------------- |
| t1-2  | Создание Docker контейнера с Nginx. Передача аргументов через ARG       |
| t3    | Запуск нескольких контейнеров через Docker Compose (Nginx + PostgreSQL) |
| t4    | Go-приложение, подключающееся к PostgreSQL                              |
| t5    | Монтирование локальной папки в контейнер (bind mount)                   |
| t6    | Многоконтейнерное приложение: Nginx + Go + PostgreSQL                   |
| t7    | Сборка образа через docker-compose                                      |
| t8    | Создание и использование пользовательской сети Docker                   |
| t9    | Использование переменных окружения через .env                           |
| t10   | Работа с Docker Hub                                                     |

## Требования

- Docker Desktop
- Docker Compose

## Запуск

### t1-2

```bash
cd t1-2
docker build -t hello-docker --build-arg GREETING="Hello, World!" .
docker run -d -p 8080:80 --name my-nginx hello-docker
```

### t3, t5, t6, t7, t8

```bash
cd t3  # или t5, t6, t7, t8
docker compose up -d
```

### t4

```bash
cd t4
docker compose up --build
```

### t9

```bash
cd t9
# Скопируй .env.example в .env и заполни переменные
cp .env.example .env
docker compose up --build -d
```

## Переменные окружения

Скопируй `.env.example` в `.env` и заполни значения:

```bash
cp .env.example .env
```
