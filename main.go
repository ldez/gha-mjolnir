package main

import (
	"context"
	"log"

	"github.com/google/go-github/v27/github"
	"github.com/ldez/ghactions"
)

func main() {
	displayVersion()

	err := ghactions.NewAction(context.Background()).OnPullRequest(action).Run()
	if err != nil {
		log.Fatal(err)
	}
}

func action(client *github.Client, event *github.PullRequestEvent) error {
	action := event.GetAction()

	if action != "closed" || !event.PullRequest.GetMerged() {
		log.Printf("skip: %q merge %v", action, event.PullRequest.GetMerged())
		return nil
	}

	owner, repoName := ghactions.GetRepoInfo()

	return closeRelatedIssues(context.Background(), client, owner, repoName, event.PullRequest, false)
}
