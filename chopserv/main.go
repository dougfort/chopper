package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/dougfort/chopper/chopserv/config"
	"github.com/dougfort/chopper/chopserv/httpserv"
	"github.com/dougfort/chopper/chopserv/types"
)

func main() {
	os.Exit(run())
}

// TODO: Use Dave Cheney's errors wrapper
// TODO: use key/value logging

func run() int {
	var cfg types.Config
	var err error

	log.Printf("info: chopserv starts")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	if cfg, err = config.Load(); err != nil {
		log.Printf("error: config.Load() failed: %s", err)
		return -1
	}

	ctx, cancel := context.WithCancel(context.Background())

	go httpserv.Serve(ctx, cfg)

	// block until sigterm
	s := <-sigChan
	log.Printf("info: received signal: %s", s)
	cancel()

	log.Printf("info: chopserv terminates normally")
	return 0
}
