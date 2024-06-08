package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestFetchPostsByUserID(t *testing.T) {
	mockPosts := []Post{
		{UserID: 1, ID: 1, Title: "Test Post 1", Body: "Body of test post 1"},
		{UserID: 1, ID: 2, Title: "Test Post 2", Body: "Body of test post 2"},
	}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check the request URL
		expectedURL := "/posts?userId=1"
		if r.URL.String() != expectedURL {
			t.Errorf("Expected URL: %s, got: %s", expectedURL, r.URL.String())
		}

		// Marshal posts to JSON
		resp, _ := json.Marshal(mockPosts)

		// Write response
		w.WriteHeader(http.StatusOK)
		w.Write(resp)
	}))
	defer server.Close()

	// Call the function to test
	posts, err := fetchPostsByUserID(1, server.URL)

	// Check for errors
	if err != nil {
		t.Errorf("fetchPostsByUserID returned an error: %v", err)
	}

	if !reflect.DeepEqual(posts, mockPosts) {
		t.Errorf("fetchPostsByUserID returned unexpected result: got %v, want %v", posts, mockPosts)
	}
}

func TestFetchCommentsByPostIDs(t *testing.T) {

	mockComments := []Comment{
		{PostID: 1, ID: 1, Name: "Test Comment 1", Email: "test1@example.com", Body: "Body of test comment 1"},
		{PostID: 2, ID: 2, Name: "Test Comment 2", Email: "test2@example.com", Body: "Body of test comment 2"},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		expectedURL := "/comments?&postId=1&postId=2"
		if r.URL.String() != expectedURL {
			t.Errorf("Expected URL: %s, got: %s", expectedURL, r.URL.String())
		}

		resp, _ := json.Marshal(mockComments)

		w.WriteHeader(http.StatusOK)
		w.Write(resp)
	}))
	defer server.Close()

	comments, err := fetchCommentsByPostIDs([]int{1, 2}, server.URL)

	if err != nil {
		t.Errorf("fetchCommentsByPostIDs returned an error: %v", err)
	}

	if !reflect.DeepEqual(comments, mockComments) {
		t.Errorf("fetchCommentsByPostIDs returned unexpected result: got %v, want %v", comments, mockComments)
	}
}

func TestFetchPostsByUserIDRange(t *testing.T) {
	testURL := "thisisatestURL"

	_, err := fetchPostsByUserID(-2, testURL)
	if err == nil {
		t.Errorf("fetchPostsByUserID did not return an error for invalid userID")
	}

	expectedError := "user ID out of range"
	if err != nil && err.Error() != expectedError {
		t.Errorf("fetchPostsByUserID returned an unexpected error: got %v, want %v", err.Error(), expectedError)
	}

	_, err2 := fetchPostsByUserID(12, testURL)
	if err2 == nil {
		t.Errorf("fetchPostsByUserID did not return an error for invalid userID")
	}

	expectedError2 := "user ID out of range"
	if err != nil && err.Error() != expectedError2 {
		t.Errorf("fetchPostsByUserID returned an unexpected error: got %v, want %v", err2.Error(), expectedError2)
	}
}
