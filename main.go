package main

import (
	"flag"
	log "github.com/Sirupsen/logrus"
	"github.com/xixitalk/parkomat/config"
	"github.com/xixitalk/parkomat/web"
	"github.com/xixitalk/parkomat/webdav"
	"os"
	"sync"
)

func main() {
	log.WithFields(log.Fields{
		"service": "main",
	}).Info("Parkomat (parkomat.io)")

	configFile := flag.String("config_file", "parkomat.toml", "Configuration File")
	flag.Parse()

	var c *config.Config
	var err error

	// If you specify environment variable, args will be overwritten
	envConfigFile := os.Getenv("PARKOMAT_CONFIG_FILE")
	if envConfigFile != "" {
		configFile = &envConfigFile
	}

	c, err = config.NewConfigFromFile(*configFile)
	if err != nil {
		log.WithFields(log.Fields{
			"service": "main",
			"path":    *configFile,
			"error":   err,
		}).Error("Can't read config file")
		return
	}

	var wg sync.WaitGroup

	// TODO: implement graceful shutdown

	s := web.NewServer(c)
	dav := webdav.NewWebDav(c)

	s.Init()

	err = dav.Init()
	if err == nil {
		s.AddHandlerFunc(c.WebDav.Mount, dav.HandlerFunc)
	}

	wg.Add(1)
	go func() {
		err = s.Serve()
		if err != nil {
			log.WithFields(log.Fields{
				"service": "main",
				"error":   err,
			}).Error("Web Error")
		}
		wg.Done()
	}()

	wg.Wait()
	log.WithFields(log.Fields{
		"service": "main",
	}).Info("Exit")
}
