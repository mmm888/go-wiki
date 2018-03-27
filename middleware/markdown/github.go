package markdown

type GithubMarkdown struct {
}

func (b GithubMarkdown) HTMLify(markdown []byte) (string, error) {

	// https://github.com/google/go-github

	return "", nil
}

func NewGithubMarkdown() *GithubMarkdown {
	return nil
}
