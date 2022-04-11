package main

import (
	"github.com/lukasl-dev/octave/config"
	"log"
	"os"
	"os/signal"
)

func main() {
	a := newApp(config.Config{
		Token: os.Getenv("TOKEN"),
		Lavalink: config.Lavalink{
			Host:       os.Getenv("LAVALINK_HOST"),
			Passphrase: os.Getenv("LAVALINK_PASSPHRASE"),
		},
	})

	if err := a.run(); err != nil {
		log.Fatalf("Failed to run app: %s\n", err)
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, os.Kill)
	<-signals
}
