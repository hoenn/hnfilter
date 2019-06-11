package store

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
)

// DataStore wraps a database
type DataStore struct {
	db *sql.DB
}

// NewDataStore returns a datastore
func NewDataStore(db *sql.DB) *DataStore {
	return &DataStore{
		db: db,
	}
}

var errCouldNotCommitTx = "could not commit transaction"
var errFailedStartingTx = "failed starting transaction"

// AddComment adds a comment to the datastore
func (d *DataStore) AddComment(ctx context.Context, c *Comment) error {
	tx, err := d.db.BeginTx(ctx, nil)
	if err != nil {
		return errors.Wrap(err, errFailedStartingTx)
	}
	_, err = insertComment(c).RunWith(tx).ExecContext(ctx)
	if err != nil {
		return errors.Wrap(err, "could not insert comment")
	}
	err = tx.Commit()
	if err != nil {
		return errors.Wrap(err, errCouldNotCommitTx)
	}
	return nil
}

// GetComment returns a comment from the datastore by id
func (d *DataStore) GetComment(ctx context.Context, id int) (*Comment, error) {
	tx, err := d.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, errors.Wrap(err, errFailedStartingTx)
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
		return nil, errors.Wrap(err, errCouldNotCommitTx)
	}

	return c, nil
}

//GetSearchResults TODO
func (d *DataStore) GetSearchResults(ctx context.Context, search string) error {
	//TODO Implement this function
	return nil
}
