package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	query, err := ioutil.ReadFile("test.gql")
	if err != nil {
		panic(err)
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	accessToken := os.Getenv("GITHUB_TOKEN") // 環境変数からアクセストークンを取得
	if accessToken == "" {
		panic("GITHUB_TOKEN is not set")
	}

	var graphqlQuery = map[string]interface{}{
		"query": string(query),
		"variables": map[string]string{
			"username": "yamasei2112",
		},
	}

	body, err := json.Marshal(graphqlQuery)
	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://api.github.com/graphql", bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}

	req.Header.Add("Authorization", "bearer "+accessToken)
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(respBody))
}
