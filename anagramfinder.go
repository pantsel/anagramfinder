package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"time"
)

func alphabetize(word string) string {
	if len(word) <= 0 {
		return ""
	}

	s := strings.Split(word, "")
	sort.Strings(s)
	alphabetized := strings.TrimSpace(strings.Join(s, ""))

	return alphabetized
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) > 0 {
			lines = append(lines, text)
		}
	}
	return lines, scanner.Err()
}

func getPath() string {
	pathPtr := flag.String("path", "", "path to the words file")
	flag.Parse()
	fmt.Println("path:", *pathPtr)
	if len(*pathPtr) <= 0 {
		log.Fatalf("You need to specify a path to the wordlist")
	}
	return *pathPtr
}

func printAnagrams(groups map[string][]string) {
	for k, v := range groups {
		_ = k
		if len(v) > 1 {
			fmt.Println(strings.Join(v, " "))
		}
	}
}

func mapWords(words []string) map[string][]string {

	groups := map[string][]string{}

	for i, s := range words {
		_ = i
		alphabetized := alphabetize(s)
		groups[alphabetized] = append(groups[alphabetized], s)
	}

	return groups
}

func main() {
	start := time.Now()
	path := getPath()
	words, err := readLines(path)

	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	groups := mapWords(words)
	printAnagrams(groups)

	elapsed := time.Since(start)
	log.Printf("-----------------------------------------\n")
	log.Printf("Execution took %s", elapsed)

}
