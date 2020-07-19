package main

import (
	"context"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/ztsu/snowboardsdb/snowboardsdb/graphql"
	"github.com/ztsu/snowboardsdb/snowboardsdb/postgres"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	addr := "0.0.0.0:80"

	bgCtx := context.Background()

	pool, err := pgxpool.Connect(bgCtx, os.Getenv("POSTGRES_DSN"))
	if err != nil {
		log.Fatalf("couldn't connect to postgres: %s", err)
	}

	mux := http.NewServeMux()

	mux.Handle("/graphql", graphql.Handler(NewStores(pool)))
	mux.Handle("/pg", playground.Handler("GraphQL playground", "/graphql"))

	s := http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	done := make(chan bool)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-quit

		log.Print("Stopping server...")

		ctx, cancel := context.WithTimeout(bgCtx, 5*time.Second)
		defer cancel()

		if err := s.Shutdown(ctx); err != nil {
			log.Fatalf("can't stop server: %s", err)
		}

		close(done)
	}()

	log.Printf("Starting webserver at %s...", addr)

	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("can't start server: %s", err)
		}
	}()

	log.Print("Visit http://localhost:80/pg for playground")

	<-done

	log.Print("Buy")
}

func NewStores(pool *pgxpool.Pool) *graphql.Stores {
	return &graphql.Stores{
		Brands:     postgres.NewBrandsStore(pool),
		Persons:    postgres.NewPersonsStore(pool),
		Catalogues: postgres.NewCataloguesStore(pool),
		Snowboards: postgres.NewSnowboardsStore(pool),
		Images:     postgres.NewImageStore(pool),
	}
}
