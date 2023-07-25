# go-tax-project

Presents a CRUD example with SOLID, interfaces,  and good practices using GoLang

```
sqlite3 db.sqlite3
CREATE TABLE orders (id VARCHAR(255) NOT NULL, price float NOT NULL, tax float NOT NULL, final_price NOT NULL, PRIMARY KEY (id));
```


```
go run main.go
```