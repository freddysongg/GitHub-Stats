package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
	"github-dashboard-backend/models"
)

type GitHubClient struct {
	client *githubv4.Client
	cache  *sync.Map
}

func NewGitHubClient(token string) *GitHubClient {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	httpClient := oauth2.NewClient(context.Background(), src)
	
	return &GitHubClient{
		client: githubv4.NewClient(httpClient),
		cache:  &sync.Map{},
	}
}

func (c *GitHubClient) GetProfile(ctx context.Context) (interface{}, error) {
	cacheKey := "profile"
	if val, ok := c.cache.Load(cacheKey); ok {
		return val, nil
	}

	var query struct {
		Viewer struct {
			Login     githubv4.String
			Name      githubv4.String
			Bio       githubv4.String
			AvatarURL githubv4.String
			Followers struct {
				TotalCount githubv4.Int
			}
			Following struct {
				TotalCount githubv4.Int
			}
		} `graphql:"viewer"`
	}

	err := c.client.Query(ctx, &query, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch profile: %w", err)
	}

	profile := map[string]interface{}{
		"login":      query.Viewer.Login,
		"name":       query.Viewer.Name,
		"bio":        query.Viewer.Bio,
		"avatar_url": query.Viewer.AvatarURL,
		"followers":  query.Viewer.Followers.TotalCount,
		"following":  query.Viewer.Following.TotalCount,
	}

	c.cache.Store(cacheKey, profile)
	return profile, nil
}

func (c *GitHubClient) GetRepositories(ctx context.Context) ([]models.RepoStats, error) {
	cacheKey := "repositories"
	if val, ok := c.cache.Load(cacheKey); ok {
		return val.([]models.RepoStats), nil
	}

	var query struct {
		Viewer struct {
			Repositories struct {
				Nodes []struct {
					ID          githubv4.String
					Name        githubv4.String
					Stargazers  struct {
						TotalCount githubv4.Int
					}
					Forks struct {
						TotalCount githubv4.Int
					}
					Issues struct {
						TotalCount githubv4.Int
					}
					PullRequests struct {
						TotalCount githubv4.Int
					}
				}
			} `graphql:"repositories(first: 100, orderBy: {field: UPDATED_AT, direction: DESC})"`
		}
	}

	err := c.client.Query(ctx, &query, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch repositories: %w", err)
	}

	var repos []models.RepoStats
	for _, node := range query.Viewer.Repositories.Nodes {
		repos = append(repos, models.RepoStats{
			ID:          string(node.ID),
			Name:        string(node.Name),
			Stars:       int(node.Stargazers.TotalCount),
			Forks:       int(node.Forks.TotalCount),
			OpenIssues:  int(node.Issues.TotalCount),
			PullRequests: int(node.PullRequests.TotalCount),
		})
	}

	c.cache.Store(cacheKey, repos)
	return repos, nil
}

func (c *GitHubClient) ClearCache() {
	c.cache.Range(func(key, value interface{}) bool {
		c.cache.Delete(key)
		return true
	})
}
