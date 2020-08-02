package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func handleError(err error) {
	if err != nil {
		log.Fatalln("Sorry, something went wrong: ", err)
	}
}

func getJokes(query string) {
	client := http.Client{}

	req, err := http.NewRequest("GET", fmt.Sprintf("https://www.reddit.com/r/dadjokes/search.json?q=%s&restrict_sr=on&sort=top&t=all&count=10", query), nil)
	req.Header.Set("User-Agent", "lukedadjokesgo")
	handleError(err)

	res, err := client.Do(req)
	handleError(err)

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	handleError(err)

	log.Println(string(body))
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

func main() {
	fmt.Println("What do you want to hear a joke about?")
	input := getUserInput()
	getJokes(input)
}
