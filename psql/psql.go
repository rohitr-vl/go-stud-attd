package psql

import "database/sql"

type PsqlRepository struct {
	db *sql.DB
}

func NewPostgreSQLClassicRepository(db *sql.DB) *PsqlRepository {
	return &PsqlRepository{
		db: db,
	}
}
