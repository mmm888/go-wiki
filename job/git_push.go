package job

import (
	"log"

	"github.com/mmm888/go-wiki/middleware/git"
)

type GitPushJob struct {
	Git *git.Git
}

// git int + git add + git commit + git push
func (j GitPushJob) Serve(data []byte) {
	if j.Git.Root == "" {
		return
	}

	log.Print("Git Init")
	if _, err := j.Git.Init(); err != nil {
		log.Print(err)
	}

	log.Print("Git Commit")
	if _, err := j.Git.Commit(); err != nil {
		log.Print(err)
	}

	log.Print("Git Push")
	if _, err := j.Git.Push(); err != nil {
		log.Print(err)
	}
}
