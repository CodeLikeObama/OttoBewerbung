package main

import (
	"fmt"
)

type Post struct {
	UserID   int    `json:"userId"`
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Body     string `json:"body"`
	Comments []Comment
}

type Comment struct {
	PostID int    `json:"postId"`
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}

var postUrl string = "https://jsonplaceholder.typicode.com"

var commentUrl string = "https://jsonplaceholder.typicode.com"

func main() {

	FetchDataAndPrint()

}

func FetchDataAndPrint() {
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
