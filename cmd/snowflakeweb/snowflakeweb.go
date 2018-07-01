package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"

	"github.com/teddyking/snowflake/api"
	"github.com/teddyking/snowflake/middleware"
	"github.com/teddyking/snowflake/web/handler"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func init() {
	configureLogging()
}

func main() {
	log.Printf("starting snowflakeweb")

	conn, err := grpc.Dial(configureServerAddress(), configureDialOptions()...)
	if err != nil {
		log.Fatalf("could not connect to server: %s", err.Error())
	}

	flakerService := api.NewFlakerClient(conn)

	handler := handler.New(configureStaticDirPath(), flakerService)

	log.Fatal(http.ListenAndServe(configureListenAddress(), handler))
}

func configureLogging() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	if os.Getenv("DEBUG") == "true" {
		log.SetLevel(log.DebugLevel)
	}
}

func configureServerAddress() string {
	serverHost := os.Getenv("SERVERHOST")
	if serverHost == "" {
		serverHost = "0.0.0.0"
	}
	serverPort := os.Getenv("SERVERPORT")
	if serverPort == "" {
		serverPort = "2929"
	}

	serverAddress := fmt.Sprintf("%s:%s", serverHost, serverPort)
	log.WithFields(log.Fields{"serverAddress": serverAddress}).Debug("configured serverAddress")

	return serverAddress
}

func configureDialOptions() []grpc.DialOption {
	dialOpts := []grpc.DialOption{grpc.WithInsecure()}

	tlsCrtPath := os.Getenv("TLSCRTPATH")
	if tlsCrtPath != "" {
		creds, err := credentials.NewClientTLSFromFile(tlsCrtPath, "")
		if err != nil {
			log.Fatalf("error reading TLS creds from '%s': %s", tlsCrtPath, err.Error())
		}

		dialOpts = []grpc.DialOption{grpc.WithTransportCredentials(creds)}
		log.WithFields(log.Fields{"tlsCrtPath": tlsCrtPath}).Debug("configured tls")
	}

	dialOpts = append(dialOpts, grpc.WithUnaryInterceptor(middleware.WithClientLogging))

	return dialOpts
}

func configureListenAddress() string {
	listenPort := os.Getenv("PORT")
	if listenPort == "" {
		listenPort = "2930"
	}

	listenAddress := fmt.Sprintf("0.0.0.0:%s", listenPort)
	log.WithFields(log.Fields{"listenAddress": listenAddress}).Debug("listenAddress configured")

	return listenAddress
}

func configureStaticDirPath() string {
	staticDirPath := os.Getenv("STATICDIR")
	if staticDirPath == "" {
		staticDirPath = filepath.Join("web", "static")
	}

	log.WithFields(log.Fields{"staticDirPath": staticDirPath}).Debug("staticDirPath configured")

	return staticDirPath
}
