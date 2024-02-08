package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"

	"github.com/PaulusSE/wbL0/api"
	"github.com/PaulusSE/wbL0/db"
	"github.com/PaulusSE/wbL0/db/sqlc"
	"github.com/PaulusSE/wbL0/util"
	_ "github.com/lib/pq"
	"github.com/nats-io/stan.go"
)

func main() {
	cache := util.New()
	db_p, err := sql.Open("postgres", "postgresql://root:secret@localhost:5432/wb_l0?sslmode=disable")
	queries := sqlc.New(db_p)

	// load from db to cache
	jsonRows, err := queries.GetOrders(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(jsonRows); i++ {
		cache.Set(jsonRows[i].OrderUid, jsonRows[i].OrderJson)
	}

	// Подключение к NATS Streaming Server
	sc, err := stan.Connect(
		"my_cluster",                          // Имя кластера NATS Streaming
		"client-123",                          // Уникальный идентификатор клиента
		stan.NatsURL("nats://localhost:4222"), // URL сервера NATS Streaming
	)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	// if err != nil {
	// 	log.Fatal(err)
	// }
	defer sc.Close()

	sub, err := sc.Subscribe("foo", func(m *stan.Msg) {
		var msgJson db.Order
		err := json.Unmarshal(m.Data, &msgJson)
		if err != nil {
		}
		queries.AddOrder(context.Background(), sqlc.AddOrderParams{
			OrderUid:  msgJson.Order_uid,
			OrderJson: m.Data,
		})
		cache.Set(msgJson.Order_uid, msgJson)
	}, stan.StartWithLastReceived())

	defer sub.Unsubscribe()
	srv := api.NewServer(sc, cache)
	srv.Start()

}
