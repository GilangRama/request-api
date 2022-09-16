package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/golang/protobuf/ptypes/timestamp"
)

type Response struct {
	Status       string `json:"status"`
	TotalResults int    `json:"totalResults"`
	Articles     []struct {
		Source []struct {
			Id   int    `json:"id"`
			Name string `json:"name"`
		} `json:"source"`
		Author      string               `json:"author"`
		Title       string               `json:"title"`
		Description string               `json:"description"`
		URL         string               `json:"url"`
		URLToImage  string               `json:"urlToImage"`
		PublishedAt *timestamp.Timestamp `json:"publishedAt"`
		Content     string               `json:"content"`
	} `json:"articles"`
}

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func main() {
	// api_key := "3e2516dbe166464f904604c16c030141"
	url := "https://newsapi.org/v2/top-headlines?country=id&apiKey=3e2516dbe166464f904604c16c030141"
	resp, getErr := http.Get(url)
	if getErr != nil {
		log.Fatal(getErr)
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	var result Response
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Cannot Unmarshal JSON")
	}

	for _, rec := range result.Articles {
		fmt.Println(rec.Title)
	}
	// body, readErr := ioutil.ReadAll(resp.Body)
	// if readErr != nil {
	// 	log.Fatal(readErr)
	// }
	// fmt.Println(string(body))
}
