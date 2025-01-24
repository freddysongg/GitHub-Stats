module github.com/freddysongg/GitHub-Stats

replace github.com/freddysongg/GitHub-Stats/internal/auth => ./internal/auth

replace github.com/freddysongg/GitHub-Stats/internal/cache => ./internal/cache

replace github.com/freddysongg/GitHub-Stats/internal/github => ./internal/github

replace github.com/freddysongg/GitHub-Stats/internal/graphql => ./internal/graphql

replace github.com/freddysongg/GitHub-Stats/internal/health => ./internal/health

replace github.com/freddysongg/GitHub-Stats/internal/kafka => ./internal/kafka

replace github.com/freddysongg/GitHub-Stats/internal/websocket => ./internal/websocket

go 1.22.5

toolchain go1.22.11

require (
	github.com/99designs/gqlgen v0.17.63
	github.com/freddysongg/GitHub-Stats/internal/auth v0.0.0-00010101000000-000000000000
	github.com/freddysongg/GitHub-Stats/internal/cache v0.0.0-00010101000000-000000000000
	github.com/freddysongg/GitHub-Stats/internal/github v0.0.0-00010101000000-000000000000
	github.com/freddysongg/GitHub-Stats/internal/graphql v0.0.0-00010101000000-000000000000
	github.com/freddysongg/GitHub-Stats/internal/health v0.0.0-00010101000000-000000000000
	github.com/freddysongg/GitHub-Stats/internal/kafka v0.0.0-00010101000000-000000000000
	github.com/freddysongg/GitHub-Stats/internal/websocket v0.0.0-00010101000000-000000000000
	github.com/sirupsen/logrus v1.9.3
	github.com/vektah/gqlparser/v2 v2.5.21
	golang.org/x/oauth2 v0.25.0
)

require (
	github.com/agnivade/levenshtein v1.2.0 // indirect
	github.com/go-viper/mapstructure/v2 v2.2.1 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/gorilla/websocket v1.5.3 // indirect
	github.com/hashicorp/golang-lru/v2 v2.0.7 // indirect
	github.com/sosodev/duration v1.3.1 // indirect
	golang.org/x/net v0.34.0 // indirect
	golang.org/x/sys v0.29.0 // indirect
)
