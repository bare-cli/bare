package utils

import (
	"context"
	"fmt"
	"log"
	"os"

	git "github.com/go-git/go-git/v5"
	"github.com/google/go-github/v39/github"
	"golang.org/x/oauth2"
)

var (
	name        = "gotest1"
	description = "This is a test repo created from golang"
	private     = true
)

func CreateRepo() {
	token := os.Getenv("GITHUB_AUTH_TOKEN")
	if token == "" {
		log.Fatal("Unauthorized : No token found")
	}

	if name == "" {
		log.Fatal("No name: New repos must be given a name")
	}

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	r := &github.Repository{Name: &name, Private: &private, Description: &description}
	repo, _, err := client.Repositories.Create(ctx, "", r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Successfully created new repo: %v\n", repo.GetName())
}

type CloneOptions git.CloneOptions

func CloneRepo(dir string, opts CloneOptions) error {
	o := git.CloneOptions(opts)

	_, err := git.PlainClone(dir, false, &o)

	return err
}
