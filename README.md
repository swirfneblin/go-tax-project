# go-tax-project

- CRUD example with SOLID, interfaces, and good practices using GoLang.
- RabbitMQ to receive messages on "Order" queue and save on SQLite3, using channels.   


1. Create table orders
```
sqlite3 db.sqlite3
CREATE TABLE orders (id VARCHAR(255) NOT NULL, price float NOT NULL, tax float NOT NULL, final_price NOT NULL, PRIMARY KEY (id));
```

2. Run docker compose to start RabbitMq server
` docker-compose up -d`

3. Open web browser http://localhost:15672/ and create a queue called 'order', so publish the following json 

```
{"id": "3", "price": 10.0, "tax": 0.2}
{"id": "2", "price": 20.0, "tax": 3.2}
{"id": "3", "price": 45.0, "tax": 1.2}

```

4. To run the `main.go` program: `go cmd/order/main.go`

