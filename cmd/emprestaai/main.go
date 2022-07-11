package main

import (
	"github.com/brunogbarros/emprestaai.git/internal/service"
	"os"
	"os/signal"
)

func main() {

	go service.Start(service.NewApp(), ":8080")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

}
