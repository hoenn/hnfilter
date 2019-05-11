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

//ErrCouldNotCommitTx xyz
var ErrCouldNotCommitTx = "could not commit transaction"

//ErrFailedStartingTx xyz
var ErrFailedStartingTx = "failed starting transaction"

// AddComment cyz
func (d *DataStore) AddComment(ctx context.Context, c *Comment) error {
	tx, err := d.db.BeginTx(ctx, nil)
	if err != nil {
		return errors.Wrap(err, ErrFailedStartingTx)
	}
	_, err = insertComment(c).RunWith(tx).ExecContext(ctx)
	if err != nil {
		return errors.Wrap(err, "could not insert comment")
	}
	err = tx.Commit()
	if err != nil {
		return errors.Wrap(err, ErrCouldNotCommitTx)
	}
	return nil
}

// GetComment xyz
func (d *DataStore) GetComment(ctx context.Context, id int) (*Comment, error) {
	tx, err := d.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, errors.Wrap(err, ErrFailedStartingTx)
	}

	var c *Comment
	row := getComment(id).RunWith(tx).QueryRowContext(ctx)
	err = row.Scan(
		&c.Author,
		&c.ID,
		&c.Kids,
		&c.Parent,
		&c.Body,
		&c.Time,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, errors.Wrap(err, "could not get comment")
	}
	err = tx.Commit()
	if err != nil {
		return nil, errors.Wrap(err, ErrCouldNotCommitTx)
	}

	return c, nil
}
