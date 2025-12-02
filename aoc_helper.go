package aoc2025

import (
	"bufio"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func LoadProblemFile(url string) []string {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env file")
	}

	sessionToken := os.Getenv("SESSION_ID")
	if sessionToken == "" {
		log.Fatal("SESSION_ID not found in .env file")
	}

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Cookie", "session="+sessionToken)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatal("error while fetching puzzle")
	}

	scanner := bufio.NewScanner(resp.Body)

	var fileInputs []string
	for scanner.Scan() {
		fileInputs = append(fileInputs, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error reading input:", err)
	}

	return fileInputs
}
