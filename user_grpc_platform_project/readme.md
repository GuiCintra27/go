run postgres server with docker

```docker-compose up``` 

run go gRCP server

```go run cmd/app/main.go -port=${port number here}```

if server starts ok, you probably will see something like

```Starting server on port ${port}```

run go gRCP client 

```go run client/main.go -port=${port number here}```

if the request is successfull you should see something like

```
  Making request on port ${port}
  token: ${user token}
```

*** make sure to change the email (in client file) everytime make a request