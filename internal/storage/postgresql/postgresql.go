package postgresql

import (
	"database/sql"
	"fmt"
)

type Storage struct {
	db *sql.DB //connection string
}

func New(storagePath string) (*Storage, error) {
	const op = "storage.postgresql.New"

	//slog.With(slog.String("op", op))

	db, err := sql.Open("pgx", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("%s: failed to ping db: %w", op, err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) addURL(urlToSave string, alias string) (int64, error) {
	const op = "storage.postgresql.addURL"

	stmt, err := s.db.Prepare(
		`INSERT INTO Order (Id, IdOfClient, GoodsinOrder) 
		VALUES (NewGuid, IdOfClient, GoodsinOrder)`)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(alias, "created")
	if err != nil {
		return 0, fmt.Errorf("%s: exec statement failed: %w", op, err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("%s: failed to get last insert id: %w", op, err)
	}

	return id, nil
}

func (s *Storage) deleteURL(urlToSave int64) error {
	const op = "storage.postgresql.deleteURL"

	stmt, err := s.db.Prepare(
		`DELETE FROM Order WHERE id=Id`)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(urlToSave)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Storage) getAllURL() ([]int64, error) {
	const op = "storage.postgresql.getAllURL"

	stmt, err := s.db.Prepare(`
		SELECT Order.Id
		FROM Order 
		INNER JOIN dbo.GoodInOrder ON Order.IdOfClient = GoodInOrder.IdOfClient 
		INNER JOIN Good ON GoodInOrder.Id = Good.Id
		`)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	row, err := stmt.Query()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	var orders []int64
	for row.Next() {
		var order int64
		err := row.Scan(&order)
		if err != nil {
			return nil, fmt.Errorf("%s: scann failed: %w", op, err)
		}
		orders = append(orders, order)
	}

	return orders, nil
}

func (s *Storage) getByIdURL(id int64) (int64, error) {
	const op = "storage.postgresql.getByIdURL"

	stmt, err := s.db.Prepare(`
	SELECT * FROM dbo.Order WHERE id=Id'
	`)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	var order int64
	err = stmt.QueryRow(id).Scan(&order)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("%s: order not found", op)
		}
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}

func (s *Storage) updateURL(oldUrl string, urlToSave string, alias string) error {
	const op = "storage.postgresql.updateURL"

	stmt, err := s.db.Prepare(`
		UPDATE dbo.Order SET IdOfClient=order.IdOfClient 
		WHERE Id=order.Id
		`)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(urlToSave, alias)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
