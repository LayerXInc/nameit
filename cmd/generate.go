package cmd

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"gopkg.in/yaml.v3"

	_ "embed"
)

func generateName(adjectives []string, nouns []string, prefix string, separator string, appendRandom bool, randomChars string, randomLength int) string {
	rngSource := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(rngSource)

	// Select a random adjective and noun
	adj := adjectives[rng.Intn(len(adjectives))]
	noun := nouns[rng.Intn(len(nouns))]

	// Combine with the configured separator
	name := adj + separator + noun

	// Prepend prefix
	if len(prefix) > 0 {
		name = prefix + separator + name
	}

	// Append random token
	if appendRandom && len(randomChars) > 0 && randomLength > 0 {
		// Generate a random token
		token := make([]byte, randomLength)
		for i := range token {
			token[i] = randomChars[rng.Intn(len(randomChars))]
		}
		name += separator + string(token)
	}

	return name
}

func generateNames(adjectives []string, nouns []string, prefix string, separator string, appendRandom bool, randomChars string, randomLength int, count int) []string {
	names := make([]string, 0, count)
	for i := 0; i < count; i++ {
		name := generateName(adjectives, nouns, prefix, separator, appendRandom, randomChars, randomLength)
		names = append(names, name)
	}
	return names
}

func outputNames(names []string) {
	switch strings.ToLower(outputFormat) {
	case "json":
		// Output as JSON array
		jsonData, err := json.MarshalIndent(names, "", "  ")
		if err != nil {
			fmt.Println("Error generating JSON:", err)
			return
		}
		fmt.Println(string(jsonData))

	case "yaml":
		// Output as YAML array
		yamlData, err := yaml.Marshal(names)
		if err != nil {
			fmt.Println("Error generating YAML:", err)
			return
		}
		fmt.Print(string(yamlData))

	case "text":
		// Output as plain text, one per line
		for _, name := range names {
			fmt.Println(name)
		}

	default:
		// Default to text format if unknown format specified
		fmt.Println("Unknown output format. Using text format:")
		for _, name := range names {
			fmt.Println(name)
		}
	}
}
