package storage

import "errors"

var (
	ErrUrlNotFound = errors.New("url not found")
	ErrUrlExist    = errors.New("url exist")
)

type StorageInter interface {
	AddURL(urlToSave string, alias string) (int64, error)

	DeleteURL(urlToSave int64) error

	GetAllURL() ([]int64, error)

	GetByIdURL(id int64) (int64, error)

	UpdateURL(oldUrl string, urlToSave string, alias string) error
}
