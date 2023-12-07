To run the application:
```
  docker-compose up
```

Enter into products database and create new products table

```
  docker-compose mysql bash and create a new products database and table

  mysql -uroot -p products and create table products (id varchar(255), name varchar(255), price float)
```

Create new product topic in kafka

```
  docker-compose kafka bash

  kafka-topics --bootstrap-server=localhost:9092 --topic=product --create
```

Run the server

```
 docker-compose goapp bash (maybe you need to go to src folder - [cd ../src/api_and_messaging_project])

  go run cmd/app/main.go
```
make a http request to http://localhost:8000/products (example in request_example.http file), or insert a new item in product topic in kafka