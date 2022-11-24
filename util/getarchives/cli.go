package getarchives

import (
	"os"
	"github.com/go-resty/resty/v2"
	//"fmt"
)

var cli *resty.Client

func req() *resty.Request {
	return cli.
			SetBaseURL(os.Getenv("hst")).
			SetHeader("accept", "application/vnd.github+json").
			SetAuthToken(os.Getenv("token")).
	R()
}

func init() {
	cli = resty.New()
}