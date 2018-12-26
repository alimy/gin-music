package main

import (
	"github.com/spf13/cobra"
	"github.com/unisx/logus"
)

var serveCmd = &cobra.Command{
	Use:   "serv",
	Short: "serve music agent",
	Long:  "music infomation restfull api agent",
	Run:   startServe,
}

func startServe(_cmd *cobra.Command, _args []string) {
	logus.Debug("start music agent")
	// TODO
}
