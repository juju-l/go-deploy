package app

//

import (
	"k8s.io/client-go/discovery"
)

func getDiscover() *discovery.DiscoveryClient {
	// var err error
	/// var dis *discovery.DiscoveryClient
	dis, err := discovery.NewDiscoveryClientForConfig(getConfig())
	if err != nil {
			panic(err)
	}
	return dis
}

func init() {
	//
}