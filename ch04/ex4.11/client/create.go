package issuectl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	issuectl "github.com/bapturp/gopl/ch04/ex4.11/model"
)

func CreateIssue(slug, title, body string) (*issuectl.Issue, error) {
	newIssue := struct {
		Title string `json:"title"`
		Body  string `json:"body,omitempty"`
	}{title, body}

	payload, err := json.Marshal(newIssue)
	if err != nil {
		return nil, err
	}

	client := http.Client{}
	fmt.Printf("%s/%s/issues\n", issuectl.RepoURL, slug)
	fmt.Println(string(payload))
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/%s/issues", issuectl.RepoURL, slug), bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	req.Header = http.Header{
		"Accept":               {"application/vnd.github+json"},
		"Authorization":        {"Bearer " + Auth()},
		"X-GitHub-Api-Version": {"2022-11-28"},
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("get issue failed: %s", resp.Status)
	}

	var result issuectl.Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}
