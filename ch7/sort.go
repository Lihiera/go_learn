package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sort"
	"strconv"
	"time"
)

type Item struct {
	Title     string
	ClickTime time.Time
	Id        int
}

type Table []*Item

func (t Table) Len() int           { return len(t) }
func (t Table) Less(i, j int) bool { return t[i].ClickTime.After(t[j].ClickTime) }
func (t Table) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }

func main() {
	file, _ := os.Open("table.json")
	defer file.Close()
	movies := new(Table)
	if err := json.NewDecoder(file).Decode(movies); err != nil {
		fmt.Println(err)
		fmt.Errorf("error when decoding: %v", err)
		return
	}
	fmt.Println("hello")

	servePage := func(w http.ResponseWriter, r *http.Request) {
		sort.Sort(*movies)
		w.Header().Set("Content-Type", "text/html")
		html := `
		<!DOCTYPE html>
		<html>
		<head>
			<meta charset="UTF-8">
			<title>Go Click Tracker</title>
		</head>
		<body>
			<h1>Item List</h1>
			<ul>`
		for _, item := range *movies {
			html += fmt.Sprintf(`<li>%s - 点击时间: %s <a href="/click?id=%d">点击</a></li>`,
				item.Title, item.ClickTime.Format("2006-01-02 15:04:05"), item.Id)
		}
		html += `</ul>
		</body>
		</html>`

		w.Write([]byte(html))
	}
	clickPage := func(w http.ResponseWriter, r *http.Request) {
		itemID, _ := strconv.Atoi(r.URL.Query().Get("id"))
		for _, item := range *movies {
			if item.Id == itemID {
				item.ClickTime = time.Now()
				break
			}
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	http.HandleFunc("/", servePage)
	http.HandleFunc("/click", clickPage)
	fmt.Println("Server running on http://localhost:8080")
	// 监听
	http.ListenAndServe(":8000", nil)

}
