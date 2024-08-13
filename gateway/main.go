package main

import (
	r "github.com/Mubinabd/car-wash/api"
	"github.com/Mubinabd/car-wash/load"
	"github.com/Mubinabd/car-wash/api/handlers"
)

func main() {

	engine := r.NewRouter(handlers.NewHandlers())
	engine.Run(config.Load().HTTPPort)
}
