package models

type RepoStats struct {
	ID          string
	Name        string
	Stars       int
	Forks       int
	OpenIssues  int
	PullRequests int
}
