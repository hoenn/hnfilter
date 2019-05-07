package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/hoenn/go-hn/pkg/hnapi"
	_ "github.com/lib/pq"
)

func main() {
	c := hnapi.NewHNClient()
	u, err := c.User("whoishiring")
	if err != nil {
		log.Fatal(err)
	}
	var posts []string
	for _, s := range u.Submitted {
		i := strconv.Itoa(s)
		posts = append(posts, i)
	}
	//For each submitted story get the post

	var stories []*hnapi.Story
	x := 0
	for _, p := range posts {
		if x > 0 {
			break
		}
		x++
		item, err := c.Item(p)
		if err != nil {
			log.Fatal(err)
		}
		switch i := item.(type) {
		case *hnapi.Story:
			stories = append(stories, i)
		default:
			//skip
		}
	}
	fmt.Println(len(stories))
	//Filter the stories by title
	var filteredStories []*hnapi.Story
	for _, s := range stories {
		if strings.Contains(s.Title, "Who is hiring?") {
			filteredStories = append(filteredStories, s)
			fmt.Println(s)
		}
	}
	//Store it in the db
	dbConn := &DBConn{
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Name:     os.Getenv("DB_NAME"),
	}
	db, err := sql.Open("postgres", dbConn.Format())
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
}

type DBConn struct {
	Username string
	Password string
	Host     string
	Port     string
	Name     string
}

func (d *DBConn) Format() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		d.Host,
		d.Port,
		d.Username,
		d.Password,
		d.Name)
}
