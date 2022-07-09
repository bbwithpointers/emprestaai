package main

import (
	"github.com/brunogbarros/emprestaai.git/internal/service"
)

func main() {

	go service.Start(service.NewApp(), ":8080")

}
