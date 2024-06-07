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
	userIDReader := bufio.NewReader(os.Stdin)
	// ReadString will block until the delimiter is entered
	userIDInput, err := userIDReader.ReadString('\n') //delim is enter (line break)
	//TODO Maybe handle empty input
	if err != nil {
		fmt.Println("An error occured while reading userID Input. Please try again", err)
		return
	}
	userIDInput = strings.TrimSuffix(userIDInput, "\n")

	fmt.Println("Please enter a Filter parameter: ")
	filterReader := bufio.NewReader(os.Stdin)
	filterInput, err := filterReader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading filter parameters. Please try again", err)
	}
	filterInput = strings.TrimSuffix(filterInput, "\n")

	inputInt, err := strconv.Atoi(userIDInput)
	posts, _ := fetchPostsByUserID(inputInt)
	postsWithComments, _ := AppendCommentToPost(posts)
	filteredPosts := filterComments(postsWithComments, filterInput)
	
	printFormattedPosts(filteredPosts)

	/*


		//TODO Tidy up below
		//postsWithComments, _ := AppendCommentToPost(posts)
		//postIDs := getPostIDs(posts)
		//comments, _ := fetchCommentsByPostID(postIDs)
		fmt.Println(input)
		//fmt.Println(posts)
		//fmt.Println(postIDs)
		//fmt.Println(comments)
		//fmt.Println(len(comments))
		//fmt.Println(AppendCommentToPost(posts))
		postsWithComments, _ := AppendCommentToPost(posts)
		printFormattedPosts(postsWithComments)
	*/
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

/*
func filterComments(posts []Post, filterParameter string) []Post {
	for i, post := range posts {
		var filteredComments []Comment
		for _, comment := range post.Comments {
			if strings.Contains(comment.Body, filterParameter) {
				filteredComments = append(filteredComments, comment)
			}
		}
		posts[i].Comments = filteredComments
	}
	return posts
}

*/

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
