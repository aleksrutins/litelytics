package dbutil

import (
	"context"
	"log"

	"github.com/aleksrutins/litelytics/ent"
	_ "github.com/lib/pq"
	"github.com/profclems/go-dotenv"
)

var Client *ent.Client

func Connect() {
	var err error
	Client, err = ent.Open("postgres", dotenv.GetString("DATABASE_URL"))

	if err != nil {
		log.Fatalf("failed to open postgres connection: %v", err)
	}

	if err := Client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed to create schema resources: %v", err)
	}
}
