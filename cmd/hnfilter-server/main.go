package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/hoenn/go-hn/pkg/hnapi"
	"github.com/hoenn/hnfilter/pkg/store"
	_ "github.com/lib/pq"
)

func main() {
	client := hnapi.NewHNClient()
	u, err := client.User("whoishiring")
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
		if x > 3 {
			break
		}
		s, err := GetStoryByID(p, client)
		if err != nil {
			//Not a story
			continue
		}
		stories = append(stories, s)
		x++
	}

	//Filter the stories by title
	var filteredStories []*hnapi.Story
	for _, s := range stories {
		if strings.Contains(s.Title, "Who is hiring?") {
			filteredStories = append(filteredStories, s)
		}
	}

	var comments []*store.Comment
	x = 0
	for _, s := range filteredStories {
		for _, c := range s.Kids {
			if x > 0 {
				break
			}
			cc, err := GetCommentByID(fmt.Sprint(c), client)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(cc)
			comments = append(comments, cc)
			x++
		}
	}

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

	ds := store.NewDataStore(db)

	for _, c := range comments {
		err := ds.AddComment(context.Background(), c)
		if err != nil {
			log.Fatal(err)
		}
	}

}

// GetStoryByID returns a story from an id
func GetStoryByID(id string, c *hnapi.HNClient) (*hnapi.Story, error) {
	item, err := c.Item(id)
	if err != nil {
		return nil, err
	}

	s, ok := item.(*hnapi.Story)
	if !ok {
		return nil, fmt.Errorf("could not get story from id:%s", id)
	}
	return s, nil
}

// GetCommentByID returns a comment in db form from an id
func GetCommentByID(id string, c *hnapi.HNClient) (*store.Comment, error) {
	item, err := c.Item(id)
	if err != nil {
		return nil, err
	}

	s, ok := item.(*hnapi.Comment)
	if !ok {
		return nil, fmt.Errorf("could not get comment from id:%s", id)
	}

	t := time.Unix(s.Time, 0)

	return &store.Comment{
		By:     s.By,
		ID:     s.ID,
		Parent: s.Parent,
		Body:   s.Text,
		Time:   t,
	}, nil
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
