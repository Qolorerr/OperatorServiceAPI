# OperatorServiceAPI

Простое REST API приложение на Go для работы операторов с обращениями пользователей. Использует PostgreSQL для хранения данных и Redis для кэширования.

## Компоненты приложения

- **HTTP Server**: Сервер на основе [Gorilla MUX](github.com/gorilla/mux), обрабатывающий запросы.
- **Роутер**: Маршруты для CRUD операций с операторами, обращениями, группами и тэгами.
- **Хэндлеры**: Логика обработки HTTP-запросов.
- **Сервисы**: Бизнес-логика приложения.
- **Репозитории**: Взаимодействие с PostgreSQL через [GORM](https://gorm.io/).
- **Кэш**: Использование Redis для кэширования часто запрашиваемых данных.
- **Миграции**: Автоматическое создание таблиц в PostgreSQL через GORM.

## Сущности и связи

### Оператор (`operators`)
- `id` (UUID) - уникальный идентификатор.
- `deleted` (boolean) - флаг удалён.
- `created_at` (timestamp) - дата создания.
- `updated_at` (timestamp) - дата обновления.

### Группа операторов (`operator_groups`)
- `id` (UUID) - уникальный идентификатор.
- `deleted` (boolean) - флаг удалён.
- `created_at` (timestamp) - дата создания.
- `updated_at` (timestamp) - дата обновления.

### Обращение (`appeals`)
- `id` (UUID) - уникальный идентификатор.
- `deleted` (boolean) - флаг удалён.
- `user_id` (UUID) - уникальный идентификатор пользователя.
- `weight` (integer) - вес обращения
- `created_at` (timestamp) - дата создания.
- `updated_at` (timestamp) - дата обновления.

### Тэг (`tags`)
- `id` (UUID) - уникальный идентификатор.
- `deleted` (boolean) - флаг удалён.
- `name` (string) - Название тэга
- `created_at` (timestamp) - дата создания.
- `updated_at` (timestamp) - дата обновления.

**Связи**:
- `operators`|`id` - `operator_id`|`operator_group_links`|`group_id` - `id`|`operator_groups`
- `operator_groups`|`id` - `group_id`|`group_tag_links`|`tag_id` - `id`|`tags`
- `tags`|`id` - `tag_id`|`appeal_tag_links`|`appeal_id` - `id`|`appeals`

## Запуск приложения

### Предварительные условия
- Установите [Go](https://go.dev/dl/) (версия 1.21+).
- Установите [Docker](https://www.docker.com/) (для запуска PostgreSQL и Redis).

### Запуск
1. Запустите PostgreSQL и Redis через Docker Compose:
   ```bash
   docker-compose up -d
   ```
2. Запустите сервер:
   ```bash
   go run main.go
   ```
   В аргументах запуска укажите следующие параметры (указаны дефолтные):
   ``` 
   --port 8080
   --redisAddr localhost:6379
   --redisUser testUser
   --redisPassword
   --redisDB 0
   --redisMaxRetries 5
   --redisDialTimeout 10
   --redisTimeout 5
   --pgAddr localhost
   --pgPort 5432
   --pgUser test_user
   --pgPassword test_password
   --pgDbName test_db
   ```

API будет доступно по адресу: `http://localhost:8080/`.

## Разворот PostgreSQL
Для установки PostgreSQL вручную:
1. Скачайте и установите [PostgreSQL](https://www.postgresql.org/download/).
2. Создайте базу данных и пользователя:
   ```sql
   CREATE USER test_user WITH PASSWORD 'test_password';
   CREATE DATABASE test_db OWNER test_user;
   \c test_db
   CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
   ```

## Разворот Redis
Для установки Redis вручную:
- Установите [Redis](https://redis.io/download/)

## Набор тестовых данных (`data.sql`)
В проекте присутствует файл с примерами данных в таблицах

## Используемые технологии
- [Gorilla MUX](github.com/gorilla/mux) - HTTP-фреймворк.
- [GORM](https://gorm.io/) - ORM для PostgreSQL.
- [Redis Go Client](https://github.com/redis/go-redis/v9) - клиент для Redis.
