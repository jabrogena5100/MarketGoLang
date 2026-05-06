package main

import (
	"context"
	"log"
	"log/slog"
	"os"

	"github.com/jabrogena5100/MarketGoLang.git/internal/env"
	"github.com/jackc/pgx/v5"
)

func main() {
	ctx := context.Background()

	cfg := config{
		addr: ":8080",
		db:   dbConfig{
			dsn: env.GetString("GOOSE_DBSTRING", "host=localhost user=postgres password=postgres dbname=ecom sslmode=disable"),
		},

	}

	// Logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// Database (postgres through Docker)
	conn, err := pgx.Connect(ctx, cfg.db.dsn)
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)

	logger.Info("connected to Database", "dsn", cfg.db.dsn)

	api := application{
		config: cfg,
		db: conn,
	}

	//mount our api
if err := api.run(api.mount()); err != nil { 
	slog.Error("server failed to start", "error", err)
	log.Printf("server has failed to start, err: %s", err)
	os.Exit(1)
}
}
