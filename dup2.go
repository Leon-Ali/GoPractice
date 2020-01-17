package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, "")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			}
			countLines(f, counts, arg)
			f.Close()
		}
	}
	for fName, inMp := range counts {
		for line, n := range inMp {
			if n > 1 {
				fmt.Printf("%s\t%d\t%s\n", fName, n, line)
			}
		}
		
	}

}

func countLines(f *os.File, counts map[string]map[string]int, fileName string) {
	input := bufio.NewScanner(f)
	innerMap := make(map[string]int)
	for input.Scan() {
		innerMap[input.Text()]++
		counts[fileName] = innerMap	
	}
	
}
