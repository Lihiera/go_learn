package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

var count = make(map[string]int)

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
	for _, link := range visit(nil, doc, "input") {
		fmt.Println(link)
	}
	for key, value := range count {
		fmt.Printf("%s: %v\n", key, value)
	}
}

func visit(links []string, n *html.Node, cate string) []string {
	if n == nil {
		return links
	}
	if n.Type == html.ElementNode && n.Data == cate {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	if n.Type == html.ElementNode {
		count[n.Data]++
	}
	// for c := n.FirstChild; c != nil; c = c.NextSibling {
	// 	links = visit(links, c)
	// }
	links = visit(links, n.FirstChild, cate)
	links = visit(links, n.NextSibling, cate)
	return links
}

// func visit2(n *html.Node) {
// 	if n == nil {
// 		return
// 	}
// 	if n.Type == html.ElementNode {
// 		count[n.Data]++
// 	}
// 	visit2(n.FirstChild)
// 	visit2(n.NextSibling)
// }
