// Issues prints a table of GitHub issues matching the search terms,
// with issues categorized by age
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopl.io/ch4/github"
)

const (
	week = 7 * 24 * time.Hour
	year = 365 * 24 * time.Hour
)

func list(t string, issues []*github.Issue) {
	fmt.Printf("issues %s:\n", t)
	for _, item := range issues {
		fmt.Printf("#%-5d %s %9.9s %.55s\n",
			item.Number, item.CreatedAt.Format("2006-01-02"), item.User.Login, item.Title)
	}
	fmt.Println()
}

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	var ages [3][]*github.Issue
	for _, item := range result.Items {
		d := time.Since(item.CreatedAt)
		switch {
		case d < week:
			ages[0] = append(ages[0], item)
		case d < year:
			ages[1] = append(ages[1], item)
		default:
			ages[2] = append(ages[2], item)
		}
	}
	fmt.Printf("%d issues total:\n\n", result.TotalCount)
	list("less than 1 week old", ages[0])
	list("less than 1 year old", ages[1])
	list("more than 1 year old", ages[2])
}
