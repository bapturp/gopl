package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/bapturp/gopl/ch04/ex4.10/github"
)

var result *github.IssuesSearchResult

func main() {
	if len(os.Args) < 1 {
		os.Exit(1)
	}

	r, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatalf("no search terms provided: %s", err)
	}
	result = r

	http.HandleFunc("/", IndexHandler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	var issueList = template.Must(template.New("issue-list").Parse(`
	<h1>{{ .TotalCount }} issues</h1>
	<table>
	<tr style='text-align: left'>
	<th>#</th>
	<th>State</th>
	<th>User</th>
	<th>Title</th>
	<th>Milestone</th>
	</tr>
	{{range .Items}}
	<tr>
	<td><a href='{{ .HTMLURL }}'>{{ .Number }}</td>
	<td>{{.State}}</td>
	<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
	<td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
	<td><a href='{{.Milestone.HTMLURL}}'>{{.Milestone.Title}}</a></td>
	</tr>
	{{end}}
	</table>
	`))

	issueList.Execute(w, result)
}
