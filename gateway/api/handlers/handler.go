package handlers

import (
	"log"

	"github.com/Mubinabd/car-wash/client"
	config "github.com/Mubinabd/car-wash/load"
)

type Handlers struct {
	Clients client.Clients
}

func NewHandlers() *Handlers {
	cfg := config.Load()
	cl, err := client.NewClients(&cfg)
	if err != nil {
		log.Fatal(err)
	}
	return &Handlers{
		Clients: *cl,
	}
}
