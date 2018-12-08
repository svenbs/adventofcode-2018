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

	s := bufio.NewScanner(f)

	var twos, threes int
	for s.Scan() {
		runes := map[rune]int{}
		for _, r := range []rune(s.Text()) {
			runes[r]++
		}
		var gotThree, gotTwo bool
		for _, r := range runes {
			if r == 2 {
				gotTwo = true
			} else if r == 3 {
				gotThree = true
			}
		}
		if gotTwo {
			twos++
		}
		if gotThree {
			threes++
		}
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Print(twos * threes)
}
