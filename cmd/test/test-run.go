package main

import (
	"fmt"
	"os"

	"github.com/nachocano/bitbucket-source/pkg/reconciler"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "check",
	Short: "nanana",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("yo")
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	Execute()

	hook, err := reconciler.CallHook(
		"damian_baar",
		"c4EpTqrm73hJRFfSXSuz",
		"digitalrigbitbucketteam",
		"embracing-nix-docker-k8s-helm-knative",
		[]string{
			"repo:push",
		})
	fmt.Println("test %v, %v", hook, err)
}
