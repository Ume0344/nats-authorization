package main

import "fmt"
import "github.com/nats-io/nats.go"

func main() {
    fmt.Println("Connecting to nats")
	url := "nats://nats-dev.dpp.vwapps.run:4222"
	nc, err := nats.Connect(url, nats.UserCredentials("/home/habiba/projects/nats-authorization/default.creds"))

	if err != nil {
		fmt.Printf("Error connecting to Nats; %s\n", err.Error())
		return
	}

	fmt.Printf("Connected to Nats successfully; %s\n", nc)
}
