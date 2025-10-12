package main

import (
	"log"

	"github.com/adedejiosvaldo/gopher_social/internal/env"
	"github.com/adedejiosvaldo/gopher_social/internal/store"
)

func main() {
	cfg := config{address: env.GetString("ADDR", ":8080")}

	store := store.NewPostgresStorage(nil)

	app := &application{config: cfg, store: store}

	mux := app.mount()

	log.Fatal(app.run(mux))

}
