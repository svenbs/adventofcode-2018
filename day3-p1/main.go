package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var fabric fabric
	s := bufio.NewScanner(f)
	for s.Scan() {
		var id, x, y, w, h int
		_, err := fmt.Sscanf(s.Text(), "#%d @ %d,%d: %dx%d", &id, &x, &y, &w, &h)
		if err != nil {
			log.Fatal(err)
		}

		fabric.addClaim(id, x, y, w, h)
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d", fabric.badlyClaimed())
}

type xy struct{ x, y int }

type fabric struct {
	m map[xy]int
}

func (f *fabric) addClaim(id, x, y, w, h int) {
	if f.m == nil {
		f.m = make(map[xy]int)
	}

	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			f.m[xy{x + i, y + j}]++
		}
	}
}

func (f *fabric) badlyClaimed() int {
	var count int
	for _, c := range f.m {
		if c > 1 {
			count++
		}
	}
	return count
}
