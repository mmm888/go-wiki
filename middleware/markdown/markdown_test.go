package markdown

import (
	"reflect"
	"testing"

	"github.com/kylelemons/godebug/pretty"
)

func TestBlackfriday(t *testing.T) {

	m := NewBlackfriday()

	input := "# aaa"

	expected, err := m.HTMLify([]byte(input))
	if err != nil {
		t.Errorf("Error blackfriday %v", err)
	}

	actual := "<h1 id=\"aaa\">aaa</h1>\n"

	if e, a := expected, actual; !reflect.DeepEqual(e, a) {
		t.Error(pretty.Compare(e, a))
	}
}

func TestGithubMarkdown(t *testing.T) {

	m := NewGithubMarkdown()

	input := "# aaa"

	expected, err := m.HTMLify([]byte(input))
	if err != nil {
		t.Errorf("Error blackfriday %v", err)
	}

	actual := "<h1>aaa</h1>"

	if e, a := expected, actual; !reflect.DeepEqual(e, a) {
		t.Error(pretty.Compare(e, a))
	}
}
