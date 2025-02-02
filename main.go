package main

import (
	"bufio"
	"fmt"
	"go-lang/blinkchat/router"
	db "go-lang/blinkchat/utils"
	"log"
	"net/http"
	"os"
	"strings"
)

func loadEnvFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Ignore comments and blank lines
		if strings.HasPrefix(line, "#") || strings.TrimSpace(line) == "" {
			continue
		}

		// Split key and value
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue // Skip malformed lines
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		// Set the environment variable
		os.Setenv(key, value)
	}

	return scanner.Err()
}

func main() {
	err := loadEnvFile(".env")
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		return
	}

	db.ConnectToMongo(os.Getenv("DB_CONNECTION_STRING"))
	log.Fatal(http.ListenAndServe(":3000", router.Router()))
}
