CREATE TABLE myschema.orders (
    order_uid varchar PRIMARY KEY,
    order_json jsonb UNIQUE NOT NULL
);