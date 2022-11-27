package getarchives

import (
	"os"
	"encoding/json"
	"strings"
	"go-deploy/app"
	//"fmt"
)

var c []struct {
	Sha string `json:"sha"`
	DownloadUrl string `json:"download_url"`
	Name string `json:"name"`
}

func Zip() {
	for _, rep := range strings.Split(os.Getenv("rep"), ",") {
			p := strings.Split(rep, "*")
			// https://api.github.com/repos/juju-l/go-deploy/commits?sha=dev&per_page=1&path=/config
			commits,_ := req().Get(""+"repos/"+os.Getenv("usr")+"/"+p[0]+"/commits"+"?sha="+p[2]+"&per_page=1&path="+p[1]+"") /**/
			/*e := */json.Unmarshal(commits.Body(), &c)
			sha := string([]rune(c[0].Sha)[:7])
			// https://api.github.com/repos/owner/rep/zipball/p[2](sha)
			// https://api.github.com/repos/juju-l/go-deploy/contents/config?ref=a34126c
			content,_ := req().Get(""+"repos/"+os.Getenv("usr")+"/"+p[0]+"/contents"+p[1]+"?"+"ref="+sha+"") /**/
			/*e := */json.Unmarshal(content.Body(), &c)
			app.Apply(get())
			// check status ...
	}
}

func init() {
	///
	/////
	//
}