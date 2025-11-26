package storage

import "errors"

var (
	ErrUrlNotFound = errors.New("url not found")
	ErrUrlExist    = errors.New("url exist")
)

type OrderService interface {
	AddURL(urlToSave string, alias string) (int64, error)

	DeleteURL(urlToSave string) error

	GetAllURL() ([]string, error)

	GetByIdURL(id string) (string, error)

	UpdateURL(oldUrl string, urlToSave string) error

	IsOrderCreatedURL(id string) (bool, error)
}
