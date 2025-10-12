package main

import (
	"log"

	"github.com/adedejiosvaldo/gopher_social/internal/db"
	"github.com/adedejiosvaldo/gopher_social/internal/env"
	"github.com/adedejiosvaldo/gopher_social/internal/store"
)

func main() {

	cfg := config{
		address: env.GetString("ADDR", ":8080"),
		dbConfig: dbConfig{
			address:      env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/social?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m")}}

	db, err := db.New(cfg.dbConfig.address, cfg.dbConfig.maxOpenConns, cfg.dbConfig.maxIdleConns, cfg.dbConfig.maxIdleTime)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	log.Println("Database connection established")
	store := store.NewPostgresStorage(db)

	app := &application{config: cfg, store: store}

	mux := app.mount()

	log.Fatal(app.run(mux))

}
