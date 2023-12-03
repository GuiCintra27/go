package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/GuiCintra27/go/api_and_messaging/internal/infra/akafka"
	"github.com/GuiCintra27/go/api_and_messaging/internal/infra/repository"
	"github.com/GuiCintra27/go/api_and_messaging/internal/infra/web"
	"github.com/GuiCintra27/go/api_and_messaging/internal/usecase"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/go-chi/chi/v5"
)

// To run the application:
// run docker-compose up
// run docker-compose mysql bash and create a new products database and table
// run docker-compose kafka bash and create a new product topic
// run docker-compose goapp bash (maybe you need to go to src folder - [cd ../src/api_and_messaging])
// run go run cmd/app/main.go
// make a http request to http://localhost:8000/products (example in request_example.http file), or insert a new item in product topic in kafka

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(host.docker.internal:3306)/products")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	repository := repository.NewProductRepositoryMySQL(db)
	createProductUseCase := usecase.NewCreateProductUseCase(repository)
	listProductsUseCase := usecase.NewListProductsUseCase(repository)

	productHandlers := web.NewProductHandlers(createProductUseCase, listProductsUseCase)

	r := chi.NewRouter()

	r.Post("/products", productHandlers.CreateProductHandler)
	r.Get("/products", productHandlers.ListProductsHandler)

	go http.ListenAndServe(":8000", r)

	msgChan := make(chan *kafka.Message)
	go akafka.Consume([]string{"product"}, "host.docker.internal:9094", msgChan)

	for msg := range msgChan {
		dto := usecase.CreateProductInputDTO{}
		err := json.Unmarshal(msg.Value, &dto)
		
		if err != nil {
			log.Println(err)
		}

		_, err = createProductUseCase.Execute(dto)

		if err != nil {
			log.Println(err)
		}
	}
}