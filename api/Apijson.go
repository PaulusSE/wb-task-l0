package api

import (
	"encoding/json"
	"net/http"

	"github.com/PaulusSE/wb-task-l0/db"
	"github.com/gin-gonic/gin"
)

type uid struct {
	Order_uid string `json:"order_uid"`
}

func (serv *Server) getJson(ctx *gin.Context) {
	orderUID := ctx.Query("order_uid")
	if orderUID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "order_uid parameter is missing"})
		return
	}

	obj, b := serv.Cache.Get(orderUID)
	if !b {
		ctx.JSON(http.StatusInternalServerError, "Non existing record")
		return
	}
	ctx.JSON(http.StatusOK, obj)
}

func (serv *Server) postJson(ctx *gin.Context) {
	var poston db.Order
	if err := ctx.ShouldBindJSON(&poston); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if (poston.Del == db.Delivery{
		Name:    "",
		Phone:   "",
		Zip:     "",
		City:    "",
		Address: "",
		Region:  "",
		Email:   "",
	}) ||
		(poston.Item == nil) ||
		(poston.Payment == db.Payments{
			Transaction:   "",
			Request_id:    "",
			Currency:      "",
			Provider:      "",
			Amount:        0,
			Payment_dt:    0,
			Bank:          "",
			Delivery_cost: 0,
			Goods_total:   0,
			Custom_fee:    0,
		}) {
		ctx.JSON(http.StatusBadRequest, "Invalid input")
		return
	}
	jposton, _ := json.Marshal(poston)
	serv.Con.Publish("foo", jposton)
	var p db.Order
	json.Unmarshal(jposton, &p)
	ctx.JSON(http.StatusOK, p)
}

func (serv *Server) homePage(c *gin.Context) {
	html := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Go Service Frontend</title>
		<style>
			body {
				font-family: Arial, sans-serif;
			}
			.container {
				margin: 20px;
			}
			.form-group {
				margin-bottom: 10px;
			}
			button {
				padding: 8px 16px;
				background-color: #4CAF50;
				color: white;
				border: none;
				cursor: pointer;
			}
			button:hover {
				background-color: #45a049;
			}
		</style>
	</head>
	<body>
		<div class="container">
			<h1>Go Service Frontend</h1>
			<div class="form-group">
				<label for="orderUid">Order UID:</label>
				<input type="text" id="orderUid" name="orderUid">
			</div>
			<button onclick="getJson()">Get JSON</button>
			<hr>
			<div class="form-group">
				<label for="postData">Post Data:</label>
				<textarea id="postData" name="postData"></textarea>
			</div>
			<button onclick="postJson()">Post JSON</button>
			<div id="response"></div>
		</div>
	
		<script>
		function getJson() {
			var orderUid = document.getElementById("orderUid").value;
			fetch('/getJson?order_uid=' + orderUid, {
				method: 'GET',
				headers: {
					'Content-Type': 'application/json'
				}
			})
			.then(response => response.json())
			.then(data => {
				document.getElementById("response").innerText = JSON.stringify(data, null, 2);
			})
			.catch(error => {
				console.error('Error:', error);
			});
		}
	
			function postJson() {
				var postData = document.getElementById("postData").value;
				fetch('/postJson', {
					method: 'POST',
					headers: {
						'Content-Type': 'application/json'
					},
					body: postData
				})
				.then(response => response.json())
				.then(data => {
					document.getElementById("response").innerText = JSON.stringify(data, null, 2);
				})
				.catch(error => {
					console.error('Error:', error);
				});
			}
		</script>
	</body>
	</html>	
`
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(html))
}
