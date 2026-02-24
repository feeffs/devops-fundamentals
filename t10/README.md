# Task 10 — Docker Hub

Образ запушен на Docker Hub: https://hub.docker.com/r/feffs/hello-docker

## Запуск образа

```bash
docker run -d -p 8080:80 feffs/hello-docker:latest
```

http://localhost:8080 — "Hello, Docker!"
