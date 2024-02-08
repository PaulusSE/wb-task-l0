create table "orders"(
    order_uid varchar PRIMARY KEY,
    order_json jsonb Unique NOT NULL
)