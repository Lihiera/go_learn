package main

import (
	"fmt"
	"io"
)

type htmlReader struct {
	data  []byte
	index int
}

type LimitedReader struct {
	R     io.Reader
	limit int64
}

func NewReader(data string) *htmlReader {
	return &htmlReader{data: []byte(data), index: 0}
}

func (r *htmlReader) Read(p []byte) (n int, err error) {
	if r.index >= len(r.data) {
		return 0, io.EOF
	}
	n = copy(p, r.data[r.index:])
	r.index += n
	return
}

func (r *LimitedReader) Read(p []byte) (n int, err error) {
	if r.limit <= 0 {
		return 0, io.EOF
	}
	if int64(len(p)) > r.limit {
		p = p[:r.limit]
	}
	n, err = r.R.Read(p)
	r.limit -= int64(n)
	return
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &LimitedReader{R: r, limit: n}
}

func main() {
	var buf [1024]byte
	//var buf2 [1024]byte
	reader := NewReader(`
	<!DOCTYPE html>
	<html lang="en">
	<head>
	    <meta charset="UTF-8">
	    <meta name="viewport" content="width=device-width, initial-scale=1.0">
	    <title>Document</title>
	</head>
	<body>
	   hahaha
	</body>
	</html>
		`)
	newreader := LimitReader(reader, 20)
	n, _ := newreader.Read(buf[:])
	fmt.Println(n)
	//n2, _ := reader.Read(buf2[:])
	fmt.Println(reader.index)
	fmt.Println(string(buf[:n]))
}
