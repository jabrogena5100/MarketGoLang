package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	repo "github.com/jabrogena5100/MarketGoLang.git/internal/adapters/postgresql/sqlc"
	"github.com/jabrogena5100/MarketGoLang.git/internal/orders"
	"github.com/jabrogena5100/MarketGoLang.git/internal/products"
	"github.com/jackc/pgx/v5"
)

//run -> graceful shutdown
//mount
func (app *application) mount() http.Handler { 
r := chi.NewRouter()



// Middle ware
r.Use(middleware.RequestID)
r.Use(middleware.RealIP)
r.Use(middleware.Logger)
r.Use(middleware.Recoverer)

r.Use(middleware.Timeout(60 * time.Second))

r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("all good"))
})

productService := products.NewService(repo.New(app.db)); 
productHandler := products.NewHandler(productService)
r.Get("/products", productHandler.ListProducts)

ordersHandler := orders.NewHandler(nil)
r.Post("/orders", ordersHandler.PlaceOrder)

return r
}

type application struct { 
config config 
// logger 
// db driver : without this application in main.go doesn't work
db *pgx.Conn
}

// run 
func (app *application) run(h http.Handler) error { 
	srv := &http.Server{Addr: app.config.addr, Handler: h, WriteTimeout: time.Second * 30, ReadTimeout: time.Second * 10, IdleTimeout: time.Minute }

	log.Printf("server has started at addr %s", app.config.addr)

	return srv.ListenAndServe()
}

type config struct { 
	addr string 
	db dbConfig
}

type dbConfig struct { 
	dsn string
}