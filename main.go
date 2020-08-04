package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

type response struct {
	Data struct {
		Children []struct {
			Data struct {
				Selftext string `json:"selftext"`
				Title    string `json:"title"`
			} `json:"data"`
		} `json:"children"`
		Before string `json:"before"`
	} `json:"data"`
}

type joke struct {
	question string
	answer   string
}

func handleError(err error) {
	if err != nil {
		log.Fatalln("Sorry, something went wrong: ", err)
	}
}

func getUserInput() string {
	var input string
	for {
		scanner := bufio.NewReader(os.Stdin)
		input, _ = scanner.ReadString('\n')
		if len(input) > 1 {
			return strings.TrimSuffix(input, "\n")
		}
		fmt.Println("You must enter something...")
	}
}

func getJokes(query string) response {
	client := http.Client{}
	var err error

	req, err := http.NewRequest("GET", fmt.Sprintf("https://www.reddit.com/r/dadjokes/search.json?q=%s&restrict_sr=on&sort=top&t=all&count=10", query), nil)
	req.Header.Set("User-Agent", "lukedadjokesgo")
	handleError(err)

	res, err := client.Do(req)
	handleError(err)

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	handleError(err)

	jokes := response{}

	err = json.Unmarshal(body, &jokes)
	handleError(err)

	return jokes
}

func randomNumber(max int) int {
	rand.Seed(time.Now().UnixNano())
	return 0 + rand.Intn(max)
}

func main() {
	fmt.Println("What do you want to hear a joke about?")
	input := getUserInput()
	jokes := getJokes(input)
	randomJoke := jokes.Data.Children[randomNumber(len(jokes.Data.Children))]
	fmt.Println(randomJoke.Data.Title)
	fmt.Println("--------------------------------------")
	fmt.Println(randomJoke.Data.Selftext)
}
