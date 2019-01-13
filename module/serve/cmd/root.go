package cmd

import (
	"github.com/alimy/gin-music/api/v1"
	"github.com/alimy/gin-music/cmd"
	"github.com/alimy/gin-music/models"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/unisx/logus"
	"net/http"
	"time"

	_ "github.com/alimy/gin-music/module/serve/openapi"
)

const (
	listenAddrDefault   = "127.0.0.1:8013" // default listen address
	certFilePathDefault = "cert.pem"       // certificate file default path
	keyFilePathDefault  = "key.pem"        // key file used in https server default path
)

var (
	address     string
	certFile    string
	keyFile     string
	enableHttps bool
	inDebug     bool
)

func init() {
	serveCmd := &cobra.Command{
		Use:   "serve",
		Short: "start to ginMusic service",
		Long:  "this cmd will start a https server to provide ginMusic service",
		Run:   serveRun,
	}

	// Parse flags for serveCmd
	serveCmd.Flags().StringVarP(&address, "addr", "a", listenAddrDefault, "service listen address")
	serveCmd.Flags().StringVarP(&certFile, "cert", "c", certFilePathDefault, "certificate path used in https connect")
	serveCmd.Flags().StringVarP(&keyFile, "key", "k", keyFilePathDefault, "key path used in https connect")
	serveCmd.Flags().BoolVarP(&enableHttps, "https", "s", false, "whether use https serve connect")
	serveCmd.Flags().BoolVarP(&inDebug, "debug", "d", false, "whether in debug mode")

	// Register serveCmd as sub-command
	cmd.Register(serveCmd)
}

func serveRun(cmd *cobra.Command, args []string) {
	setup()

	// Instance a default gin engine
	e := gin.Default()

	// Install api router
	api.InstallDefault(e)

	// Setup http.Server
	server := &http.Server{
		Handler: e,
		Addr:    address,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// Start http.Server
	if enableHttps {
		logus.Info("start listen and serve", logus.String("address", address))
		server.ListenAndServeTLS(certFile, keyFile)
	} else {
		logus.Info("listen and serve",
			logus.String("address", address),
			logus.Bool("enableHttps", enableHttps))
		server.ListenAndServe()
	}
}

func setup() {
	if !inDebug {
		logus.InProduction()
		gin.SetMode(gin.ReleaseMode)
	}

	// initial models with MemoryProfile
	if err := models.Register(models.MemoryProfile); err != nil {
		panic(err)
	}
}
