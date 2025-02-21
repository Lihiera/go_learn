package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type link struct {
	Rel string `json:"img"`
}

func main() {
	number := os.Args[1]
	url := fmt.Sprintf("https://xkcd.com/%s/info.0.json", number)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching URL:", err)
	}
	defer resp.Body.Close()
	Decoder := json.NewDecoder(resp.Body)
	result := link{}
	if err := Decoder.Decode(&result); err != nil {
		fmt.Println("Error decoding JSON:", err)
	}
	fmt.Println(result.Rel)
}
