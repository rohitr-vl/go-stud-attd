package repository

import (
	"clean-architecture/psql/query_collection"
	"context"
	"log"

	"github.com/jackc/pgx/pgtype"
)

type Repository interface {
	GetUser(ctx context.Context, id int64) (User, error)
}

func run() error {
	queries := query_collection.New(conn)

	// list all users
	users, err := queries.ListUsers(ctx)
	if err != nil {
		return err
	}
	log.Println(users)

	// create an author
	insertedUser, err := queries.CreateUser(ctx, query.CreateUserParams{
		Name:  "Brian Kernighan",
		Email: "bk@testing.com",
		Bio:   pgtype.Text{String: "Co-author of The C Programming Language and The Go Programming Language", Valid: true},
	})
	if err != nil {
		return err
	}
	log.Println(insertedUser)

	return nil
}
