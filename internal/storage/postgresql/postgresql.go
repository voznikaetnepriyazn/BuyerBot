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
	const op = "storage.postgresql.addURL"

	stmt, err := s.db.Prepare("INSERT INTO dbo.Order (Id, IdOfClient, GoodsinOrder) VALUES ('{Guid.NewGuid}, {order.IdOfClient}, {order.GoodsinOrder}')")
	if err != nil{
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	res, err := stmt.Exec(urlToSave, alias)
	if err != nil{
		if postgresql, ok := err.(postgresql.Error); ok && postgresqlErr.ExtendedCode == postgresql.ErrConstraint {
			return 0, fmt.Errorf("%s :%w", op, storage.ErrURLExists)
		}

		return 0, fmt.Errorf("%s: %w", op, err)
	}

	id, err := res.LastInsertId()
	if err != nil{
		return 0, fmt.Errorf("%s: failed to get last insert id: %w", op, err)
	}
}

func (s *Storage) deleteURL(urlToSave string, alias string) (int64, error) {
	const op = "storage.postgresql.addURL"

	stmt, err := s.db.Prepare("DELETE FROM dbo.Order WHERE id='{Id}'")
	if err != nil{
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	res, err := stmt.Exec(urlToSave, alias)
	if err != nil{
		if postgresql, ok := err.(postgresql.Error); ok && postgresqlErr.ExtendedCode == postgresql.ErrConstraint {
			return 0, fmt.Errorf("%s :%w", op, storage.ErrURLExists)
		}

		return 0, fmt.Errorf("%s: %w", op, err)
	}

	id, err := res.LastInsertId()
	if err != nil{
		return 0, fmt.Errorf("%s: failed to get last insert id: %w", op, err)
	}
}

func (s *Storage) getAllURL(urlToSave string, alias string) (int64, error) {
	const op = "storage.postgresql.addURL"

	stmt, err := s.db.Prepare("SELECT Order.Id, Good.Id FROM dbo.Order INNER JOIN dbo.GoodInOrder ON Order.IdOfClient = GoodInOrder.IdOfClient INNER JOIN Good ON GoodInOrder.Id = Good.Id")
	if err != nil{
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	res, err := stmt.Exec(urlToSave, alias)
	if err != nil{
		if postgresql, ok := err.(postgresql.Error); ok && postgresqlErr.ExtendedCode == postgresql.ErrConstraint {
			return 0, fmt.Errorf("%s :%w", op, storage.ErrURLExists)
		}

		return 0, fmt.Errorf("%s: %w", op, err)
	}

	id, err := res.LastInsertId()
	if err != nil{
		return 0, fmt.Errorf("%s: failed to get last insert id: %w", op, err)
	}
}

func (s *Storage) getByIdURL(urlToSave string, alias string) (int64, error) {
	const op = "storage.postgresql.addURL"

	stmt, err := s.db.Prepare("SELECT * FROM dbo.Order WHERE id='{Id}'")
	if err != nil{
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	res, err := stmt.Exec(urlToSave, alias)
	if err != nil{
		if postgresql, ok := err.(postgresql.Error); ok && postgresqlErr.ExtendedCode == postgresql.ErrConstraint {
			return 0, fmt.Errorf("%s :%w", op, storage.ErrURLExists)
		}

		return 0, fmt.Errorf("%s: %w", op, err)
	}

	id, err := res.LastInsertId()
	if err != nil{
		return 0, fmt.Errorf("%s: failed to get last insert id: %w", op, err)
	}
}

func (s *Storage) updateURL(urlToSave string, alias string) (int64, error) {
	const op = "storage.postgresql.addURL"

	stmt, err := s.db.Prepare("UPDATE dbo.Order SET IdOfClient='{order.IdOfClient}' WHERE Id='{order.Id}'")
	if err != nil{
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	res, err := stmt.Exec(urlToSave, alias)
	if err != nil{
		if postgresql, ok := err.(postgresql.Error); ok && postgresqlErr.ExtendedCode == postgresql.ErrConstraint {
			return 0, fmt.Errorf("%s :%w", op, storage.ErrURLExists)
		}

		return 0, fmt.Errorf("%s: %w", op, err)
	}

	id, err := res.LastInsertId()
	if err != nil{
		return 0, fmt.Errorf("%s: failed to get last insert id: %w", op, err)
	}
}
