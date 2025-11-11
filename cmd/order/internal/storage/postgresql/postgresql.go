package postgresql

import (
	"fmt"
	"database/sql"
	"log/slog"

	"github.com/jackc/pgx/v5"
)

type Storage struct{
	db *sql.DB //connection string
}

func New(storagePath string) (*Storage, error) {
	const op = "storage.postgresql.New"

	slog.With(slog.String("op", op))

	db, err := sql.Open("postgresql", storagePath)
	if err != nil{
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	stmt, err := db.Prepare(
		CREATE TABLE IF NOT EXIST order(
			id INTEGER PRIMARY KEY
		);
	CREATE INDEX IF NOT EXISTS index ON order(id);
	)
	if err != nil{
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec()
	if err != nil{
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) addURL(urlToSave string, alias string) (int64, error) {

}

func (s *Storage) deleteURL(urlToSave string, alias string) (int64, error) {

}

func (s *Storage) getAllURL(urlToSave string, alias string) (int64, error) {

}

func (s *Storage) getByIdURL(urlToSave string, alias string) (int64, error) {

}

func (s *Storage) updateURL(urlToSave string, alias string) (int64, error) {

}
