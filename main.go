package main

import (
	"fmt"
	"os"
	"context"
	"bufio"
	"flag"
	"log"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func main() {
	user := flag.String("user", "", "github username")
	permission := flag.String("permission", "push", "One of pull, push or admin. Default: push")
	organisation := flag.String("organisation", "Alfresco", "Organisation. Default: Alfresco")
	file := flag.String("file", "", "Filename containing all the repos")

	flag.Parse()

	if *user == "" {
		log.Fatal("Username cannot be empty")
	}

	if *file == "" {
		log.Fatal("Filename cannot be empty")
	}

	// http context needed for github API
	ctx := context.Background()

	// Log in with API Token and create a client
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_AUTH_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	opt := &github.RepositoryAddCollaboratorOptions{Permission: *permission}

	// Open file containing list of repos
	fileHandle, err := os.Open(*file)
	if err != nil {
		log.Fatal(fmt.Sprintf("Couldn't open file: %s", *file))
	}
	defer fileHandle.Close()

	scanner := bufio.NewScanner(fileHandle)

	for scanner.Scan() {
		name := scanner.Text()

		_, err := client.Repositories.AddCollaborator(ctx, *organisation, name, *user, opt)
		if err != nil {
			fmt.Println(err)
		}
	}
}
