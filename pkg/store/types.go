package store

import (
	"fmt"
	"time"
)

// Comment contains the data to store a comment
type Comment struct {
	Author string
	ID     int
	Kids   []int
	Parent int
	Body   string
	Time   time.Time
}

//DBConn is the required info for a postgres connection string
type DBConn struct {
	Username string
	Password string
	Host     string
	Port     string
	Name     string
}

//Format converts the string to a postgres conneciton string
func (d *DBConn) Format() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		d.Host,
		d.Port,
		d.Username,
		d.Password,
		d.Name)
}
