package app

import (
	"k8s.io/client-go/dynamic"
)

var cli dynamic.Interface

func getClient() {
	cli, _ = dynamic.NewForConfig(getConfig())
}

func init() {
	getClient() //
}