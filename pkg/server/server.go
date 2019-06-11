package server

import (
	"context"
	"errors"

	"github.com/hoenn/hnfilter/pkg/store"
)

//This is the datastore abstraction layer

//Server wraps a datastore
type Server struct {
	d store.Store
}

//GetCommentAuthor is an example function
func (s *Server) GetCommentAuthor(id int) (string, error) {
	c, err := s.d.GetComment(context.Background(), id)
	if err != nil {
		return "", errors.New("could not get author for comment")
	}

	return c.Author, nil
}
