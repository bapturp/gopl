package issuectl

import (
	"encoding/json"
	"fmt"
	issuectl "issuectl/model"
	"net/http"
)

func GetIssue(n int, slug string) (*issuectl.Issue, error) {
	client := http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s/issues/%d", issuectl.RepoURL, slug, n), nil)
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
