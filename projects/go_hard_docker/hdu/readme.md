# hdu - Hard docker UI

проект-эксперимент по реализации управления докером с помощью хардкорного html-интерфейса

## Линки

- https://habr.com/ru/articles/449038/
- https://echo.labstack.com/docs
- https://pkg.go.dev/github.com/docker/docker/client#section-readme

## Packages

```bash
go get github.com/docker/docker/client
go get github.com/labstack/echo/v4
```

## Docker API Version issue

Если возникнет ошибка версии api докера и клиента, можно в переменной окружения указать реальную версию

```bash
DOCKER_API_VERSION=1.43 go run ./cmd/app/
```
