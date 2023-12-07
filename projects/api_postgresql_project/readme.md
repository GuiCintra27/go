This is a server that only aims to register todos.

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