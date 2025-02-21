package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const token = ""
const owner = "Lihiera"
const repo = "go_learn"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

type IssueRequest struct {
	Title  string   `json:"title"`
	Body   string   `json:"body"`
	Labels []string `json:"labels"`
}

type IssueUpdateRequest struct {
	Title  string   `json:"title"`
	Body   string   `json:"body"`
	State  string   `json:"state"` // open or closed
	Labels []string `json:"labels"`
}

func createIssue(title, body string) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/issues", owner, repo)
	issueRequest := IssueRequest{
		Title:  title,
		Body:   body,
		Labels: []string{"bug"}, // 添加标签
	}

	// 将请求体转换为 JSON 格式
	reqBody, err := json.Marshal(issueRequest)
	if err != nil {
		fmt.Println("Error marshalling request:", err)
		return
	}

	// 创建 HTTP 请求
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// 设置请求头
	req.Header.Set("Authorization", "token "+token)
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	if resp.StatusCode == http.StatusCreated {
		fmt.Println("Issue created successfully!")
		fmt.Println("Response:", string(respBody))
	} else {
		fmt.Println("Failed to create issue. Status:", resp.Status)
	}
}

func listIssues() {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/issues?state=open", owner, repo)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "token "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Failed to search issue. Status:", resp.StatusCode)
		return
	}

	var result []Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Println("Failed to decode. error:", err)
		return
	}
	for _, item := range result {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}

func updateIssue(issueNumber int, title, body, state string) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/issues/%d", owner, repo, issueNumber)
	issueUpdateRequest := IssueUpdateRequest{
		Title:  title,
		Body:   body,
		State:  state, // "open" 或 "closed"
		Labels: []string{"bug"},
	}

	// 将请求体转换为 JSON 格式
	reqBody, err := json.Marshal(issueUpdateRequest)
	if err != nil {
		fmt.Println("Error marshalling request:", err)
		return
	}

	// 创建 HTTP 请求
	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(reqBody))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// 设置请求头
	req.Header.Set("Authorization", "token "+token)
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	if resp.StatusCode == http.StatusOK {
		fmt.Println("Issue updated successfully!")
		fmt.Println("Response:", string(respBody))
	} else {
		fmt.Println("Failed to update issue. Status:", resp.Status)
	}
}

func main() {
	// 创建一个新的 issue
	//createIssue("Bug found in Go API", "There's a bug in the GitHub API client for Go.")
	//listIssues()
	updateIssue(1, "Bug fixed in Go API", "The bug has been fixed.", "closed")
}
