package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	var goodsList = template.Must(template.New("goodsList").Parse(`
		<h1>goodslist</h1>
<table>
<tr style='text-align: left'>
  <th>Item</th>
  <th>Price</th>
</tr>
{{range $item, $price := .}}
<tr>
  <td>{{$item}}</td>
  <td>{{$price}}</td>
</tr>
{{end}}
</table>
`))
	if err := goodsList.Execute(w, db); err != nil {
		log.Fatal(err)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	if item == "" || price == "" {
		fmt.Fprintln(w, "Please provide an item and its price")
		return
	}
	dprice, _ := strconv.Atoi(price)
	if _, ok := db[item]; ok {
		fmt.Fprintf(w, "There has been %v, %v: %v\n", item, item, db[item])
		return
	}
	db[item] = dollars(dprice)
	fmt.Fprintf(w, "Create data successfully, %v: %v\n", item, price)
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	if item == "" || price == "" {
		fmt.Fprintln(w, "Please provide an item and its price")
		return
	}
	oldPrice, ok := db[item]
	if !ok {
		fmt.Fprintf(w, "no such item of %v\n", item)
		return
	}
	dprice, _ := strconv.Atoi(price)
	db[item] = dollars(dprice)
	fmt.Fprintf(w, "Update %v's price from %v to %v successfully\n", item, oldPrice, dollars(dprice))
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if item == "" {
		fmt.Fprintln(w, "Please provide an item")
		return
	}
	_, ok := db[item]
	if !ok {
		fmt.Fprintf(w, "no such item of %v\n", item)
		return
	}
	delete(db, item)
	fmt.Fprintf(w, "Delete %v successfully\n", item)
}

func main() {
	db := database{"shoes": 50, "socks": 5}
	mux := http.NewServeMux()
	mux.Handle("/list", http.HandlerFunc(db.list))
	mux.Handle("/price", http.HandlerFunc(db.price))
	mux.HandleFunc("/create", db.create)
	mux.HandleFunc("/update", db.update)
	mux.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}
