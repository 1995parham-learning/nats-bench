package common

import (
	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
)

func Connect(url string) *nats.EncodedConn {
	nc, err := nats.Connect(url)
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Infof("Connected to %s from %v", nc.ConnectedAddr(), nc.DiscoveredServers())

	c, err := nats.NewEncodedConn(nc, nats.GOB_ENCODER)
	if err != nil {
		logrus.Fatal(err)
	}

	return c
}
