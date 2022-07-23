package main

import (
	"github.com/brunogbarros/emprestaai.git/internal/config"
	"github.com/brunogbarros/emprestaai.git/internal/service"
	"os"
	"os/signal"
)

func main() {
	go service.Start(service.NewApp(), string(rune(config.Port)))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

}
