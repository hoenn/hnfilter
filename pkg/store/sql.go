package store

import (
	//driver:
	_ "github.com/lib/pq"
	sq "gopkg.in/Masterminds/squirrel.v1"
)

func insertPost(c *Comment) sq.InsertBuilder {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	return psql.Insert(
		"comments",
	).Columns(
		"author",
		"id",
		"parent",
		"body",
		"time",
	).Values(
		c.By,
		c.ID,
		c.Parent,
		c.Body,
		c.Time,
	)
}
