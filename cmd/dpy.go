package main

import (
	"go-tools/util"
	"go-tools/cmd/k8sic"
	"os";"fmt"//
)

var rootCmd = util.NewCommand("kic", "=> ...\njujuGo/kic app - Realize what you think and help you solve what you think ! Is our purpose", "")

func main() {
	//
	///rootCmd := &cobra.Command{}
	//
	rootCmd.AddCommand(
	 k8sic.NewGet(),
	 k8sic.NewCluster(),
	 k8sic.NewDocker(),
	)
	if err := rootCmd.Execute();err != nil { //
	/*;*/fmt.Println(err);os.Exit(-1)
	}
}

func init() {
	//
}