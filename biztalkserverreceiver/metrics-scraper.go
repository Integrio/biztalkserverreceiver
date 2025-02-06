package biztalkserverreceiver

import (
	"context"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/zap"
)

type biztalkservermetricsScraper struct {
	host          component.Host
	cancel        context.CancelFunc
	logger        *zap.Logger
	nextConsumer  consumer.Metrics
	config        *Config
	simpleGauge   int64
	simpleCounter int64
}

func (smr *biztalkservermetricsScraper) scrapeMetrics() {
	smr.simpleCounter += smr.simpleGauge
	smr.logger.Sugar().Infof("Counter %d", smr.simpleCounter)
}

func (smr *biztalkservermetricsScraper) Start(ctx context.Context, host component.Host) error {
	smr.host = host
	ctx = context.Background()
	ctx, smr.cancel = context.WithCancel(ctx)
	interval, _ := time.ParseDuration(smr.config.Interval)
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				smr.logger.Info("Start processing metrcis")
				ts := pcommon.NewTimestampFromTime(time.Now())

				smr.scrapeMetrics()

				metrics := pmetric.NewMetrics()

				m1 := pmetric.NewMetric()
				m1.SetName("simplecounter")
				m1.SetDescription("A Simple Counter")
				m1.SetUnit("unit")
				var dp1 pmetric.NumberDataPoint
				sum := m1.SetEmptySum()
				sum.SetIsMonotonic(true)
				sum.SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
				dp1 = sum.DataPoints().AppendEmpty()
				dp1.SetTimestamp(ts)
				dp1.SetStartTimestamp(ts)
				dp1.SetIntValue(smr.simpleCounter)
				// append to metrics
				newMetric1 := metrics.ResourceMetrics().AppendEmpty().ScopeMetrics().AppendEmpty().Metrics().AppendEmpty()
				m1.MoveTo(newMetric1)

				//consume
				err := smr.nextConsumer.ConsumeMetrics(ctx, metrics)
				if err != nil {
					smr.logger.Error(err.Error())
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	return nil
}

func (smr *biztalkservermetricsScraper) Shutdown(ctx context.Context) error {
	if smr.cancel != nil {
		smr.cancel()
	}
	return nil
}
