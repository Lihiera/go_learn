package outline

import (
	"golang.org/x/net/html"
)

// func main() {
// 	resp, err := http.Get("https://www.lsta.media.kyoto-u.ac.jp/internal-wiki/index.php")
// 	if err != nil {
// 		fmt.Printf("Cant access the site: %v\n", err)
// 	}
// 	defer resp.Body.Close()
// 	if resp.StatusCode != http.StatusOK {
// 		fmt.Printf("API 请求失败，状态码: %d\n", resp.StatusCode)
// 		return
// 	}
// 	doc, err := html.Parse(resp.Body)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
// 		os.Exit(1)
// 	}
// 	var depth int
// 	inlinestart := func(n *html.Node) {
// 		if n.Type == html.ElementNode {
// 			//fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
// 			fmt.Printf("%*s<%s", depth*2, "", n.Data)
// 			for _, a := range n.Attr {
// 				fmt.Printf(" %s=%q", a.Key, a.Val)
// 			}
// 			fmt.Println(">")
// 			depth++
// 		}
// 		if n.Type == html.TextNode {
// 			fmt.Printf("%*sText:%q\n", depth*2, "", n.Data)
// 			// fmt.Println("Text Node")
// 		}
// 		if n.Type == html.CommentNode {
// 			fmt.Printf("%*sComment:%q\n", depth*2, "", n.Data)
// 			// fmt.Println("Text Node")
// 		}
// 	}
// 	inlineend := func(n *html.Node) {
// 		if n.Type == html.ElementNode {
// 			depth--
// 			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
// 		}
// 	}
// 	forEachNode(doc, inlinestart, inlineend)
// }

func ForEachNode(n *html.Node, pre, post func(n *html.Node)) {
	// if n.FirstChild == nil {
	// 	briefElement(n)
	// 	return
	// }
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ForEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

// func startElement(n *html.Node) {
// 	if n.Type == html.ElementNode {
// 		//fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
// 		fmt.Printf("%*s<%s", depth*2, "", n.Data)
// 		for _, a := range n.Attr {
// 			fmt.Printf(" %s=%q", a.Key, a.Val)
// 		}
// 		fmt.Println(">")
// 		depth++
// 	}
// 	if n.Type == html.TextNode {
// 		fmt.Printf("%*sText:%q\n", depth*2, "", n.Data)
// 		// fmt.Println("Text Node")
// 	}
// 	if n.Type == html.CommentNode {
// 		fmt.Printf("%*sComment:%q\n", depth*2, "", n.Data)
// 		// fmt.Println("Text Node")
// 	}
// }
// func endElement(n *html.Node) {
// 	if n.Type == html.ElementNode {
// 		depth--
// 		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
// 	}
// }

// func briefElement(n *html.Node) {
// 	if n.Type == html.ElementNode {
// 		fmt.Printf("%*s<%s", depth*2, "", n.Data)
// 		for _, a := range n.Attr {
// 			fmt.Printf(" %s=%q", a.Key, a.Val)
// 		}
// 		fmt.Print("/>\n")
// 	}
// }
