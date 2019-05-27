package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	rps uint32
)

func init() {
	rootCmd.AddCommand(loadDBCmd)
}

var loadDBCmd = &cobra.Command{
	Use:   "LoadDB",
	Short: "Loads a postgres database with posts",
	Long:  "Loads a postgres database with hackernews who's hiring posts with rate limiting",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hello")
	},
}
