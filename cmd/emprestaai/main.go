package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/brunogbarros/emprestaai.git/internal/service"
)

func main() {

	go service.Start(service.NewApp(), ":8080")
	go service.Start(service.NewHealthApp(), ":8081")

	sinal := make(chan os.Signal, 1)
	signal.Notify(sinal, os.Interrupt, syscall.SIGTERM)
	switch <-sinal {
	case os.Interrupt:
		log.Fatal("Interrupted signal")
	case syscall.SIGTERM:
		log.Fatal("Interrupted by user signal")
	}

}
