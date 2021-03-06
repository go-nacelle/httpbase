package httpbase

import (
	"fmt"
	"time"
)

type Config struct {
	HTTPHost           string `env:"http_host" default:"0.0.0.0"`
	HTTPPort           int    `env:"http_port" default:"5000"`
	HTTPCertFile       string `env:"http_cert_file"`
	HTTPKeyFile        string `env:"http_key_file"`
	RawShutdownTimeout int    `env:"http_shutdown_timeout" default:"5"`

	ShutdownTimeout time.Duration
}

var ErrBadCertConfig = fmt.Errorf("cert file and key file must both be supplied or both be omitted")

func (c *Config) PostLoad() error {
	if (c.HTTPCertFile == "") != (c.HTTPKeyFile == "") {
		return ErrBadCertConfig
	}

	c.ShutdownTimeout = time.Duration(c.RawShutdownTimeout) * time.Second
	return nil
}
