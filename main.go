package main

import (
  "flag"
	"context"
	"fmt"
	"github.com/google/go-github/v31/github"
)

var (
  client = github.NewClient(nil)
	ctx = context.Background()

	org = "elastic"
	project = "elasticsearch"
)

func main() {
  // func StringVar(p *string, name string, value string, usage string)
  flag.StringVar(&org, "o", "elastic", "github user/org name")
  flag.StringVar(&project, "p", "elasticsearch", "repository name")
	flag.Parse()
	
	latestRelease, _, err := client.Repositories.GetLatestRelease(ctx, org, project)
  
	if err != nil {
		panic(err)
	}

	fmt.Println(*latestRelease.TagName)
}
