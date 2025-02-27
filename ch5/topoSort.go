package main

import (
	"fmt"
)

var prereqs = map[string][]string{
	"algorithms":     {"data structures"},
	"calculus":       {"linear algebra"},
	"linear algebra": {"calculus"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items map[string]bool)
	visitAll = func(items map[string]bool) {
		for key := range items {
			newkeys := make(map[string]bool)
			if !seen[key] {
				seen[key] = true
				for _, item := range m[key] {
					newkeys[item] = true
				}
				visitAll(newkeys)
				order = append(order, key)
			}
		}
	}
	// var keys []string
	// for key := range m {
	// 	keys = append(keys, key)
	// }
	// sort.Strings(keys)
	keys := make(map[string]bool)
	for key := range m {
		keys[key] = true
	}
	visitAll(keys)
	return order
}
