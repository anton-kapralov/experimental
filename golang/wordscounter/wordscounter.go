package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	wordMap := make(map[string]int)
	for _, filename := range os.Args[1:] {
		file, err := os.Open(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			continue
		}
		countWords(file, wordMap)
	}

	for word, n := range wordMap {
		fmt.Printf("%s\t%d\n", word, n)
	}
}

func countWords(f *os.File, wordMap map[string]int) {
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		wordMap[scanner.Text()]++
	}
}
