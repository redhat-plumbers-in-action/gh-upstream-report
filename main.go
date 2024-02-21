package main

import (
	"fmt"

	"github.com/cli/go-gh/v2/pkg/api"
)

func main() {
	fmt.Println("hi world, this is the gh-upstream-report extension!")
	client, err := api.DefaultRESTClient()
	if err != nil {
		fmt.Println(err)
		return
	}
	response := struct {Login string}{}
	err = client.Get("user", &response)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("running as %s\n", response.Login)
}

// For more examples of using go-gh, see:
// https://github.com/cli/go-gh/blob/trunk/example_gh_test.go

// Next Steps
// - run 'cd gh-upstream-report; gh extension install .; gh upstream-report' to see your new extension in action
// - run 'go build && gh upstream-report' to see changes in your code as you develop
// - run 'gh repo create' to share your extension with others

// For more information on writing extensions:
// https://docs.github.com/github-cli/github-cli/creating-github-cli-extensions
