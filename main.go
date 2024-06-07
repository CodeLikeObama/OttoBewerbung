package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
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

func main() {
	fmt.Print("Please enter userID: ")
	reader := bufio.NewReader(os.Stdin)
	// ReadString will block until the delimiter is entered
	input, err := reader.ReadString('\n') //delim is enter (line break)
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return
	}
	//TODO Maybe handle empty input

	// remove the delimeter from the string
	input = strings.TrimSuffix(input, "\n")
	inputInt, err := strconv.Atoi(input)
	posts, _ := fetchPostsByUserID(inputInt)
	//postIDs := getPostIDs(posts)
	//comments, _ := fetchCommentsByPostID(postIDs)
	fmt.Println(input)
	//fmt.Println(posts)
	//fmt.Println(postIDs)
	//fmt.Println(comments)
	//fmt.Println(len(comments))
	fmt.Println(AppendCommentToPost(posts))
}

func fetchPostsByUserID(userID int) ([]Post, error) {
	//handle invalid Input
	if userID > 10 {
		return nil, errors.New("User ID out of range")
	}
	//input valid
	userIDString := strconv.Itoa(userID)
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts?userId=" + userIDString)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()            //verbindung schließen, wenn alles gelesen ist
	body, err := io.ReadAll(resp.Body) //html Body wird gelesen
	if err != nil {
		return nil, err
	}
	var posts []Post
	err = json.Unmarshal(body, &posts) //body wird gelesen und in ein Pointer zu Posts geschrieben
	if err != nil {
		return nil, err
	}

	return posts, err
}

func getPostIDs(post []Post) []int {
	var postIDS []int
	for _, post := range post {
		postIDS = append(postIDS, post.ID)
	}
	return postIDS
}

/*
func fetchCommentsByPostID(postIDs []int) ([]Comment, error) {
	var allComments []Comment

	for _, postID := range postIDs {
		url := fmt.Sprintf("https://jsonplaceholder.typicode.com/comments?postId=%d", postID)
		resp, err := http.Get(url)
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

		allComments = append(allComments, comments...) //... variadic function so that each slice what we append can be andy size  and not a fixed size
	}

	return allComments, nil
}
*/

func fetchCommentsByPostID(postID int) ([]Comment, error) {
	postIDString := strconv.Itoa(postID)
	resp, err := http.Get("https://jsonplaceholder.typicode.com/comments?postId=" + postIDString)
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

func AppendCommentToPost(posts []Post) ([]Post, error) {
	for i, post := range posts {
		comments, err := fetchCommentsByPostID(post.ID)
		if err != nil {
			return nil, err
		}
		posts[i].Comments = comments
	}
	return posts, nil
}
