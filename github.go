package main

import (
	"fmt"
	"time"
	"net/url"
	"strings"
	"net/http"
	"encoding/json"
	"os"
	"log"
)

const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount  int  `json:"totalCount"`
	Items       []*Issue
}

type Issue struct {
	Number      int
	HTMLURL     string `json:"htmlUrl"`
	Title       string
	State       string
	User		*User
	CreatedAt   time.Time `json:"createdAt"`
	Body        string
}

type User struct {
	Login       string
	HTMLURL     string `json:"htmlUrl"`
}

func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	log.Printf("请求" + IssuesURL + "?q=" + q)
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}
	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func main() {
	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}
