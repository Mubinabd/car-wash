package handlers

import (
	"github.com/Mubinabd/car-wash/internal/pkg/load"
	client "github.com/Mubinabd/car-wash/internal/service"
)

type Handlers struct {
	Clients *client.Clients
}

func NewHandlers() (*Handlers, error) {
	clients, err := client.NewClients(&load.Config{})
	if err != nil {
		return nil, err
	}	

	return &Handlers{
		Clients: clients,
	}, nil
}
