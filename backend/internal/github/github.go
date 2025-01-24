package github

import (
	"context"
	"fmt"
	"time"

	"github.com/google/go-github/v58/github"
	"golang.org/x/oauth2"
)

type Service struct {
	client *github.Client
	cache  Cache
}

type Cache interface {
	Get(key string) (interface{}, bool)
	Set(key string, value interface{}, ttl time.Duration)
}

func NewService(token string, cache Cache) *Service {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(context.Background(), ts)

	return &Service{
		client: github.NewClient(tc),
		cache:  cache,
	}
}

func (s *Service) GetUser(ctx context.Context, login string) (*User, error) {
	cacheKey := fmt.Sprintf("user:%s", login)
	if cached, found := s.cache.Get(cacheKey); found {
		return cached.(*User), nil
	}

	user, _, err := s.client.Users.Get(ctx, login)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	result := &User{
		Login:     user.GetLogin(),
		Name:      user.GetName(),
		AvatarURL: user.GetAvatarURL(),
		Bio:       user.GetBio(),
		Company:   user.GetCompany(),
		Location:  user.GetLocation(),
		CreatedAt: user.GetCreatedAt().Time,
		Followers: user.GetFollowers(),
		Following: user.GetFollowing(),
	}

	s.cache.Set(cacheKey, result, 5*time.Minute)
	return result, nil
}

type User struct {
	Login       string
	Name        string
	AvatarURL   string
	Bio         string
	Company     string
	Location    string
	CreatedAt   time.Time
	Followers   int
	Following   int
}
