package cmd

import (
	"github.com/alimy/gin-music/api/v1"
	"github.com/alimy/gin-music/cmd"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/unisx/logus"
	"net/http"
	"time"

	_ "github.com/alimy/gin-music/module/serve"
	_ "github.com/alimy/gin-music/pkg/portal"
)

const (
	certFilePathDefault = "cert.pem" // certificate file default path
	keyFilePathDefault  = "key.pem"  // key file used in https server default path
)

var (
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
	e := gin.Default()

	// Install api router
	api.InstallDefault(e)

	// Setup http.Server
	server := &http.Server{
		Handler: e,
		Addr:    "127.0.0.1:8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// Start http.Server
	if enableHttps {
		logus.Info("listen and serve in https://:8080")
		server.ListenAndServeTLS(certFile, keyFile)
	} else {
		logus.Info("listen and serve in http://:8080")
		server.ListenAndServe()
	}
}
