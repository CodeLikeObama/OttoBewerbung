package main

import (
	"reflect"
	"testing"
)

func TestFilterCommentsValid(t *testing.T) {
	mockPosts := []Post{
		{
			UserID: 1,
			ID:     1,
			Title:  "Test Post",
			Body:   "Body of test post",
			Comments: []Comment{
				{PostID: 1, ID: 1, Name: "Test Comment 1", Email: "test1@example.com", Body: "this Should appear in the Filter"},
				{PostID: 1, ID: 2, Name: "Test Comment 2", Email: "test2@example.com", Body: "this Shouldnt appear in the Filter"},
				{PostID: 1, ID: 2, Name: "Test Comment 2", Email: "test2@example.com", Body: "this is a comment with\nin it"},
			},
		},
	}

	expectedPost := []Post{
		{
			UserID: 1,
			ID:     1,
			Title:  "Test Post",
			Body:   "Body of test post",
			Comments: []Comment{
				{PostID: 1, ID: 1, Name: "Test Comment 1", Email: "test1@example.com", Body: "this Should appear in the Filter"},
			},
		},
	}
	filteredPosts := filterComments(mockPosts, "this Should appear in the Filter")

	if !reflect.DeepEqual(filteredPosts, expectedPost) {
		t.Errorf("filterComments returned unexpected results: got %v want %v", filteredPosts, expectedPost)
	}

}

// TODO Maybe test input with \n ?
func TestFilterCommentsEdgeCase(t *testing.T) {
	mockPosts := []Post{
		{
			UserID: 1,
			ID:     1,
			Title:  "Test Post",
			Body:   "Body of test post",
			Comments: []Comment{
				{PostID: 1, ID: 1, Name: "Test Comment 1", Email: "test1@example.com", Body: "this Should appear in the Filter"},
				{PostID: 1, ID: 2, Name: "Test Comment 2", Email: "test2@example.com", Body: "this Shouldnt appear in the Filter"},
				{PostID: 1, ID: 2, Name: "Test Comment 2", Email: "test2@example.com", Body: "this is a comment with\nin it"},
			},
		},
	}

	expectedPost := []Post{
		{
			UserID: 1,
			ID:     1,
			Title:  "Test Post",
			Body:   "Body of test post",
			Comments: []Comment{
				{PostID: 1, ID: 2, Name: "Test Comment 2", Email: "test2@example.com", Body: "this is a comment with\nin it"},
			},
		},
	}

	filteredPosts := filterComments(mockPosts, "this is a comment with in it")

	if !reflect.DeepEqual(filteredPosts, expectedPost) {
		t.Errorf("filterComments returned unexpected results: got %v want %v", filteredPosts, expectedPost)
	}

}

func TestFilterCommentsEmptyInput(t *testing.T) {
	mockPosts := []Post{
		{
			UserID: 1,
			ID:     1,
			Title:  "Test Post",
			Body:   "Body of test post",
			Comments: []Comment{
				{PostID: 1, ID: 1, Name: "Test Comment 1", Email: "test1@example.com", Body: "this Should appear in the Filter"},
				{PostID: 1, ID: 2, Name: "Test Comment 2", Email: "test2@example.com", Body: "this Shouldnt appear in the Filter"},
				{PostID: 1, ID: 2, Name: "Test Comment 2", Email: "test2@example.com", Body: "this is a comment with\nin it"},
			},
		},
	}

	unfliteredPosts := filterComments(mockPosts, "")

	if !reflect.DeepEqual(unfliteredPosts, mockPosts) {
		t.Errorf("filterComments returned unexpected results: got %v want %v", unfliteredPosts, mockPosts)
	}

}

func TestFilterCommentsNonexistentInput(t *testing.T) {
	mockPosts := []Post{
		{
			UserID: 1,
			ID:     1,
			Title:  "Test Post",
			Body:   "Body of test post",
			Comments: []Comment{
				{PostID: 1, ID: 1, Name: "Test Comment 1", Email: "test1@example.com", Body: "this Should appear in the Filter"},
				{PostID: 1, ID: 2, Name: "Test Comment 2", Email: "test2@example.com", Body: "this Shouldnt appear in the Filter"},
				{PostID: 1, ID: 2, Name: "Test Comment 2", Email: "test2@example.com", Body: "this is a comment with\nin it"},
			},
		},
	}

	expectedPost := []Post{
		{
			UserID: 1,
			ID:     1,
			Title:  "Test Post",
			Body:   "Body of test post",
		},
	}

	filteredPosts := filterComments(mockPosts, "this text is not included in any comment")

	if !reflect.DeepEqual(filteredPosts, expectedPost) {
		t.Errorf("filterComments returned unexpected results: got %v want %v", filteredPosts, expectedPost)
	}
}
