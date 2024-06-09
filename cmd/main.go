package main

import (
	"fmt"
)

const postUrl string = "https://jsonplaceholder.typicode.com"
const commentUrl string = "https://jsonplaceholder.typicode.com"

/*
FetchAndPrintData calls the CLI() Function, takes its input and fetches the Post's with it and appends
the comments to it and then prints them according to the format.
*/
func FetchAndPrintData() {
	userIDInt, filterInput := CLI()
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

func main() {
	FetchAndPrintData()
}
