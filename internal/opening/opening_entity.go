package opening

import (
	"github.com/google/uuid"
)

type Opening struct {
	ID       uuid.UUID `db:"id"`
	Role     string    `db:"role"`
	Company  string    `db:"company"`
	Location string    `db:"location"`
	Remote   bool      `db:"remote"`
	Link     string    `db:"link"`
	Salary   int64     `db:"salary"`
}
