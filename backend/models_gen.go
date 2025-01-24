// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package backend

type Language struct {
	Name  string  `json:"name"`
	Color *string `json:"color,omitempty"`
}

type Query struct {
}

type Repository struct {
	Name           string      `json:"name"`
	Description    *string     `json:"description,omitempty"`
	URL            string      `json:"url"`
	StargazerCount int         `json:"stargazerCount"`
	ForkCount      int         `json:"forkCount"`
	Languages      []*Language `json:"languages"`
}

type User struct {
	Login        string        `json:"login"`
	Name         *string       `json:"name,omitempty"`
	AvatarURL    string        `json:"avatarUrl"`
	Bio          *string       `json:"bio,omitempty"`
	Company      *string       `json:"company,omitempty"`
	Location     *string       `json:"location,omitempty"`
	WebsiteURL   *string       `json:"websiteUrl,omitempty"`
	Repositories []*Repository `json:"repositories"`
}