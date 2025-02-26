package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	resp, err := http.Get("https://www.lsta.media.kyoto-u.ac.jp/internal-wiki/index.php")
	if err != nil {
		fmt.Printf("Cant access the site: %v\n", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("API 请求失败，状态码: %d\n", resp.StatusCode)
		return
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	visitText(doc)
}

func visitText(n *html.Node) {
	if n == nil {
		return
	}
	if n.Type == html.ElementNode && (n.Data == "script" || n.Data == "style") {
		return
	}
	if n.Type == html.TextNode {
		fmt.Println(n.Data)
	}
	visitText(n.FirstChild)
	visitText(n.NextSibling)
}
