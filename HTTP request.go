package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"body"`
}

func main() {
	//URL to send the GET request
	url := "https://jsonplaceholder.typicode.com/posts/1"

	//Send the GET request
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("HTTP GET request failed: %s\n", err)
		return
	}

	//Ensure the response body is closed after we are done with it
	defer response.Body.Close()

	//Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Print("failed to read the response body: %s\n", err)
		return
	}

	//parse the JSON response into a Post struct
	var post Posterr = json.Unmarshal(body, &post)
	if err != nil {
		fmt.Printf("Failed to parse JSON: %s\n", err)
		return
	}
	//print the parsed post
	fmt.Printf("UserID: %d\n", post.UserID)
	fmt.Printf("ID: %d", post.ID)
	fmt.Printf("Title: %d", post.Title)
	fmt.Printf("")
}
