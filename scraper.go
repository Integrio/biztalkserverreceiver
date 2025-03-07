package biztalkserverreceiver

import (
	"context"
	"errors"
	"fmt"
	"github.com/Integrio/biztalkserverreceiver/internal/client"
	"github.com/Integrio/biztalkserverreceiver/internal/metadata"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/scraper/scrapererror"
	"go.uber.org/zap"
	"strings"
	"time"
)

type biztalkservermetricsScraper struct {
	host     component.Host
	cancel   context.CancelFunc
	logger   *zap.Logger
	config   *Config
	client   *client.Client
	mb       *metadata.MetricsBuilder
	settings component.TelemetrySettings
}

type instanceInfo struct {
	serviceType string
	application string
}

var errLoggerNotInitialized = errors.New("logger not initialized")

func newScraper(logger *zap.Logger, config *Config, settings receiver.Settings) *biztalkservermetricsScraper {
	return &biztalkservermetricsScraper{
		logger:   logger,
		config:   config,
		mb:       metadata.NewMetricsBuilder(config.MetricsBuilderConfig, settings),
		settings: settings.TelemetrySettings,
	}
}

func (smr *biztalkservermetricsScraper) Start(ctx context.Context, host component.Host) (err error) {
	smr.host = host
	ctx, smr.cancel = context.WithCancel(ctx)

	if smr.logger == nil {
		return errLoggerNotInitialized
	}

	clientBuilder := client.NewClientBuilder(smr.config.Endpoint)
	switch smr.config.Auth {
	case ntlmAuth:
		smr.logger.Debug("Continuing with NTLM authentication")
		clientBuilder.UseNtlmAuth(smr.config.Username, smr.config.Password)
		break
	case basicAuth:
		smr.logger.Debug("Continuing with basic authentication")
		clientBuilder.UseBasicAuth(smr.config.Username, smr.config.Password)
		break
	default:
		smr.logger.Debug("Continuing without authentication")
	}

	if smr.client, err = clientBuilder.Build(); err != nil {
		return err
	}

	return nil
}

func (smr *biztalkservermetricsScraper) Shutdown(ctx context.Context) error {
	if smr.cancel != nil {
		smr.cancel()
	}
	return nil
}

func (smr *biztalkservermetricsScraper) scrape(ctx context.Context) (pmetric.Metrics, error) {
	if smr.client == nil {
		return pmetric.NewMetrics(), errors.New("client not initialized")
	}

	scrapeErrors := scrapererror.ScrapeErrors{}

	smr.scrapeOrchestrations(ctx, &scrapeErrors)
	smr.scrapeReceiveLocations(ctx, &scrapeErrors)
	smr.scrapeSendPorts(ctx, &scrapeErrors)
	smr.scrapeSendPortGroups(ctx, &scrapeErrors)
	smr.scrapeHostInstances(ctx, &scrapeErrors)

	instanceInfoMap := map[string]*instanceInfo{}
	instanceInfoMap = smr.scrapeSuspendedInstances(ctx, &scrapeErrors)
	smr.scrapeSuspendedMessages(ctx, &scrapeErrors, instanceInfoMap)

	rb := smr.mb.NewResourceBuilder()
	rb.SetBiztalkName(smr.config.Endpoint)
	return smr.mb.Emit(metadata.WithResource(rb.Emit())), scrapeErrors.Combine()
}

func (smr *biztalkservermetricsScraper) scrapeOrchestrations(ctx context.Context, errors *scrapererror.ScrapeErrors) {
	var orchestrationStatusMap = map[string]int64{
		"unenlisted": 0,
		"started":    1,
		"enlisted":   2,
	}

	orchestrations, err := smr.client.GetOrchestrations(ctx)
	if err != nil {
		errors.Add(err)
		return
	}

	for _, orch := range orchestrations {
		statusStr := strings.ToLower(orch.Status)

		value, valueOk := orchestrationStatusMap[statusStr]
		status, statusOk := metadata.MapAttributeOrchestrationStatus[statusStr]
		if !valueOk || !statusOk {
			err := fmt.Errorf("orchestration %s has invalid status %s", orch.FullName, orch.Status)
			errors.AddPartial(1, err)
			continue
		}

		now := pcommon.NewTimestampFromTime(time.Now())
		smr.mb.RecordBiztalkOrchestrationsStatusDataPoint(now, value, status,
			orch.FullName, orch.Description, orch.Host, orch.ApplicationName)
	}
}

func (smr *biztalkservermetricsScraper) scrapeReceiveLocations(ctx context.Context, errors *scrapererror.ScrapeErrors) {
	receiveLocations, err := smr.client.GetReceiveLocations(ctx)
	if err != nil {
		errors.Add(err)
		return
	}

	for _, rl := range receiveLocations {
		now := pcommon.NewTimestampFromTime(time.Now())
		value := int64(boolToInt(rl.Enable))
		smr.mb.RecordBiztalkReceiveLocationsEnabledDataPoint(now, value, rl.Enable, rl.Name, rl.Description)
	}
}

