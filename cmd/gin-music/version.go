package main

import (
	"fmt"
	"github.com/alimy/gin-music/version"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use: "version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s (APIVersion:%s)\nBuildTime:%s\nBuildGitSHA:%s\n",
			version.Version, version.APIVersion,
			version.BuildTime, version.BuildGitHash)
	},
}
