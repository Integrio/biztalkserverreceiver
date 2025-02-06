package biztalkserverreceiver

import (
	"context"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
)

var (
	typeStr = component.MustNewType("biztalkserver")
)

const (
	defaultInterval = 1 * time.Minute
)

func createDefaultConfig() component.Config {
	return &Config{
		Endpoint: "http://localhost/biztalk/",
		Interval: string(defaultInterval),
	}
}

func NewFactory() receiver.Factory {
	return receiver.NewFactory(
		typeStr,
		createDefaultConfig,
		receiver.WithMetrics(createMetricsReceiver, component.StabilityLevelAlpha))
}

func createMetricsReceiver(_ context.Context, params receiver.Settings, conf component.Config, consumer consumer.Metrics) (receiver.Metrics, error) {
	logger := params.Logger
	smrCfg := conf.(*Config)
	simplemetricReceiver := &biztalkservermetricsScraper{
		logger:       logger,
		nextConsumer: consumer,
		config:       smrCfg,
	}
	return simplemetricReceiver, nil
}
