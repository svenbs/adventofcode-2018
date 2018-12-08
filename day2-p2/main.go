package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	var ids []string
	for s.Scan() {
		ids = append(ids, s.Text())
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Strings(ids)
	for i := 0; i < len(ids); i++ {
		solution, ok := compare(ids[i], ids[i+1])
		if ok {
			fmt.Println(solution)
			return
		}
	}
}

func compare(a, b string) (string, bool) {
	var diffs, pos int
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			diffs++
			pos = i
		}
		if diffs > 1 {
			return "", false
		}
	}
	if diffs != 1 {
		return "", false
	}
	return a[:pos] + a[pos+1:], true
}
