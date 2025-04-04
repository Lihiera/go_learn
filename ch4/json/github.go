package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const IssuesURL = "https://api.github.com/search/issues"

// type IssuesSearchResult struct {
// 	TotalCount int `json:"total_count"`
// 	Items      []*Issue
// }

// type Issue struct {
// 	Number    int
// 	HTMLURL   string `json:"html_url"`
// 	Title     string
// 	State     string
// 	User      *User
// 	CreatedAt time.Time `json:"created_at"`
// 	Body      string    // in Markdown format
// }

// type User struct {
// 	Login   string
// 	HTMLURL string `json:"html_url"`
// }

func main() {
	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	timecate := make(map[string][]*Issue)
	for _, item := range result.Items {
		timecate[timePass(item.CreatedAt)] = append(timecate[timePass(item.CreatedAt)], item)
	}
	for key := range timecate {
		fmt.Print(key + ":\n")
		for _, item := range timecate[key] {
			fmt.Printf("#%-5d %9.9s %.55s\n",
				item.Number, item.User.Login, item.Title)
			fmt.Println(item.CreatedAt)
		}

	}

}

func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

	// We must close resp.Body on all execution paths.
	// (Chapter 5 presents 'defer', which makes this simpler.)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

func timePass(t time.Time) string {
	now := time.Now()
	pass := now.Sub(t)
	if pass <= 30*24*time.Hour {
		return "less than a month"
	}
	if pass <= 365*24*time.Hour {
		return "less than a year"
	}
	return "more than a year"
}
