package main

import (
	"strings"
)

/*
filterComments takes a slice of Post's and a filter parameter and filters the Comment's of the posts according to it
and returns the new Post's with filtered Comment's
*/
func filterComments(posts []Post, filterParameter string) []Post {
	if filterParameter == "" {
		return posts
	}

	normalizedFilterParameter := strings.ReplaceAll(filterParameter, "\n", " ")
	for i, post := range posts {
		var filteredComments []Comment
		for _, comment := range post.Comments {
			normalizedBody := strings.ReplaceAll(comment.Body, "\n", " ")
			if strings.Contains(normalizedBody, normalizedFilterParameter) {
				filteredComments = append(filteredComments, comment)
			}
		}
		posts[i].Comments = filteredComments
	}
	return posts
}
