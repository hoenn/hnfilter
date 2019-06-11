package store

import "context"

//Store represents what a store must be capable of
type Store interface {
	AddComment(context.Context, *Comment) error
	GetComment(context.Context, int) (*Comment, error)
	//FIXME this definition is probably not correct
	GetSearchResults(context.Context, string) error
}
