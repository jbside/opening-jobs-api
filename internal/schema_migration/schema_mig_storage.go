package schemamigration

import "github.com/jmoiron/sqlx"

type SchemaMigrationStorage struct {
	db *sqlx.DB
}

func NewSchemaMigrationStorage(db *sqlx.DB) *SchemaMigrationStorage {
	return &SchemaMigrationStorage{
		db: db,
	}
}

func (storage *SchemaMigrationStorage) appliedMigrations() (map[string]bool, error) {
	rows, err := storage.db.Query(`SELECT filename FROM schema_migrations`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	applied := make(map[string]bool)
	for rows.Next() {
		var f string
		rows.Scan(&f)
		applied[f] = true
	}
	return applied, nil
}

func (storage *SchemaMigrationStorage) markMigrationApplied(filename string) error {
	_, err := storage.db.Exec(`INSERT INTO schema_migrations (filename) VALUES ($1)`, filename)
	return err
}

func (storage *SchemaMigrationStorage) unmarkMigration(filename string) error {
	_, err := storage.db.Exec(`DELETE FROM schema_migrations WHERE filename = $1`, filename)
	return err
}

func (storage *SchemaMigrationStorage) execSql(sql string) error {
	_, err := storage.db.Exec(sql)
	return err
}
