package repo

import (
	"context"

	"github.com/go-pg/pg/v10"
)

func ConnectDB() *pg.DB{
	db := pg.Connect(&pg.Options{
		Addr:     ":5432",
		User:     "user",
		Password: "pass",
		Database: "db_name",
	})
    defer db.Close()

	ctx := context.Background()

if err := db.Ping(ctx); err != nil {
    panic(err)
}
return db
}
