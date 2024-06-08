package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
)

func fetchPostsByUserID(userID int, url string) ([]Post, error) {
	var posts []Post

	var path = "/posts?userId=" + strconv.Itoa(userID)
	//handle invalid Input
	if userID > 10 || userID < 1 {
		return nil, errors.New("user ID out of range")
	}
	//input valid

	resp, err := http.Get(url + path)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()            //verbindung schließen, wenn alles gelesen ist
	body, err := io.ReadAll(resp.Body) //html Body wird gelesen
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &posts) //body wird gelesen und in ein Pointer zu Posts geschrieben
	if err != nil {
		return nil, err
	}

	return posts, err
}

func fetchCommentsByPostIDs(postID []int, url string) ([]Comment, error) {
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

	var comments []Comment
	err = json.Unmarshal(body, &comments)
	if err != nil {
		return nil, err
	}

	return comments, nil
}

func getPostIDs(post []Post) []int {
	var postIDS []int
	for _, post := range post {
		postIDS = append(postIDS, post.ID)
	}
	return postIDS
}

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

	return posts, err
}
