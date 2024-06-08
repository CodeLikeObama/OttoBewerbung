package main

import (
	"strings"
)

func filterComments(posts []Post, filterParameter string) []Post {
	if filterParameter == "" {
		return posts
	}
	for i, post := range posts {
		var filteredComments []Comment
		for _, comment := range post.Comments {
			normalizedBody := strings.ReplaceAll(comment.Body, "\n", " ") //need to replace the \n with blank space in the body if you search it
			if strings.Contains(normalizedBody, filterParameter) {
				filteredComments = append(filteredComments, comment)
			}
		}
		posts[i].Comments = filteredComments
	}
	return posts
}