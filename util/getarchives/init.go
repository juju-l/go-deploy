package getarchives

import (
	"os"
	"encoding/json"
	"strings"
)

var c struct {Commit struct {Sha string}}

func Get() {
	for _, rep := range strings.Split(os.Getenv("rep"), ",") {
			p := strings.Split(rep, "*")
			rsp,_ := req().Get(""+"repos/"+os.Getenv("usr")+"/"+p[0]+"/branches/"+p[2]+"") /**/
			/*e := */json.Unmarshal(rsp.Body(), &c)
			/*s := */string([]rune(c.Commit.Sha)[:7])
			// https://api.github.com/repos/owner/rep/zipball/[ s ] getarchive
			//apply
			// check status ...
	}
}

func init() {
	///
	/////
	//
}