package markdown

type Markdown interface {
	HTMLify([]byte) (string, error)
}
