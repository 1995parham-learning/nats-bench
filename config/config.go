package config

type Config struct {
	URL          string
	UseTLS       bool
	Instances    int
	MessageCount int
	MessageSize  int
}
