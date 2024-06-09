package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

const postUrl string = "https://jsonplaceholder.typicode.com"
const commentUrl string = "https://jsonplaceholder.typicode.com"

/*
FetchAndPrintData takes its input and fetches the Post's with it and appends
the comments to it and then prints them according to the format.
*/
func FetchAndPrintData(userIDInt int, filterInput string) {

	posts, err := fetchPostsByUserID(userIDInt, postUrl)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	postsWithComments, err := appendCommentsToPosts(posts, commentUrl)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	filteredPosts := filterComments(postsWithComments, filterInput)

	printFormattedPosts(filteredPosts)
}

/*
CLI reads and returns the userID and filter provided by the user
*/
func CLI() {
	var userId int
	var filterInput string

	var cmd = &cobra.Command{
		Use: "otto",
		Run: func(cmd *cobra.Command, args []string) {
			FetchAndPrintData(userId, filterInput)
		},
	}

	cmd.Flags().IntVarP(&userId, "userId", "u", 0, "Specify a userID")

	cmd.MarkFlagRequired("userId") //handles Err by itself

	cmd.Flags().StringVarP(&filterInput, "filter", "f", "", "Specify a filterParameter")

	err := cmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}
