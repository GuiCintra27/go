#GO WITH DOCKER >>>

# FROM golang:1.20 as builder

# WORKDIR /app

# COPY /deploy .

# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server

# FROM scratch

# COPY --from=builder /app/server /server

# ENTRYPOINT [ "/server" ]

FROM golang:latest

WORKDIR /go/app

RUN apt-get update && apt-get install -y librdkafka-dev

CMD ["tail", "-f", "/dev/null"]