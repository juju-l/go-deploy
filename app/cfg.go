package app

//

import (
	// "net/http"
	"k8s.io/client-go/rest"
	// "net/url"
)

func getConfig() *rest.Config {
	// var err error
	cfg, err := rest.InClusterConfig()
	if err != nil {
			panic(err)
	}
	// cfg.Proxy = func(*http.Request) (*url.URL, error) { return url.Parse(os.Getenv("proxy")) }
	return cfg
}

func init() {
	//
}