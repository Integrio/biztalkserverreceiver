package biztalkserverreceiver

import (
	"errors"
	"github.com/Integrio/biztalkserverreceiver/biztalkserverreceiver/internal/metadata"
	"go.opentelemetry.io/collector/scraper/scraperhelper"
	"go.uber.org/multierr"
)

var errNoEndpoint = errors.New("no endpoint specified")
var errNoUsername = errors.New("no username specified")
var errNoPassword = errors.New("no password specified")

type Config struct {
	scraperhelper.ControllerConfig `mapstructure:",squash"`
	metadata.MetricsBuilderConfig  `mapstructure:",squash"`
	Endpoint                       string `mapstructure:"endpoint"`
	Auth                           string `mapstructure:"auth"`
	Username                       string `mapstructure:"username"`
	Password                       string `mapstructure:"password"`
}

const (
	ntlmAuth  = "ntlm"
	basicAuth = "basic"
)

// Validate checks if the receiver configuration is valid
func (cfg *Config) Validate() error {
	var err error

	if cfg.Endpoint == "" {
		err = multierr.Append(err, errNoEndpoint)
	}

	switch cfg.Auth {
	case basicAuth, ntlmAuth:
		if cfg.Username == "" {
			err = multierr.Append(err, errNoUsername)
		}
		if cfg.Password == "" {
			err = multierr.Append(err, errNoPassword)
		}
	default:
	}

	return err
}
