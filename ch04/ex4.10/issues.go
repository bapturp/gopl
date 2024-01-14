// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/bapturp/gopl/ch04/ex4.10/github"
)

// !+
func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	PrintIssues(result)
}

func PrintIssues(result *github.IssuesSearchResult) {
	fmt.Printf("%d issues:\n", result.TotalCount)

	now := time.Now()
	oneMonthAgo := now.AddDate(0, -1, 0)
	oneYearAgo := now.AddDate(-1, 0, 0)

	fmt.Printf("\nIssue(s) less than a month old.\n")
	for _, item := range result.Items {
		if item.CreatedAt.After(oneMonthAgo) && item.CreatedAt.Before(oneYearAgo) {
			fmt.Printf("#%-5d %9.9s %.55s\n",
				item.Number, item.User.Login, item.Title)
		}
	}

	fmt.Printf("\nIssue(s) less than a year old.\n")
	for _, item := range result.Items {
		if item.CreatedAt.After(oneYearAgo) && item.CreatedAt.Before(oneMonthAgo) {
			fmt.Printf("#%-5d %9.9s %.55s\n",
				item.Number, item.User.Login, item.Title)
		}
	}

	fmt.Printf("\nIssue(s) more than a year old.\n")
	for _, item := range result.Items {
		if item.CreatedAt.After(oneYearAgo) {
			fmt.Printf("#%-5d %9.9s %.55s\n",
				item.Number, item.User.Login, item.Title)
		}
	}
}
