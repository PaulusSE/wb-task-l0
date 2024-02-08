package db

import "time"

type Order struct {
	Order_uid          string    `json:"order_uid"`
	Track_number       string    `json:"track_number"`
	Entry              string    `json:"entry"`
	Del                Delivery  `json:"delivery"`
	Payment            Payments  `json:"payment"`
	Item               []Items   `json:"items"`
	Locale             string    `json:"locale"`
	Internal_signature string    `json:"internal_signature"`
	Customer_id        string    `json:"customer_id"`
	Delivery_service   string    `json:"delivery_service"`
	Shardkey           string    `json:"shardkey"`
	Sm_id              int       `json:"sm_id"`
	Data_created       time.Time `json:"data_created"`
	Cof_shard          string    `json:"cof_shard"`
}

type Delivery struct {
	Name    string
	Phone   string
	Zip     string
	City    string
	Address string
	Region  string
	Email   string
}

type Payments struct {
	Transaction   string
	Request_id    string
	Currency      string
	Provider      string
	Amount        int
	Payment_dt    int
	Bank          string
	Delivery_cost int
	Goods_total   int
	Custom_fee    int
}

type Items struct {
	Chrt_id      int
	Track_Number string
	Price        int
	Rid          string
	Name         string
	Sale         int
	Size         string
	Total_price  int
	Nm_id        int
	Brand        string
	Status       int
}

type Test struct {
	Name string
	Face int
}
