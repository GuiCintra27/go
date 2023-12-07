###Objective of the project

The project aims to create a isolated enviroment training my docker concepts, here the docker-compose create Postgres enviroment, create a database into it and todos table too.

Integrate golang with a SQL database (PostgresSQL) and create a CRUD api.

To run postgres and create database and tables

```
  docker-compose up -d
```

If you need to instal any package

```
  go mod tidy
```

To run the server

```
  go run main.go
```

After that, you should see something like

```
Server running on port 9000
```

To see the requests that you can make to api look the request_example.http file