// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github"
)

// !+
func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	PrintIssues(result)
	// fmt.Printf("%d issues:\n", result.TotalCount)
	// for _, item := range result.Items {
	// 	fmt.Printf("#%-5d %9.9s %.55s\n",
	// 		item.Number, item.User.Login, item.Title)
	// }
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

// func getDate(item *github.Issue) string {
// 	now := time.Now()
// 	oneDayAgo := now.AddDate(0,0,-1)
// 	oneMonthAgo := now.AddDate(0,-1,0)
// 	oneYearAgo := now.AddDate(-1,0,0)
// 	if item.CreatedAt.Before(oneYearAgo) {
// 		return ""
// 	}
// 	if item.CreatedAt
// }

//!-

/*
//!+textoutput
$ go build gopl.io/ch4/issues
$ ./issues repo:golang/go is:open json decoder
13 issues:
#5680    eaigner encoding/json: set key converter on en/decoder
#6050  gopherbot encoding/json: provide tokenizer
#8658  gopherbot encoding/json: use bufio
#8462  kortschak encoding/json: UnmarshalText confuses json.Unmarshal
#5901        rsc encoding/json: allow override type marshaling
#9812  klauspost encoding/json: string tag not symmetric
#7872  extempora encoding/json: Encoder internally buffers full output
#9650    cespare encoding/json: Decoding gives errPhase when unmarshalin
#6716  gopherbot encoding/json: include field name in unmarshal error me
#6901  lukescott encoding/json, encoding/xml: option to treat unknown fi
#6384    joeshaw encoding/json: encode precise floating point integers u
#6647    btracey x/tools/cmd/godoc: display type kind of each named type
#4237  gjemiller encoding/base64: URLEncoding padding is optional
//!-textoutput
*/
