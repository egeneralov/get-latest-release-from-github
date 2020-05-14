package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/google/go-github/v31/github"
)

var (
	client  = github.NewClient(nil)
	ctx     = context.Background()
	org     string
	project string

	err           error
	latestRelease *github.RepositoryRelease
)

func main() {
	flag.StringVar(&org, "o", "egeneralov", "github user/org name")
	flag.StringVar(&project, "p", "get-latest-release-from-github", "repository name")
	flag.Parse()

	latestRelease, _, err = client.Repositories.GetLatestRelease(ctx, org, project)

	if err != nil {
		panic(err)
	}

	fmt.Println(*latestRelease.TagName)
}
