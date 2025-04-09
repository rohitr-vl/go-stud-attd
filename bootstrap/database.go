package bootstrap

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"clean-architecture/psql"
)

func NewPsqlDatabase(env *Environment) *psql.PsqlRepository {
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dbHost := env.DBHost
	dbPort := env.DBPort
	dbUser := env.DBUser
	dbPass := env.DBPass

	psqlURI := fmt.Sprintf("postgres://%s:%s@%s:%s", dbUser, dbPass, dbHost, dbPort)

	if dbUser == "" || dbPass == "" {
		psqlURI = fmt.Sprintf("mongodb://%s:%s", dbHost, dbPort)
	}
	db, err := sql.Open("pgx", psqlURI)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    psqlPtr := psql.NewPostgreSQLClassicRepository(db)

	return psqlPtr
}
