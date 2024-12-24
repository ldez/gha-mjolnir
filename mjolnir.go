package main

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/google/go-github/v62/github"
)

var issueRegex = `(?:https://github\.com/[\w-]+/[\w-]+/issues/\d+|#\d+)`

func buildGlobalFixesIssueRegex() *regexp.Regexp {
	actions := strings.Join([]string{"close", "closes", "closed", "fix", "fixes", "fixed", "resolve", "resolves", "resolved"}, "|")
	endPattern := `(?:[\n\r\s,\.]|$)`
	regexString := `(?i)(?:` + actions + `)[\s:]+` + issueRegex + `(?:[, ]+` + issueRegex + `)*` + endPattern
	return regexp.MustCompile(regexString)
}

var (
	issueRE            = regexp.MustCompile(issueRegex)
	globalFixesIssueRE = buildGlobalFixesIssueRegex()
)

// closeRelatedIssues Closes issues listed in the PR description.
func closeRelatedIssues(ctx context.Context, client *github.Client, owner string, repositoryName string, pr *github.PullRequest, dryRun bool) error {
	issueNumbers := parseIssueFixes(pr.GetBody(), owner, repositoryName)

	repo, _, err := client.Repositories.Get(ctx, owner, repositoryName)
	if err != nil {
		return fmt.Errorf("unable to access repository %s/%s: %w", owner, repositoryName, err)
	}

	for _, issueNumber := range issueNumbers {
		log.Printf("PR #%d: closes issue #%d, add milestones %s", pr.GetNumber(), issueNumber, pr.Milestone.GetTitle())
		if !dryRun {
			err := closeIssue(ctx, client, owner, repositoryName, pr, issueNumber)
			if err != nil {
				return fmt.Errorf("unable to close issue #%d: %w", issueNumber, err)
			}
		}

		// Add comment if needed
		if pr.Base.GetRef() != repo.GetDefaultBranch() {
			message := fmt.Sprintf("Closed by #%d.", pr.GetNumber())

			log.Printf("PR #%d: issue #%d, add comment: %s", pr.GetNumber(), issueNumber, message)

			if !dryRun {
				err := addComment(ctx, client, owner, repositoryName, issueNumber, message)
				if err != nil {
					return fmt.Errorf("unable to add comment on issue #%d: %w", issueNumber, err)
				}
			}
		}
	}

	return nil
}

func closeIssue(ctx context.Context, client *github.Client, owner string, repositoryName string, pr *github.PullRequest, issueNumber int) error {
	var milestone *int
	if pr.Milestone != nil {
		milestone = pr.Milestone.Number
	}

	issueRequest := &github.IssueRequest{
		Milestone: milestone,
		State:     github.String("closed"),
	}

	_, _, err := client.Issues.Edit(ctx, owner, repositoryName, issueNumber, issueRequest)
	return err
}

func addComment(ctx context.Context, client *github.Client, owner string, repositoryName string, issueNumber int, message string) error {
	issueComment := &github.IssueComment{
		Body: github.String(message),
	}
	_, _, err := client.Issues.CreateComment(ctx, owner, repositoryName, issueNumber, issueComment)
	return err
}

func parseIssueFixes(text string, owner string, repoName string) []int {
	var issueNumbers []int

	fixMatches := globalFixesIssueRE.FindAllStringSubmatch(text, -1)

	for _, fixMatch := range fixMatches {
		fmt.Println(fixMatch)

		issueMatches := issueRE.FindAllString(fixMatch[0], -1)
		for _, issue := range issueMatches {
			if strings.HasPrefix(issue, "#") {
				numb, err := strconv.ParseInt(strings.TrimPrefix(issue, "#"), 10, 64)
				if err != nil {
					log.Println(err)
				} else {
					issueNumbers = append(issueNumbers, int(numb))
				}
			} else if strings.HasPrefix(issue, "https://github.com/") {
				urlParts := strings.Split(issue, "/")
				if len(urlParts) >= 5 && urlParts[len(urlParts)-2] == "issues" {
					issueOwner := urlParts[len(urlParts)-4]
					issueRepo := urlParts[len(urlParts)-3]
					if issueOwner == owner && issueRepo == repoName {
						issueNumber, err := strconv.Atoi(urlParts[len(urlParts)-1])
						if err == nil {
							issueNumbers = append(issueNumbers, issueNumber)
						} else {
							log.Println(err)
						}
					} else {
						log.Println("Skipping issue from different repository:", issue)
					}
				}
			}
		}
	}
	return issueNumbers
}
