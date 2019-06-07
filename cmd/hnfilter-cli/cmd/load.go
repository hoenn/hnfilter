package cmd

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/hoenn/go-hn/pkg/hnapi"
	"github.com/hoenn/hnfilter/pkg/store"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	rps int
)

func init() {
	rootCmd.AddCommand(loadDBCmd)
	loadDBCmd.Flags().IntVarP(&rps, "rps", "r", 10, "How many requests to make per second")
}

var loadDBCmd = &cobra.Command{
	Use:   "LoadDB",
	Short: "Loads a postgres database with posts",
	Long:  "Loads a postgres database with hackernews who's hiring posts with rate limiting",
	Run: func(cmd *cobra.Command, args []string) {
		//Get user, iterate through posts to find monthly posts, grab the comments on each story, bulk load into db
		client := hnapi.NewHNClient()
		u, err := client.User("whoishiring")
		if err != nil {
			log.Fatal(errors.Wrap(err, "could not get user to load"))
		}
		fs := getUserTitleFilteredPosts(client, u, "Who is hiring?")
		//For each story, get all comments and bulk load those in database in separate transactions
		for _, s := range fs {
			cs, err := getAllCommentsForStory(client, s)
			if err != nil {
				log.Fatal(errors.Wrap(err, "could not get comments for stories"))
			}

			//Replace this with a transactional bulk loader
			for _, c := range cs {
				err := ds.AddComment(context.Background(), c)
				if err != nil {
					//IDs that fail can be reinserted from another job or retried
					log.Print(fmt.Sprintf("Unable to bulk load comments for story ID: %v", s.ID))
				}
			}
		}

	},
}

func getUserTitleFilteredPosts(client *hnapi.HNClient, u *hnapi.HNUser, filterStr string) []*hnapi.Story {
	var posts []string
	for _, s := range u.Submitted {
		i := strconv.Itoa(s)
		posts = append(posts, i)
	}
	//For each submitted story get the post
	var stories []*hnapi.Story
	for _, p := range posts {
		s, err := getStoryByID(p, client)
		if err != nil {
			//Not a story
			continue
		}
		stories = append(stories, s)
	}

	//Filter the stories by title
	var filteredStories []*hnapi.Story
	for _, s := range stories {
		if strings.Contains(s.Title, filterStr) {
			filteredStories = append(filteredStories, s)
		}
	}

	return filteredStories
}

func getAllCommentsForStory(client *hnapi.HNClient, story *hnapi.Story) ([]*store.Comment, error) {
	var comments []*store.Comment
	for _, c := range story.Kids {
		cc, err := getCommentByID(fmt.Sprint(c), client)
		if err != nil {
			return nil, err
		}
		comments = append(comments, cc)
	}
	return comments, nil
}

// getStoryByID returns a story from an id
func getStoryByID(id string, c *hnapi.HNClient) (*hnapi.Story, error) {
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

// getCommentByID returns a comment in db form from an id
func getCommentByID(id string, c *hnapi.HNClient) (*store.Comment, error) {
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
		Author: s.By,
		ID:     s.ID,
		Parent: s.Parent,
		Body:   s.Text,
		Time:   t,
	}, nil
}
