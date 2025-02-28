package main

import (
	"fmt"
)

var prereqs = map[string][]string{
	"algorithms":      {"data structures"},
	"calculus":        {"linear algebra"},
	"higher algebra":  {"calculus"},
	"complex algebra": {"calculus", "algorithms"},
	"linear algebra":  {"calculus"},
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
	cantStudy := make(map[string]bool)
	var visitAll func(items map[string]bool, chain []string)
	visitAll = func(items map[string]bool, chain []string) {
		hasCycle := false
		for key := range items {
			if cantStudy[key] {
				for _, item := range chain {
					cantStudy[item] = true
				}
				continue
			}
			for _, appended := range chain {
				if appended == key {
					hasCycle = true
					break
				}
			}
			if hasCycle {
				fmt.Print("Detected cycle: ")
				for _, item := range append(chain, key) {
					cantStudy[item] = true
					fmt.Printf("%q -> ", item)
				}
				fmt.Println()
				return
			}
			newkeys := make(map[string]bool)
			if !seen[key] {
				for _, item := range m[key] {
					newkeys[item] = true
				}
				visitAll(newkeys, append(chain, key))
				if !cantStudy[key] {
					seen[key] = true
					order = append(order, key)
				}
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
	visitAll(keys, nil)
	return order
}
