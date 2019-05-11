package store

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
)

// DataStore xyz
type DataStore struct {
	db *sql.DB
}

// NewDataStore xyz
func NewDataStore(db *sql.DB) *DataStore {
	return &DataStore{
		db: db,
	}
}

// AddComment cyz
func (d *DataStore) AddComment(ctx context.Context, c *Comment) error {
	tx, err := d.db.BeginTx(ctx, nil)
	if err != nil {
		return errors.Wrap(err, "failed starting transaction")
	}
	_, err = insertPost(c).RunWith(tx).ExecContext(ctx)
	if err != nil {
		return errors.Wrap(err, "could not insert comment")
	}
	err = tx.Commit()
	if err != nil {
		return errors.Wrap(err, "failed to commit transaction")
	}
	return nil
}
