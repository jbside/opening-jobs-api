package opening

import (
	"database/sql"
	"openingjobs/pkg/response"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type OpeningStorage struct {
	db *sqlx.DB
	qb *squirrel.StatementBuilderType
}

func newOpeningStorage(db *sqlx.DB, qb *squirrel.StatementBuilderType) *OpeningStorage {
	return &OpeningStorage{
		db: db,
		qb: qb,
	}
}

func (storage *OpeningStorage) createOpening(opening *Opening) error {
	var err error

	query, args, err := storage.qb.Insert("opening").
		Columns("role", "company", "location", "remote", "link", "salary").
		Values(opening.Role, opening.Company, opening.Location, opening.Remote, opening.Link, opening.Salary).
		ToSql()

	if err != nil {
		return err
	}

	err = storage.db.Get(opening, query, args...)

	if err != nil {
		return err
	}

	return nil
}

func (storage *OpeningStorage) updateOpening(id *uuid.UUID, opening *Opening) error {
	var err error

	queryBuilder := storage.qb.Update("opening")
	queryBuilder.Set("role", opening.Role)
	queryBuilder.Set("company", opening.Company)
	queryBuilder.Set("location", opening.Location)
	queryBuilder.Set("remote", opening.Remote)
	queryBuilder.Set("link", opening.Link)
	queryBuilder.Set("salary", opening.Salary)
	queryBuilder.Where(squirrel.Eq{"id": id})

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		if err == sql.ErrNoRows {
			return response.NotFoundError()
		}
		return err
	}

	err = storage.db.Get(opening, query, args...)

	if err != nil {
		return err
	}

	return nil
}

func (storage *OpeningStorage) deleteOpening(id *uuid.UUID) error {
	var err error

	query, args, err := storage.qb.Delete("opening").
		Where(squirrel.Eq{"id": id}).
		ToSql()

	if err != nil {
		return err
	}

	_, err = storage.db.Exec(query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (storage *OpeningStorage) listOpenings() ([]*OpeningResponseDTO, error) {
	var err error

	query, args, err := storage.qb.Select("role", "company", "location", "remote", "link", "salary").
		From("opening").
		ToSql()

	if err != nil {
		return nil, err
	}

	var openings []*OpeningResponseDTO
	err = storage.db.Select(&openings, query, args...)
	if err != nil {
		return nil, err
	}

	return openings, nil
}

func (storage *OpeningStorage) showOpening(id uuid.UUID) (*OpeningResponseDTO, error) {
	var err error

	query, args, err := storage.qb.
		Select("role", "company", "location", "remote", "link", "salary").
		From("opening").
		Where(squirrel.Eq{"id": id}).
		ToSql()

	if err != nil {
		return nil, err
	}

	var opening OpeningResponseDTO
	err = storage.db.Get(&opening, query, args...)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, response.NotFoundError()
		}
		return nil, err
	}

	return &opening, nil
}
