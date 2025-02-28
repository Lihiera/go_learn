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
	//nodes := ElementsByTagName(doc, "a", "img")
	nodes := ElementsByTagName(doc, "script")
	fmt.Println(len(nodes))

}

func ElementsByTagName(n *html.Node, name ...string) []*html.Node {
	var result []*html.Node
	var visit func(n *html.Node)
	visit = func(n *html.Node) {
		if n.Type == html.ElementNode {
			for _, nm := range name {
				if n.Data == nm {
					result = append(result, n)
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			visit(c)
		}
	}
	visit(n)
	return result
}
