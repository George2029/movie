package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: movie <movie_name>")
		return
	}

	searchTerm := strings.ToLower(os.Args[1])
	dir := "/mnt/hdd/movies"

	err := os.Chdir(dir)
	if err != nil {
		fmt.Printf("Failed to change directory: %v\n", err)
		return
	}

	files, err := filepath.Glob("*")
	if err != nil {
		fmt.Printf("Failed to read directory: %v\n", err)
		return
	}

	matches := []string{}
	for _, file := range files {
		if strings.Contains(strings.ToLower(file), searchTerm) {
			matches = append(matches, file)
		}
	}

	switch len(matches) {
	case 0:
		fmt.Println("No matches found.")
	case 1:
		launchBook(matches[0])
	default:
		if len(matches) > 10 {
			fmt.Println("Too many results, be more specific.")
			return
		}
		fmt.Println("Multiple matches found. Please select a number:")
		for i, match := range matches {
			fmt.Printf("%d: %s\n", i+1, match)
		}

		reader := bufio.NewReader(os.Stdin)
		choice, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Failed to read input: %v\n", err)
			return
		}

		choiceNum, err := strconv.Atoi(strings.TrimSpace(choice))
		if err != nil || choiceNum < 1 || choiceNum > len(matches) {
			fmt.Println("Invalid choice.")
			return
		}

		launchBook(matches[choiceNum-1])
	}
}

func launchBook(movie string) {
	fmt.Printf("Launching movie: %s\n", movie)
	cmd := exec.Command("vlc", "-f", movie)
	err := cmd.Start()
	if err != nil {
		fmt.Printf("Failed to launch movie: %v\n", err)
	}
}
