-- name: GetOrders :many
SELECT * from "orders";

-- name: AddOrder :one
INSERT INTO "orders" (order_uid,order_json) VALUES($1,$2) RETURNING *;