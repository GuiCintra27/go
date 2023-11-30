package main

import (
	"encoding/json"
	"fmt"
)

type Client struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone,omitempty"`
}

func (c Client) FullName() {
	fmt.Printf(" %s %s\n", c.FirstName, c.LastName)
}

type InternationalClient struct {
	Client
	Country   string `json:"country"`
}

func main() {
	client := Client{
		FirstName: "Client",
		LastName:  "Normal",
		Email:     "jdoe@me.com",
		Phone:     "555-555-5555",
	}

	InternationalClient := InternationalClient{
		Client:    Client{
			FirstName: "Client",
			LastName:  "International",
			Email:     "jdoe@me.com",
			Phone:     "555-555-5555",
		},
		Country:   "US",
	}

	fmt.Println("Client:")
	client.FullName()

	fmt.Println("\nInternational Client:")
	InternationalClient.FullName()

	clientJson, err := json.Marshal(client)

	if err != nil {
		panic(err)
	}

	fmt.Println("\nClient JSON:")

	fmt.Printf("%s\n", string(clientJson))

	jsonClient2 := `{
		"first_name": "Client2",
		"last_name": "Normal",
		"email": "jdoe@me.com",
		"phone": "555-555-5555"
	}`

	var client2 Client

	json.Unmarshal([]byte(jsonClient2), &client2)

	fmt.Println("\nClient 2:\n First Name:", client2.FirstName)
	
}
