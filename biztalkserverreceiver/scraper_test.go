package biztalkserverreceiver

import (
	"context"
	"github.com/Integrio/biztalk-server-go/client"
	"github.com/Integrio/biztalkserverreceiver/biztalkserverreceiver/internal/metadata"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/scraper/scraperhelper"
	"go.uber.org/zap"
	"testing"
)

func TestScraperStart(t *testing.T) {
	testcases := []struct {
		desc     string
		testFunc func(t *testing.T)
	}{
		{
			desc: "bad config",
			testFunc: func(t *testing.T) {
				s := &biztalkservermetricsScraper{
					config: &Config{
						ControllerConfig:     scraperhelper.NewDefaultControllerConfig(),
						MetricsBuilderConfig: metadata.DefaultMetricsBuilderConfig(),
						Endpoint:             "http://localhost/biztalk",
						Auth:                 "basic",
					},
					settings: componenttest.NewNopTelemetrySettings(),
				}

				err := s.Start(context.Background(), componenttest.NewNopHost())
				require.ErrorIs(t, err, errLoggerNotInitialized)
			},
		},
		{
			desc: "no auth configured",
			testFunc: func(t *testing.T) {
				s := &biztalkservermetricsScraper{
					config: &Config{
						ControllerConfig:     scraperhelper.NewDefaultControllerConfig(),
						MetricsBuilderConfig: metadata.DefaultMetricsBuilderConfig(),
						Endpoint:             "http://localhost/biztalk",
					},
					settings: componenttest.NewNopTelemetrySettings(),
					logger:   zap.NewNop(),
				}

				if err := s.Start(context.Background(), componenttest.NewNopHost()); err != nil {
					t.Errorf("azureScraper.start() error = %v", err)
				}

				require.Implements(t, (*client.HttpRequestDoer)(nil), s.client.Client)
				_, ok := s.client.Client.(*client.MiddlewareDoer)
				require.Falsef(t, ok, "HttpRequestDoer should not be of type client.MiddlewareDoer when no auth is configured")
			},
		},
		{
			desc: "basic auth configured",
			testFunc: func(t *testing.T) {
				s := &biztalkservermetricsScraper{
					config: &Config{
						ControllerConfig:     scraperhelper.NewDefaultControllerConfig(),
						MetricsBuilderConfig: metadata.DefaultMetricsBuilderConfig(),
						Endpoint:             "http://localhost/biztalk",
						Auth:                 "basic",
						Username:             "username",
						Password:             "password",
					},
					settings: componenttest.NewNopTelemetrySettings(),
					logger:   zap.NewNop(),
				}

				if err := s.Start(context.Background(), componenttest.NewNopHost()); err != nil {
					t.Errorf("azureScraper.start() error = %v", err)
				}

				require.IsType(t, (*client.MiddlewareDoer)(nil), s.client.Client)
			},
		},
		{
			desc: "ntlm auth configured",
			testFunc: func(t *testing.T) {
				s := &biztalkservermetricsScraper{
					config: &Config{
						ControllerConfig:     scraperhelper.NewDefaultControllerConfig(),
						MetricsBuilderConfig: metadata.DefaultMetricsBuilderConfig(),
						Endpoint:             "http://localhost/biztalk",
						Auth:                 "ntlm",
						Username:             "username",
						Password:             "password",
					},
					settings: componenttest.NewNopTelemetrySettings(),
					logger:   zap.NewNop(),
				}

				if err := s.Start(context.Background(), componenttest.NewNopHost()); err != nil {
					t.Errorf("azureScraper.start() error = %v", err)
				}

				require.IsType(t, (*client.MiddlewareDoer)(nil), s.client.Client)
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.desc, tc.testFunc)
	}
}
