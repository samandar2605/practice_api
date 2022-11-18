package storage

import (
	"github.com/jmoiron/sqlx"
	"github.com/samandar2605/practice_api/storage/postgres"
	"github.com/samandar2605/practice_api/storage/repo"
)

type StorageI interface {
	Students() repo.StudentStorageI
}

type storagePg struct {
	studentRepo repo.StudentStorageI
}

func NewStoragePg(db *sqlx.DB) StorageI {
	return &storagePg{
		studentRepo: postgres.NewStudent(db),
	}
}

func (s *storagePg) Students() repo.StudentStorageI {
	return s.studentRepo
}
