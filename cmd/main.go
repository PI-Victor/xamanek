package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/cloudflavor/xamanek/pkg/server"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)

	ticker := time.Ticker{}
	errChan := make(chan error, 1)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGSTOP)

	go func() {
		for sig := range c {
			ticker.Stop()
			logrus.Infof("\n%s received, stopping...", sig.String())
			os.Exit(0)
		}
	}()
	go server.StartHTTPServer(errChan)

	t := (time.Duration(2) * time.Second)
	duration := time.NewTicker(t)

	for range duration.C {
		logrus.Debugf("Sleeping... %s", time.Now())
	}

}
