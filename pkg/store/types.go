package store

import "time"

// Comment contains the data to store a comment
type Comment struct {
	Author string
	ID     int
	Kids   []int
	Parent int
	Body   string
	Time   time.Time
}
