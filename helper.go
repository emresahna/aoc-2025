package aoc2025

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func LoadFromSource(url string) ([]string, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env file")
	}

	sessionToken := os.Getenv("SESSION_ID")
	if sessionToken == "" {
		return nil, fmt.Errorf("SESSION_ID not found in .env file")
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Cookie", "session="+sessionToken)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error while fetching puzzle")
	}

	return readInput(resp.Body)
}

func LoadFromFile(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	return readInput(file)
}

func readInput(input io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(input)

	var fileInputs []string
	for scanner.Scan() {
		fileInputs = append(fileInputs, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error while reading: %+v", err)
	}

	return fileInputs, nil
}
