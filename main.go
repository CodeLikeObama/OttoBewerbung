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
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type Comment struct {
	PostID int    `json:"post_id"`
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
	// remove the delimeter from the string
	input = strings.TrimSuffix(input, "\n")
	inputInt, err := strconv.Atoi(input)
	fmt.Println(input)
	fmt.Println(fetchPosts(inputInt))
}

func fetchPosts(userID int) ([]Post, error) {
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
