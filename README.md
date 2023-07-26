# go-tax-project

- CRUD example with SOLID, interfaces, and good practices using GoLang.
- RabbitMQ to receive messages on "Order" queue and save on SQLite3, using channels.   

### DOCKER steps 
Before starting prepare your image, execute this docker steps:
```
docker build -t [YOUR_USER]/go-tax-project .
docker push [YOUR_USER]/go-tax-project
```

### Create table orders
```
sqlite3 db.sqlite3
CREATE TABLE orders (id VARCHAR(255) NOT NULL, price float NOT NULL, tax float NOT NULL, final_price NOT NULL, PRIMARY KEY (id));
```

### Run docker compose to start RabbitMq server
` docker-compose up -d`

### Creating and publishing message on the queue

Open web browser http://localhost:15672/ and create a queue called 'order', so publish the following json 

```
{"id": "3", "price": 10.0, "tax": 0.2}
{"id": "2", "price": 20.0, "tax": 3.2}
{"id": "3", "price": 45.0, "tax": 1.2}

```

### Creating kubernetes cluster (using kind)

Create cluster and apply manifests inside k8s path.
```
kind create cluster 
kubectl appply -f k8s
``` 

### Running
To run the `main.go` program: `go cmd/order/main.go`   
If you need debug code, might it run nodemon comand to watch changes  
```
nodemon --exec go run cmd/api/main.go --signal SIGTERM
```

### * Using docker 
If you prefer uses docker, run: `docker run -p 8888:8888 swirfneblin/go-tax-project`   

