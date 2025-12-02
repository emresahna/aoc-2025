package aoc2025

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var client *http.Client

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env file")
	}

	client = &http.Client{}
}

func LoadInput(url string) ([]string, error) {
	sessionToken := os.Getenv("SESSION_ID")
	if sessionToken == "" {
		return nil, fmt.Errorf("SESSION_ID not found in .env file")
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Cookie", "session="+sessionToken)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error while fetching puzzle")
	}

	scanner := bufio.NewScanner(resp.Body)

	var fileInputs []string
	for scanner.Scan() {
		fileInputs = append(fileInputs, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error while reading: %+v", err)
	}

	return fileInputs, nil
}
