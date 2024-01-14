package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"

	client "github.com/bapturp/gopl/ch04/ex4.11/client"
	"github.com/tj/go-editor"
)

func main() {
	args := &Arguments{}
	if len(os.Args) < 2 {
		fmt.Fprint(os.Stderr, "Usage issuectl <get|create|update>")
		os.Exit(1)
	}
	ParseVerb(os.Args[1], args)

	switch args.Verb {
	case "get":
		Get(args)
	case "create":
		Create(args)
	}
}

type Arguments struct {
	Verb        string
	Slug        string
	IssueId     int
	Title       string
	Description string
}

func ParseVerb(verb string, args *Arguments) error {
	rVerb := regexp.MustCompile(`get|create|update|delete`)
	if rVerb.MatchString(verb) {
		args.Verb = verb
		return nil
	} else {
		return fmt.Errorf("verb not valid: %s\n", verb)
	}
}

func ParseSlug(slug string, args *Arguments) error {
	rSlug := regexp.MustCompile(`^[\-a-zA-Z0-9]+\/[\-a-zA-Z0-9]+$`)
	if rSlug.MatchString(slug) {
		args.Slug = slug
		return nil
	} else {
		return fmt.Errorf("slug not valid: %s\n", slug)
	}
}

func ParseIssueId(issueId string, args *Arguments) error {
	issueid, err := strconv.Atoi(os.Args[3])
	if err != nil {
		return fmt.Errorf("issue Id not valid: %s\n", issueId)
	} else {
		args.IssueId = issueid
		return nil
	}
}

func ParseIssueTitle(title string, args *Arguments) error {
	rTitle := regexp.MustCompile(`[a-zA-Z0-9]+`)
	if rTitle.MatchString(title) {
		args.Title = title
		return nil
	} else {
		return fmt.Errorf("title not valid: %s\n", title)
	}
}

func Get(args *Arguments) {
	if len(os.Args) < 3 {
		fmt.Fprint(os.Stderr, "Usage issuectl get <org/repo> <issueId>")
		os.Exit(1)
	}

	ParseSlug(os.Args[2], args)
	ParseIssueId(os.Args[3], args)

	issue, err := client.GetIssue(args.IssueId, args.Slug)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("%.55s %.55s\n", issue.Title, issue.Description)
}

func Create(args *Arguments) {
	if len(os.Args) < 4 {
		fmt.Fprintln(os.Stderr, "Usage issuectl get <org/repo> <issue_title>")
		os.Exit(1)
	}
	ParseSlug(os.Args[2], args)
	ParseIssueTitle(os.Args[3], args)

	body, err := editor.Read()
	if err != nil {
		log.Fatalf("error: %s\n", err)
	}

	newIssue, err := client.CreateIssue(args.Slug, args.Title, string(body))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(newIssue)
}

func GetBody() string {
	b, err := editor.Read()
	if err != nil {
		log.Fatalf("error: %s\n", err)
	}
	return string(b)
}
