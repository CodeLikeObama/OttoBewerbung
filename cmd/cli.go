package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

/*
readUserID prompts the STDOUT so the user can enter the wanted userID and returns it as an Integer
*/
func readUserID() (int, error) {
	fmt.Print("Please enter userID: ")
	userIDReader := bufio.NewReader(os.Stdin)

	userIDInput, err := userIDReader.ReadString('\n')
	if err != nil {
		return 0, errors.New("an error occured while reading userID Input. Please try again")
	}
	userIDInput = strings.TrimSuffix(userIDInput, "\n")

	if userIDInput == "" {
		return 0, errors.New("please enter userID and try again")
	}

	userIDInt, err := strconv.Atoi(userIDInput)
	if err != nil {
		return 0, errors.New("userID must be an Integer, please try again")
	}

	return userIDInt, nil
}

/*
readFilterInput prompts the STDOUT so the user can enter the wanted filter parameter to filter the comments of posts and returns the filter parameter
*/
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

/*
CLI reads and returns the userID and filter provided by the user
*/
func CLI() (int, string) {
	var userId int
	var filterInput string

	var cmd = &cobra.Command{
		Use: "otto",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(args)
			fmt.Println(userId)
		},
	}

	cmd.PersistentFlags().IntVarP(&userId, "userId", "uid", 0, "Specify a userID")

	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(userId)
	userId, err := readUserID()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	filterInput = readFilterInput()

	return userId, filterInput
}
