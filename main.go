package main

import (
	"flag"
	"github.com/avegao/gocondi"
	"github.com/heroku/rollrus"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"context"
	"time"
	"net/http"
)

const (
	version = "1.0.0"
)

var (
	debug                 = flag.Bool("debug", false, "Print debug logs")
	froniusServiceAddress = flag.String("fronius-service-address", "fronius:50000", "Fronius Service address. Default = fronius:50000")
	buildDate             string
	commitHash            string
	container             *gocondi.Container
	parameters            map[string]interface{}
	server                *http.Server
)

func initContainer() {
	flag.Parse()

	parameters = map[string]interface{}{
		"build_date":              buildDate,
		"debug":                   *debug,
		"commit_hash":             commitHash,
		"version":                 version,
		"fronius_service_address": *froniusServiceAddress,
	}

	logger := initLogger()
	gocondi.Initialize(logger)
	container = gocondi.GetContainer()

	for name, value := range parameters {
		container.SetParameter(name, value)
	}
}

func initLogger() *logrus.Logger {
	logLevel := logrus.InfoLevel
	environment := "release"
	log := logrus.New()

	if *debug {
		logLevel = logrus.DebugLevel
		environment = "debug"
	} else {
		hook := rollrus.NewHook(fmt.Sprintf("%v", parameters["rollbar_token"]), environment)
		log.Hooks.Add(hook)
	}

	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetFormatter(&logrus.TextFormatter{})
	log.SetLevel(logLevel)

	return log
}

func initHttpServer() {
	router := initRouter()
	router.Run(":8080")

	go func() {
		if err := http.ListenAndServe(":8080", router); err != nil {
			container.GetLogger().WithError(err).Panicf("Error creating server")
		} else {
			container.GetLogger().Infof("Listening to 0.0.0.0:8080")
		}
	}()
}

func closeHttpServer() {
	if nil != server {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			container.GetLogger().WithError(err).Fatalf("Server shutdown error")
		}

		container.GetLogger().Debugf("HTTP server closed")
	}
}

func handleInterrupt() {
	gracefulStop := make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)
	go func() {
		<-gracefulStop
		powerOff()
	}()
}

func powerOff() {
	container.GetLogger().Infof("Shutting down...")

	closeHttpServer()

	os.Exit(0)
}

func main() {
	initContainer()
	handleInterrupt()

	logger := container.GetLogger()
	logger.Infof("IoT Fronius Push Service v%s started (commit %s, build date %s)", container.GetStringParameter("version"), container.GetStringParameter("commit_hash"), container.GetStringParameter("build_date"))

	initHttpServer()
}
