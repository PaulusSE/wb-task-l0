# WB Tech: level # 0 (Golang)
Этот проект является тестовым заданием для WB Tech. Он представляет собой демонстрационный сервис с простейшим интерфейсом, отображающий данные о заказе. Данные о заказе представлены в формате JSON и прилагаются к заданию.

### Цель проекта
Целью проекта является показать навыки разработки на языке go, а также работу с базой данных PostgreSQL и брокером сообщений nats-streaming.

### Функциональность проекта
##### Проект состоит из следующих компонентов:

- Сервис, который подключается к каналу в nats-streaming и получает данные о заказе в формате JSON. Сервис записывает данные в базу данных PostgreSQL и кэширует их в памяти. В случае падения сервиса, он восстанавливает кэш из базы данных.
- HTTP-сервер, который выдает данные о заказе по id из кэша. Сервер также отображает данные в виде HTML-страницы с помощью шаблонизатора html/template.
- Скрипт, который публикует данные о заказе в канал nats-streaming для проверки работы подписки.

### Зависимости проекта
##### Проект использует следующие зависимости:

- nats-streaming - брокер сообщений, который обеспечивает надежную доставку данных между сервисами.
- PostgreSQL - реляционная база данных, которая хранит данные о заказе.
- sqlc - генератор кода для работы с базой данных с помощью SQL-запросов.
- migrate - инструмент для управления миграциями базы данных.
- [html/template] - стандартный пакет для генерации HTML-страниц из шаблонов.

### Установка проекта
##### Для установки проекта необходимо выполнить следующие шаги:

Клонировать репозиторий проекта:
```bash
git clone https://github.com/yourname/wb-task-l0.git
cd wb-task-l0
```

##### Установить зависимости:
```bash
go mod download

```
##### Запустить nats-streaming и PostgreSQL с помощью docker:
```bash
docker run -p 4222:4222 -ti nats-streaming:latest -cid my_cluster 
docker run -d --name postgresql_container -e POSTGRES_PASSWORD=secret -p 5432:5432 postgres
```

##### Создать базу данных и пользователя для проекта:
```bash
psql -h localhost -U postgres -c "CREATE DATABASE wb_l0;"
psql -h localhost -U postgres -c "CREATE USER wb_l0 WITH PASSWORD 'secret';"
psql -h localhost -U postgres -c "GRANT ALL PRIVILEGES ON DATABASE wb_l0 TO wb_l0;"
```

##### Применить миграции базы данных с помощью migrate:
```bash
migrate -path db/migrations -database "postgres://wb_l0:secret@localhost:5432/wb_l0?sslmode=disable" -verbose up
```

##### Скопировать пример файла .env и заполнить его своими значениями:
```bash
cp .env.example .env
```

##### Запустить сервис и HTTP-сервер:
```bash
go run main.go

```
Открыть [http://localhost:8080] в браузере и ввести id заказа для получения данных.
Использование проекта
Для использования проекта необходимо запустить скрипт, который публикует данные о заказе в канал nats-streaming:

```bash
go run publisher.go

```

[========]

Скрипт считывает данные из файла model.json и отправляет их в канал orders. Сервис подписан на этот канал и получает данные, записывает их в базу данных и кэш. HTTP-сервер выдает данные из кэша по запросу.

[========]

`Лицензия проекта
Проект распространяется под лицензией MIT. См. файл LICENSE для подробностей.`