package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

const layout = "2006-01-02"

func main() {

	var grassURL string = loadEnvURL()

	data_date, data_count := getGrass(grassURL)

	for i := 0; i < len(data_count); i++ {
		t := time.Now()
		if t.Format(layout) == data_date[i] {
			fmt.Printf("this is data_date %s \ndata count is %d \n", t.Format(layout), data_count[i])
		}

	}
}

func loadEnvURL() string {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Printf(".envの読み込みに失敗しました: %v", err)
	}

	url := os.Getenv(("GrassURL"))

	return url
}

func getGrass(url string) ([500]string, [500]int) {
	var data_count [500]int
	var data_date [500]string
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("urlの取得がうまくできませんでした")
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	regex1 := regexp.MustCompile(`data-count`)
	regex2 := regexp.MustCompile(`data-date`)
	regex3 := regexp.MustCompile(`ry="2" `)

	result := regex3.Split(string(body), -1)
	for i := 0; i < len(result); i++ {
		arr := strings.Split(result[i], " ")

		var temp_arr []string

		for j := 0; j < len(arr); j++ {
			if regex2.MatchString(arr[j]) {
				temp_arr = strings.Split(arr[j], "\"")
				data_date[i] = temp_arr[1]

			} else if regex1.MatchString(arr[j]) {
				temp_arr = strings.Split(arr[j], "\"")
				num, _ := strconv.Atoi(temp_arr[1])
				data_count[i] = num
			}

		}
	}
	return data_date, data_count
}
