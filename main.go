package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	query, err := ioutil.ReadFile("test.gql")
	if err != nil {
		panic(err)
	}

	var graphqlQuery = map[string]interface{}{
		"query": string(query),
		"variables": map[string]string{
			"username": "yamasei2112", // GitHubのユーザー名を設定
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

	req.Header.Add("Authorization", "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")
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
