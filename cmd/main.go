package main

import (
	"log"
	"log/slog"
	"os"
)

func main() {

	cfg := config{
		addr: ":8080",
		db:   dbConfig{},
	}

	api := application{
		config: cfg,
	}

// Logging
logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

slog.SetDefault(logger)


	//mount our api
if err := api.run(api.mount()); err != nil { 
	slog.Error("server failed to start", "error", err)
	log.Printf("server has failed to start, err: %s", err)
	os.Exit(1)
}
}
