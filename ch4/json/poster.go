package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const apiKey = "1d948c9f"

type Movie struct {
	Title  string
	Poster string
}

func main() {
	url := fmt.Sprintf("https://www.omdbapi.com/?t=%s&apikey=%s", "The Avengers", apiKey)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Errorf("API 请求失败，状态码: %d", resp.StatusCode)
		return
	}
	result := Movie{}
	contents := json.NewDecoder(resp.Body)
	if err := contents.Decode(&result); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Title: " + result.Title)
	fmt.Println("Link: " + result.Poster)
	if result.Poster == "N/A" {
		return
	}
	resp2, err := http.Get(result.Poster)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp2.Body.Close()
	if resp2.StatusCode != http.StatusOK {
		fmt.Printf("下载失败，状态码: %d", resp2.StatusCode)
		return
	}
	filename := result.Title + ".jpg"
	file, err := os.Create(filename)
	if err != nil {
		return
	}
	defer file.Close()
	_, err = io.Copy(file, resp2.Body)
	if err != nil {
		fmt.Println(err)
	}
}
