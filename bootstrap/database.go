package bootstrap

import (
	"context"
	"database/sql"
	"log"
)

type PostgreSQLClassicRepository struct {
	db *sql.DB
}

func NewPsqlDatabase(env *Environment) psql.Client {
	return &PostgreSQLClassicRepository{
		db: db,
	}
}

func ClosePsqlDBConnection(client psql.Client) {
	if client == nil {
		return
	}

	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connection to PostgreSql closed.")
}
