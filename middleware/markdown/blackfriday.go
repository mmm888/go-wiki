package markdown

import "github.com/russross/blackfriday"

const (
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

type Blackfriday struct {
	extFlags  blackfriday.Extensions
	htmlFlags blackfriday.HTMLFlags
}

func (b Blackfriday) HTMLify(markdown []byte) (string, error) {
	renderer := blackfriday.NewHTMLRenderer(blackfriday.HTMLRendererParameters{Flags: b.htmlFlags, Title: "", CSS: ""})
	html := blackfriday.Run(markdown, blackfriday.WithExtensions(b.extFlags), blackfriday.WithRenderer(renderer))

	return string(html), nil
}

func NewBlackfriday() *Blackfriday {
	return &Blackfriday{
		extFlags:  extFlags,
		htmlFlags: htmlFlags,
	}
}
