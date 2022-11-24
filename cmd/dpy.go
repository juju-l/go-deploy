package main

import (
	"context"
	git "github.com/google/go-github/v48/github"
	"fmt"
)

var cli = git.NewClient(nil)

func main() {
	var ctx = context.Background()
			b,_,_ := cli.Repositories.GetBranch(ctx, "juju-l", "go-deploy", "dev", false)
			opt := &git.RepositoryContentGetOptions{ Ref: *b.Commit.SHA } // *b.GetCommit() ->sha
			/*u,_,_ :=*/ cli.Repositories.GetArchiveLink(ctx, "juju-l", "go-deploy", git.Zipball, opt, true /**/)
			// r := resty.New().R().SetOutput(*b.Commit.SHA).Get(u.String()) // downloadZip
			_,d,_,_ := cli.Repositories.GetContents(ctx, "juju-l", "go-deploy", "", opt)
	fmt.Println(d)
}

func init() {
	//
}