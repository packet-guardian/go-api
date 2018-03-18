package main

import (
	"fmt"

	"github.com/packet-guardian/go-api"
)

const (
	pgURL      = "http://localhost:8080"
	pgUsername = "admin"
	pgPassword = "admin"
)

func main() {
	fmt.Println("Packet Guardian status example")

	pg, _ := api.Connect(pgURL)
	pg.Login(pgUsername, pgPassword)
	status, err := pg.GetStatus()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%#v\n", status)
}
