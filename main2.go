package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Login string `json:"login"`
}

type Repository struct {
	Name string `json:"name"`
}

func main() {
	username := "yamasei2112"

	user, err := getUserInfo(username)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	fmt.Printf("Username: %s\n", user.Login)

	repos, err := getUserRepos(username)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	fmt.Println("Repositories:")
	for _, repo := range repos {
		fmt.Println(repo.Name)
	}
}

func getUserInfo(username string) (*User, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s", username)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var user User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func getUserRepos(username string) ([]Repository, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/repos", username)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var repos []Repository
	if err := json.NewDecoder(resp.Body).Decode(&repos); err != nil {
		return nil, err
	}

	return repos, nil
}
