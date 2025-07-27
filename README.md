# Auth-Service

Сервис аутентификации на Go с PostgreSQL и Docker.

---

## Запуск

1. Склонировать репозиторий  
2. Запустить сервис командой:

``` bash

docker-compose -f docker-compose.yml up -d

```

Все сервисы (БД, миграции, auth-service) поднимутся автоматически.

---

## API

- **POST** `/api/auth/token?guid={user_guid}`  
  Получить пару токенов (access и refresh) для пользователя с GUID.

- **POST** `/api/auth/refresh`  
  Обновить пару токенов. В теле JSON с полями `access` и `refresh`.

- **GET** `/api/auth/user`  
  Получить GUID текущего пользователя (требуется access token в заголовке Authorization).

- **POST** `/api/auth/logout`  
  Выход (деавторизация) пользователя.
  
---

## Swagger 

Для удобного тестирования и изучения API доступен Swagger UI по адресу:
``` bash

http://localhost:8080/swagger/index.html

```
