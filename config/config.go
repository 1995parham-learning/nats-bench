package config

import (
	"fmt"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/providers/posflag"
	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

type Config struct {
	URL          string `koanf:"url"`
	UseTLS       bool   `koanf:"use-tls"`
	Instances    int    `koanf:"instances"`
	MessageCount int    `koanf:"message-count"`
	MessageSize  int    `koanf:"message-size"`
	Subject      string `koanf:"subject"`
}

func Register(fs *pflag.FlagSet) {
	fs.StringP(
		"url", "u",
		nats.DefaultURL,
		fmt.Sprintf("nats server url(s) e.g. %s", nats.DefaultURL),
	)

	fs.IntP(
		"instances", "i",
		1,
		"number of running instances",
	)

	fs.IntP(
		"message-count", "m",
		1,
		"number of messages",
	)

	fs.IntP(
		"message-size", "b",
		1,
		"size of messages (bytes)",
	)

	fs.StringP(
		"subject", "t",
		"test-topic",
		"subject",
	)

	fs.BoolP(
		"use-tls", "s",
		false,
		"Use TLS Secure Connection",
	)
}

func New(fs *pflag.FlagSet) Config {
	var instance Config

	k := koanf.New(".")
	if err := k.Load(posflag.Provider(fs, ".", k), nil); err != nil {
		logrus.Errorf("error loading config.yml: %s", err)
	}

	if err := k.Unmarshal("", &instance); err != nil {
		logrus.Fatalf("error unmarshalling config: %s", err)
	}

	logrus.Infof("following configuration is loaded:\n%+v", instance)

	return instance
}
