package git

import (
	"io/ioutil"
	"os"
	"testing"
)

const (
	testGitDir = "test"
	testGitURL = "https://github.com/hoge/foo.git"
)

var testGit *Git

func SetUp() error {
	if err := os.RemoveAll(testGitDir); err != nil {
		return err
	}

	testGit = NewGit(testGitDir, testGitURL)
	if _, err := testGit.Init(); err != nil {
		return err
	}

	return nil
}

func End() error {
	err := os.RemoveAll(testGitDir)
	return err
}

func TestMain(m *testing.M) {
	var err error
	err = SetUp()
	if err != nil {
		panic(err)
	}

	run := func() int {
		defer End()
		return m.Run()
	}

	if s := run(); s != 0 {
		os.Exit(s)
	}
}

func TestGitInit(t *testing.T) {
	if _, err := testGit.Init(); err != nil {
		t.Errorf("Error git init %v", err)
	}
}

func TestGitCommit(t *testing.T) {

	testFile := "test/aaa"

	if err := ioutil.WriteFile(testFile, []byte("aaa"), os.ModePerm); err != nil {
		t.Errorf("Error write testFile %v", err)
	}

	if _, err := testGit.Commit(); err != nil {
		t.Errorf("Error git add/commit %v", err)
	}
}

func TestGitDiff(t *testing.T) {
	if _, err := testGit.Diff(testGit.Root); err != nil {
		t.Errorf("Error git diff %v", err)
	}

	//if e, a := expected, actual; !reflect.DeepEqual(e, a) {
	//	t.Error(pretty.Compare(e, a))
	//}
}

func TestGitPush(t *testing.T) {
	if _, err := testGit.Push(); err != nil {
		t.Errorf("Error git push %v", err)
	}
}

func TestGitReset(t *testing.T) {
	if _, err := testGit.Reset(); err != nil {
		t.Errorf("Error git reset %v", err)
	}
}
