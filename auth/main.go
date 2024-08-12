package main

import (
	"github.com/Mubinabd/car-wash/config"
	"github.com/Mubinabd/car-wash/pkg/app"
)

func main() {
	cfg := config.Load()
	app.Run(&cfg)
}
