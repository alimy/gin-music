package cmd

import (
	"context"
	"github.com/alimy/gin-music/cmd"
	"github.com/alimy/gin-music/module/serve"
	_ "github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var AppContext = context.Background()

const (
	certFilePathDefault = "cert.pem" // certificate file default path
	keyFilePathDefault  = "key.pem"  // key file used in https server default path
)

var (
	muxType     string
	certFile    string
	keyFile     string
	enableHttps bool
)

func init() {
	serveCmd := &cobra.Command{
		Use:   "serve",
		Short: "start to okmall service",
		Long:  "this cmd will start a https server to provide okmall service",
		Run:   serveRun,
	}

	// Parse flags for serveCmd
	serveCmd.Flags().StringVarP(&certFile, "cert", "c", certFilePathDefault, "certificate path used in https connect")
	serveCmd.Flags().StringVarP(&keyFile, "key", "k", keyFilePathDefault, "key path used in https connect")
	serveCmd.Flags().BoolVarP(&enableHttps, "https", "s", false, "whether use https serve connect")

	// Register serveCmd as sub-command
	cmd.Register(serveCmd)
}

func serveRun(cmd *cobra.Command, args []string) {
	// TODO
	serve.StartService()
}
