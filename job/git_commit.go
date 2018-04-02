package job

import (
	"log"

	"github.com/mmm888/go-wiki/middleware/git"
)

type GitCommitJob struct {
	Git *git.Git
}

func (j GitCommitJob) Serve(data []byte) {
	if j.Git.Root == "" {
		return
	}

	if _, err := j.Git.Init(); err != nil {
		log.Print(err)
	}

	if _, err := j.Git.Commit(); err != nil {
		log.Print(err)
	}
}
