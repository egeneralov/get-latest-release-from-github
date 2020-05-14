package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/google/go-github/v31/github"
	"strings"
)

var (
	client  = github.NewClient(nil)
	ctx     = context.Background()
	org     string
	project string

	removePrefix bool
	withShortSha bool

	tagName       string
	err           error
	latestRelease *github.RepositoryRelease
)

func main() {
	flag.StringVar(&org, "o", "egeneralov", "github user/org name")
	flag.StringVar(&project, "p", "get-latest-release-from-github", "repository name")
	flag.BoolVar(&removePrefix, "remove-prefix", false, "remove v prefix from tagname")
	flag.BoolVar(&withShortSha, "with-short-sha", false, "tag-shortsha")
	flag.Parse()

	latestRelease, _, err = client.Repositories.GetLatestRelease(ctx, org, project)

	if err != nil {
		panic(err)
	}

	tagName = *latestRelease.TagName
	if removePrefix {
		tagName = strings.TrimPrefix(tagName, "v")
	}

	if withShortSha {
		commits, _, err := client.Repositories.ListCommits(
			ctx,
			org,
			project,
			&github.CommitsListOptions{
				SHA: *latestRelease.TagName,
			},
		)
		if err != nil {
			panic(err)
		}

		for _, c := range commits {
			tagName = tagName + "-" + string(*c.SHA)[:8]
			break
		}
	}

	fmt.Println(tagName)
}
