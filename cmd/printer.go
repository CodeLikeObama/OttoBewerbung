package main

import (
	"fmt"
)

/*
printFormattedPosts takes a slice of Post's and prints their attributes nicely formatted to the STDOUT
*/
func printFormattedPosts(posts []Post) {
	for _, post := range posts {
		fmt.Printf("UserID: %d\n", post.UserID)
		fmt.Printf("ID: %d\n", post.ID)
		fmt.Printf("Title: %s\n", post.Title)
		fmt.Printf("Body: %s\n", post.Body)
		fmt.Println("Comments: ")
		for _, comment := range post.Comments {
			fmt.Printf("\tPostID: %d\n", comment.PostID)
			fmt.Printf("\tID: %d\n", comment.ID)
			fmt.Printf("\tName: %s\n", comment.Name)
			fmt.Printf("\tEmail: %s\n", comment.Email)
			fmt.Printf("\tBody: %s\n", comment.Body)
			fmt.Println()
		}
		fmt.Println("----------------------------------------")
	}
}
