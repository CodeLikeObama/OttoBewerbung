package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readUserID() (int, error) {
	fmt.Print("Please enter userID: ")
	userIDReader := bufio.NewReader(os.Stdin)
	// ReadString will block until the delimiter is entered
	userIDInput, err := userIDReader.ReadString('\n') //delim is enter (line break)
	if err != nil {
		return 0, errors.New("an error occured while reading userID Input. Please try again")
	}
	userIDInput = strings.TrimSuffix(userIDInput, "\n")
	//handling empty input
	if userIDInput == "" {
		return 0, errors.New("please enter userID and try again")
	}

	userIDInt, err := strconv.Atoi(userIDInput)
	if err != nil {
		return 0, errors.New("userID must be an Integer, please try again")
	}

	return userIDInt, nil
}

func readFilterInput() string {
	fmt.Println("Please enter a Filter parameter: ")
	filterReader := bufio.NewReader(os.Stdin)
	filterInput, err := filterReader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading filter parameters. Please try again", err)
	}
	filterInput = strings.TrimSuffix(filterInput, "\n")

	return filterInput
}

func CLI() (int, string) {

	userId, err := readUserID()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	filterInput := readFilterInput()

	return userId, filterInput
}
