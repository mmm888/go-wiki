package markdown

import (
	"context"

	"github.com/google/go-github/github"
)

type GithubMarkdown struct {
	client *github.Client
	option *github.MarkdownOptions
}

func (g *GithubMarkdown) HTMLify(markdown []byte) (string, error) {

	html, _, err := g.client.Markdown(context.Background(), string(markdown), g.option)

	return html, err
}

func NewGithubMarkdown() *GithubMarkdown {
	return &GithubMarkdown{
		client: github.NewClient(nil),
		option: &github.MarkdownOptions{Mode: "gfm", Context: "google/go-github"},
	}
}
