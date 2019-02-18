package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/google/go-github/v24/github"
	"golang.org/x/oauth2"
)

const (
	// https://developer.github.com/v3/activity/events/types/#pullrequestevent
	pullRequestEventName = "pull_request"
)

func action(dryRun bool) error {
	gitHubToken := os.Getenv("GITHUB_TOKEN")
	eventName := os.Getenv("GITHUB_EVENT_NAME")
	eventPath := os.Getenv("GITHUB_EVENT_PATH")

	if len(gitHubToken) == 0 {
		return errors.New("GITHUB_TOKEN is required")
	}

	if eventName != pullRequestEventName {
		return fmt.Errorf("invalid event type: %q", eventName)
	}

	event := &github.PullRequestEvent{}
	err := readEvent(eventPath, event)
	if err != nil {
		return err
	}

	action := event.GetAction()

	if action != "closed" || !event.PullRequest.GetMerged() {
		log.Printf("skip: %q merge %v", action, event.PullRequest.GetMerged())
		return nil
	}

	ctx := context.Background()
	client := newGitHubClient(ctx, gitHubToken)

	owner, repoName := getRepoInfo()

	return closeRelatedIssues(ctx, client, owner, repoName, event.PullRequest, dryRun)
}

func newGitHubClient(ctx context.Context, token string) *github.Client {
	var client *github.Client
	if len(token) == 0 {
		client = github.NewClient(nil)
	} else {
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		)
		tc := oauth2.NewClient(ctx, ts)
		client = github.NewClient(tc)
	}
	return client
}

func readEvent(eventPath string, event interface{}) error {
	content, err := ioutil.ReadFile(eventPath)
	if err != nil {
		return err
	}

	return json.Unmarshal(content, event)
}

func getRepoInfo() (string, string) {
	githubRepository := os.Getenv("GITHUB_REPOSITORY")

	parts := strings.SplitN(githubRepository, "/", 2)

	return parts[0], parts[1]
}
