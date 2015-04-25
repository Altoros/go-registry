package go_registry

import (
	"flag"
	bosherr "github.com/cloudfoundry/bosh-agent/errors"
	boshlog "github.com/cloudfoundry/bosh-agent/logger"
	biregistry "github.com/cloudfoundry/bosh-init/registry"

	"fmt"
	"os"
	"os/signal"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "c", "", "Configuration File")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
}

func newLogger() boshlog.Logger {

	logLevelString := os.Getenv("GO_REGISTRY_LOG_LEVEL")
	level := boshlog.LevelNone
	if logLevelString != "" {
		var err error
		level, err = boshlog.Levelify(logLevelString)
		if err != nil {
			err = bosherr.WrapError(err, "Invalid GO_REGISTRY_LOG_LEVEL value")
		}
	}

	return boshlog.NewLogger(level)
}

func main() {
	logger := newLogger()
	tag := "goRegistry"
	if configFile == "" {
		logger.Error(tag, "No config file specified")
	}
	config, err := InitFromFile(configFile)
	if err != nil {
		logger.Error(tag, err.Error())
	}

	registryServerManager := biregistry.NewServerManager(logger)

	_, err = registryServerManager.Start(config.Username, config.Password, config.Host, config.Port)
	if err != nil {
		logger.Error(tag, err.Error())
	}

	// // registryServer
	// go func() {
	// 	broker.Start()
	// }()
	// handleSignals()
	// broker.Stop()
}

func handleSignals() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c
}
