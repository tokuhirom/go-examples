package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type User struct {
	Screen_name string
}

type Twit struct {
	User User
	Text string
}

func main() {
	response, err := http.Get("http://twitter.com/statuses/public_timeline.json")
	if err == nil {
		body, _ := ioutil.ReadAll(response.Body)
		var twits []Twit
		json.Unmarshal(body, &twits)

		for _, twit := range twits {
			fmt.Printf("%s: %s\n", twit.User.Screen_name, twit.Text)
		}
	}
}
