package schemamigration

import "openingjobs/pkg/config"

var (
	UseCase *SchemaMigrationUseCase
)

func InitializeSchemaMigrationDBConext() error {
	db := config.GetDB()

	storage := NewSchemaMigrationStorage(db)
	UseCase = NewSchemaMigrationUseCase(storage)

	return nil
}
