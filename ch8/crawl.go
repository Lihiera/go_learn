package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"

	"golang.org/x/net/html"
	"gopl.io/ch5/links"
)

var domain string
var count int

func crawl(link item) items {
	if link.dep > 0 {
		return items{nil, 0}
	}

	fmt.Println(link.url, "/t", link.dep)
	downloadUrl(link.url)
	urls, err := links.Extract(link.url)
	if err != nil {
		log.Print(err)
	}
	return items{urls, link.dep + 1}
}

type items struct {
	urls []string
	dep  int
}

type item struct {
	url string
	dep int
}

func main() {
	worklist := make(chan items) // lists of URLs, may have duplicates
	unseenLinks := make(chan item)
	var wg sync.WaitGroup
	parsedURL, err := url.Parse(os.Args[1])
	if err != nil {
		log.Fatal(err)
	} // de-duplicated URLs
	domain = parsedURL.Hostname()
	// Add command-line arguments to worklist.
	go func() { worklist <- items{os.Args[1:], 0} }()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() { worklist <- foundLinks }()
			}
		}()
	}
	go func() {
		wg.Wait()
		close(unseenLinks)
		close(worklist)
	}()
	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	wg.Add(1)
	os.Mkdir(domain, 0755)
	os.Chdir(domain)
	os.Exit(0)
	for links := range worklist {
		for _, url := range links.urls {
			if !seen[url] {
				seen[url] = true
				wg.Add(1)
				unseenLinks <- item{url, links.dep}

			}

		}
		wg.Done()
	}
}

func isHost(link string) bool {
	parsedURL, err := url.Parse(link)
	if err != nil {
		return false // URL 解析失败，视为不在域名下
	}
	host := parsedURL.Hostname() // 只取主机名，不包含端口

	// 判断是否是 domain 或其子域名
	return host == domain
}

func downloadUrl(link string) {
	fullUrl := link
	if !strings.HasPrefix(link, "http://") && !strings.HasPrefix(link, "https://") {
		fullUrl = "https://" + link
	}
	resp, err := http.Get(fullUrl)
	if err != nil {
		log.Print(err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Print("failed to fetch page, status code: %d", resp.StatusCode)
		return
	}
	file, err := os.Create(fmt.Sprintf("%d.html", count))
	if err != nil {
		log.Print(err)
		return
	}
	defer file.Close()
	count++
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		log.Print(err)
		return
	}
}

func extract(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // ignore bad URLs
				}
				if !isHost(link.String()) {
					continue
				}
				links = append(links, link.String())
			}
		}
		if n.Type == html.ElementNode && n.Data == "link" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				relativePath := resp.Request.URL.Path(a.Val)
				if err != nil {
					continue // ignore bad URLs
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}
