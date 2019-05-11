package store

import (
	"fmt"

	sq "gopkg.in/Masterminds/squirrel.v1"
)

var psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

var commentsCols = []string{
	"author",
	"id",
	"parent",
	"body",
	"time",
}

func insertComment(c *Comment) sq.InsertBuilder {
	return psql.Insert(
		"comments",
	).Columns(
		commentsCols...,
	).Values(
		c.Author,
		c.ID,
		c.Parent,
		c.Body,
		c.Time,
	)
}

func getComment(id int) sq.SelectBuilder {
	return psql.Select(
		commentsCols...,
	).From(
		"comments",
	).Where(
		sq.Eq{
			"id": fmt.Sprint(id),
		},
	)
}
