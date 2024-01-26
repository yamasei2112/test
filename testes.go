//main.go

package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	var grassURL string = loadEnvURL()
}

func loadEnvURL() string {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Printf(".envの読み込みに失敗しました: %v", err)
	}

	url := os.Getenv(("GrassURL"))

	return url
}
