package domain

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/mmm888/go-wiki/middleware/variable"
	blackfriday "gopkg.in/russross/blackfriday.v2"
)

const (
	defaultFileName = "README.md"

	// blackfriday
	// Extensions
	extFlags = blackfriday.Footnotes | // Pandoc-style footnotes
		blackfriday.NoEmptyLineBeforeBlock | // No need to insert an empty line to start a (code, quote, ordered list, unordered list) block
		blackfriday.HeadingIDs | // specify heading IDs  with {#id}
		blackfriday.Titleblock | // Titleblock ala pandoc
		blackfriday.AutoHeadingIDs | // Create the heading ID from the text
		blackfriday.BackslashLineBreak | // Translate trailing backslashes into line breaks
		blackfriday.DefinitionLists

	// HTMLFlags and Renderer
	htmlFlags = blackfriday.HTMLFlagsNone |
		blackfriday.FootnoteReturnLinks | //Generate a link at the end of a footnote to return to the source
		blackfriday.Smartypants | //Enable smart punctuation substitutions
		blackfriday.SmartypantsFractions | // Enable smart fractions (with Smartypants)
		blackfriday.SmartypantsDashes | // Enable smart dashes (with Smartypants)
		blackfriday.SmartypantsLatexDashes | // Enable LaTeX-style dashes (with Smartypants)
		blackfriday.SmartypantsAngledQuotes | // Enable angled double quotes (with Smartypants) for double quotes rendering
		blackfriday.SmartypantsQuotesNBSP // Enable French guillemets Â» (with Smartypants)
)

type ShowUseCase struct {
}

type ShowInput struct {
	Path       string
	CommonVars *variable.CommonVars
}

type ShowOutput struct {
	Path     string
	Query    string
	Tree     string
	Contents string
}

// ファイルを読み込み、Markdwon to HTML の結果を返す
func (s *ShowUseCase) Get(in *ShowInput) (*ShowOutput, error) {
	root := in.CommonVars.Name
	fpath := filepath.Join(root, in.Path)

	path, err := checkDirTrav(root, fpath)
	if err != nil {
		return nil, err
	}

	tree, err := dirTree(root, root)
	if err != nil {
		return nil, err
	}

	// root ディレクトリ以外の場所 or 存在しない path のチェック
	if path == "" {
		return &ShowOutput{Path: in.Path, Tree: tree}, nil
	}

	// ディレクトリの場合は defaultFile にアクセス
	if fi, _ := os.Lstat(path); fi.IsDir() {
		path = filepath.Join(path, defaultFileName)

		if _, err := os.Stat(path); err != nil {
			return &ShowOutput{Path: in.Path, Tree: tree}, err
		}
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	renderer := blackfriday.NewHTMLRenderer(blackfriday.HTMLRendererParameters{Flags: 0, Title: "", CSS: ""})
	out := blackfriday.Run(data, blackfriday.WithExtensions(extFlags), blackfriday.WithRenderer(renderer))

	return &ShowOutput{Path: in.Path, Tree: tree, Contents: string(out)}, nil
}