func (smr *biztalkservermetricsScraper) scrapeSendPorts(ctx context.Context, errors *scrapererror.ScrapeErrors) {
	var sendPortStatusMap = map[string]int64{
		"bound":   0,
		"started": 1,
		"stopped": 2,
	}

	sendPorts, err := smr.client.GetSendPorts(ctx)
	if err != nil {
		errors.Add(err)
		return
	}

	for _, sp := range sendPorts {
		statusStr := strings.ToLower(sp.Status)

		value, valueOk := sendPortStatusMap[statusStr]
		status, statusOk := metadata.MapAttributeSendPortStatus[statusStr]
		if !valueOk || !statusOk {
			err := fmt.Errorf("send port %s has invalid status %s", sp.Name, sp.Status)
			errors.AddPartial(1, err)
			continue
		}

		now := pcommon.NewTimestampFromTime(time.Now())
		smr.mb.RecordBiztalkSendPortsStatusDataPoint(now, value, status, sp.Name, sp.Description, sp.ApplicationName)
	}
}

func (smr *biztalkservermetricsScraper) scrapeSendPortGroups(ctx context.Context, errors *scrapererror.ScrapeErrors) {
	var sendPortGroupStatusMap = map[string]int64{
		"bound":   0,
		"started": 1,
		"stopped": 2,
	}

	spGroups, err := smr.client.GetSendPortGroups(ctx)
	if err != nil {
		errors.Add(err)
		return
	}

	for _, spg := range spGroups {
		statusStr := strings.ToLower(spg.Status)

		value, valueOk := sendPortGroupStatusMap[statusStr]
		status, statusOk := metadata.MapAttributeSendPortGroupStatus[statusStr]
		if !valueOk || !statusOk {
			err := fmt.Errorf("send port group %s has invalid status %s", spg.Name, spg.Status)
			errors.AddPartial(1, err)
			continue
		}

		now := pcommon.NewTimestampFromTime(time.Now())
		smr.mb.RecordBiztalkSendportGroupsStatusDataPoint(now, value, status, spg.Name, spg.Description, spg.ApplicationName)
	}
}

func (smr *biztalkservermetricsScraper) scrapeHostInstances(ctx context.Context, errors *scrapererror.ScrapeErrors) {
	var hostInstanceStatusMap = map[string]int64{
		"stopped": 0,
		"running": 1,
		"unknown": 2,
	}

	hostInstances, err := smr.client.GetHostInstances(ctx)
	if err != nil {
		errors.Add(err)
	}

	for _, hi := range hostInstances {
		statusStr := strings.ToLower(hi.ServiceState)

		value, valueOk := hostInstanceStatusMap[statusStr]
		status, statusOk := metadata.MapAttributeHostInstanceStatus[statusStr]
		if !valueOk || !statusOk {
			err := fmt.Errorf("host instance %s has invalid service state %s", hi.Name, hi.ServiceState)
			errors.AddPartial(1, err)
			continue
		}

		now := pcommon.NewTimestampFromTime(time.Now())
		smr.mb.RecordBiztalkHostInstancesStatusDataPoint(now, value, status, hi.Name, hi.HostName)
	}
}

func (smr *biztalkservermetricsScraper) scrapeSuspendedInstances(ctx context.Context, errors *scrapererror.ScrapeErrors) map[string]*instanceInfo {
	instances, err := smr.client.GetOperationalDataInstances(ctx)
	if err != nil {
		errors.Add(err)
		return make(map[string]*instanceInfo)
	}

	instanceInfoMap := make(map[string]*instanceInfo)
	for _, inst := range instances {
		if inst.InstanceStatus != "Suspended" && inst.InstanceStatus != "SuspendedNotResumable" {
			continue
		}

		now := pcommon.NewTimestampFromTime(time.Now())
		smr.mb.RecordBiztalkSuspendedInstancesDataPoint(now, 1, inst.Application, inst.ServiceType, inst.HostName, inst.Class)

		if _, ok := instanceInfoMap[inst.ServiceTypeID.String()]; !ok {
			instanceInfoMap[inst.ServiceTypeID.String()] = &instanceInfo{
				serviceType: inst.ServiceType,
				application: inst.Application,
			}
		}
	}

	return instanceInfoMap
}

func (smr *biztalkservermetricsScraper) scrapeSuspendedMessages(ctx context.Context, errors *scrapererror.ScrapeErrors, instanceInfoMap map[string]*instanceInfo) {
	messages, err := smr.client.GetOperationalDataMessages(ctx)
	if err != nil {
		errors.Add(err)
		return
	}

	for _, msg := range messages {
		if msg.Status != "Suspended" && msg.Status != "SuspendedNotResumable" {
			continue
		}

		now := pcommon.NewTimestampFromTime(time.Now())

		info, ok := instanceInfoMap[msg.ServiceTypeID.String()]
		if !ok {
			smr.logger.Warn("found suspended message with no correlating suspended instance", zap.String("serviceTypeId", msg.ServiceTypeID.String()))
			smr.mb.RecordBiztalkSuspendedMessagesDataPoint(now, 1, "", "", "")
		} else {
			smr.mb.RecordBiztalkSuspendedMessagesDataPoint(now, 1, info.application, info.serviceType, msg.HostName)
		}
	}
}

func boolToInt(b bool) int {
	var i int
	if b {
		i = 1
	} else {
		i = 0
	}
	return i
}
