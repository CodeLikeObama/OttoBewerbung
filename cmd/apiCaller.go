package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
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

/*
fetchPostsByUserID fetches the posts by a specific userID and returns the corresponding Post's
*/
func fetchPostsByUserID(userID int, url string) ([]Post, error) {
	var posts []Post

	var path = "/posts?userId=" + strconv.Itoa(userID)
	//handle invalid input
	if userID > 10 || userID < 1 {
		return nil, errors.New("user ID out of range")
	}

	resp, err := http.Get(url + path)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &posts)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

/*
fetchCommentsByPostIDs fetches the comments for a slice of postIDs and returns them as Comment's
*/
func fetchCommentsByPostIDs(postID []int, url string) ([]Comment, error) {
	var comments []Comment
	query := "/comments?"

	for _, postID := range postID {
		query += "&postId=" + strconv.Itoa(postID)
	}

	resp, err := http.Get(url + query)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &comments)
	if err != nil {
		return nil, err
	}

	return comments, nil
}

/*
getPostIDs gets all the IDs of a Post and returns them as a slice of Integers
*/
func getPostIDs(post []Post) []int {
	var postIDS []int

	for _, post := range post {
		postIDS = append(postIDS, post.ID)
	}
	return postIDS
}

/*
appendCommentsToPosts takes posts and a URl and matches the comments to the corresponding Post's and returns them
*/
func appendCommentsToPosts(posts []Post, url string) ([]Post, error) {
	postIDs := getPostIDs(posts)
	comments, err := fetchCommentsByPostIDs(postIDs, url)

	if err != nil {
		return nil, err
	}

	commentsByPostID := make(map[int][]Comment)
	for _, comment := range comments {
		commentsByPostID[comment.PostID] = append(commentsByPostID[comment.PostID], comment)
	}

	for i, post := range posts {
		posts[i].Comments = commentsByPostID[post.ID]
	}

	return posts, nil
}
