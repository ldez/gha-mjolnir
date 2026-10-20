package main

import (
	"context"
	"log"

	"github.com/google/go-github/v93/github"
	"github.com/ldez/ghactions"
)

func main() {
	displayVersion()

	act, err := ghactions.NewAction(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	err = act.
		OnPullRequest(func(client *github.Client, event *github.PullRequestEvent) error {
			return action(client, event)
		}).
		OnPullRequestTarget(func(client *github.Client, event *github.PullRequestTargetEvent) error {
			return action(client, event)
		}).
		Run()
	if err != nil {
		log.Fatal(err)
	}
}

type pullRequestBasedEvent interface {
	GetAction() string
	GetPullRequest() *github.PullRequest
}

func action(client *github.Client, event pullRequestBasedEvent) error {
	action := event.GetAction()
	pr := event.GetPullRequest()

	if action != "closed" || !pr.GetMerged() {
		log.Printf("skip: %q merge %v", action, pr.GetMerged())
		return nil
	}

	owner, repoName := ghactions.GetRepoInfo()

	return closeRelatedIssues(context.Background(), client, owner, repoName, pr, false)
}
