package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	// Check that a file name was provided
	if len(os.Args) < 2 {
		fmt.Println("Usage: secret-santa [file]")
		return
	}

	// Open the TSV file
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create a new scanner to read the file
	scanner := bufio.NewScanner(file)

	// Create a list of names
	var names []string
	// Create a map of exclusions
	exclusions := make(map[string]string)

	// Read the names and exclusions from the file
	for scanner.Scan() {
		// Split the line on the tab character
		fields := strings.Split(scanner.Text(), "\t")
		// Append the first field to the list of names
		names = append(names, fields[0])
		// Split the second field on the space character
		if len(fields) >= 2 {
			exclusions[fields[0]] = fields[1]
		} else {
			exclusions[fields[0]] = ""
		}
	}

	// Initialize the random number generator
	rand.Seed(time.Now().UnixNano())

	// Shuffle the names
	for i := range names {
		// Pick a random index that is not on the exclusion list
		j := i
		for {
			j = rand.Intn(len(names))
			if !strings.Contains(exclusions[names[i]], names[j]) {
				break
			}
		}
		names[i], names[j] = names[j], names[i]
	}

	// Print the name pairs
	for i := 0; i < len(names)-1; i++ {
		fmt.Printf("%s gives a gift to %s\n", names[i], names[i+1])
	}
	// The last person gives a gift to the first person
	fmt.Printf("%s gives a gift to %s\n", names[len(names)-1], names[0])
}
