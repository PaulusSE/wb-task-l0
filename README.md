# wb_l0 task

Демонстрационный сервис заказов (Golang)
Этот проект представляет собой демонстрационный сервис, написанный на языке программирования Go, который использует технологии nats-streaming и PostgreSQL для обработки данных о заказах.


## Команды для запуска
Запуск nats-streaming
docker run -p 4222:4222 -ti nats-streaming:latest -cid my_cluster 
Запуск PostgreSQL
docker run -d --name postgresql_container -e POSTGRES_PASSWORD=secret -p 5432:5432 postgres
Используемые инструменты
Go: Язык программирования для разработки сервиса.
PostgreSQL: Реляционная база данных для хранения данных о заказах.
nats-streaming: Система сообщений для подписки на каналы и обработки заказов.
SQLC: Инструмент для работы с базой данных через SQL.
Пример запуска миграций
Для управления миграциями используется инструмент golang-migrate.

migrate -path db/migrations -database "databaseUrl" -verbose up 
migrate -path db/migrations -database "databaseUrl" -verbose down