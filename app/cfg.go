package app

//

import (
	"k8s.io/client-go/rest"
)

func getConfig() *rest.Config {
	// var err error
	/// var cfg *rest.Config
	cfg, err := rest.InClusterConfig()
	if err != nil {
			panic(err)
	}
	return cfg
}

func init() {
	//
}