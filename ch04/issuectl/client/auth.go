package issuectl

import "os"

func Auth() string {
	return os.Getenv("GH_TOKEN")
}
