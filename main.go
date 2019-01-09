package main

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

var (
	githubAccessToken = os.Getenv("GITHUB_ACCESS_TOKEN")
)

func main() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: githubAccessToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	opt := &github.RepositoryListOptions{Visibility: "private", Affiliation: "owner"}
	repos, _, err := client.Repositories.List(ctx, "", opt)
	if err != nil {
		panic(err)
	}
	for _, r := range repos {
		fmt.Printf("Repo: %s\n", r.GetName())
		collaborators, _, err := client.Repositories.ListCollaborators(ctx, r.GetOwner().GetLogin(), r.GetName(), nil)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("Collaborators: %d\n", len(collaborators))
	}
}
