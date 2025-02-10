package client

import (
	"github.com/google/uuid"
)

// CertificateInfo Encryption certificate
type CertificateInfo struct {
	CommonName string `json:"CommonName,omitempty"`
	Thumbprint string `json:"Thumbprint,omitempty"`
}

// HostInstance The host instance.
type HostInstance struct {
	HostName      string `json:"HostName,omitempty"`
	HostType      string `json:"HostType,omitempty"`
	IsDisabled    bool   `json:"IsDisabled,omitempty"`
	Logon         string `json:"Logon,omitempty"`
	NTGroupName   string `json:"NTGroupName,omitempty"`
	Name          string `json:"Name,omitempty"`
	RunningServer string `json:"RunningServer,omitempty"`
	ServiceState  string `json:"ServiceState,omitempty"`
}

// Instance The instance.
type Instance struct {
	Adapter              string     `json:"Adapter,omitempty"`
	Application          string     `json:"Application,omitempty"`
	Class                string     `json:"Class,omitempty"`
	CreationTime         string     `json:"CreationTime,omitempty"`
	ErrorCode            string     `json:"ErrorCode,omitempty"`
	ErrorDescription     string     `json:"ErrorDescription,omitempty"`
	HostName             string     `json:"HostName,omitempty"`
	Id                   *uuid.UUID `json:"Id,omitempty"`
	InstanceStatus       string     `json:"InstanceStatus,omitempty"`
	MessageBoxDb         string     `json:"MessageBoxDb,omitempty"`
	MessageBoxServer     string     `json:"MessageBoxServer,omitempty"`
	PendingJobSubmitTime string     `json:"PendingJobSubmitTime,omitempty"`
	PendingOperation     string     `json:"PendingOperation,omitempty"`
	ProcessingServer     string     `json:"ProcessingServer,omitempty"`
	ServiceType          string     `json:"ServiceType,omitempty"`
	ServiceTypeID        *uuid.UUID `json:"ServiceTypeID,omitempty"`
	SuspendTime          string     `json:"SuspendTime,omitempty"`
	Uri                  string     `json:"Uri,omitempty"`
}

// MessageBodyTracking Message body tracking
type MessageBodyTracking struct {
	AfterReceivePipeline  bool `json:"AfterReceivePipeline,omitempty"`
	AfterSendPipeline     bool `json:"AfterSendPipeline,omitempty"`
	BeforeReceivePipeline bool `json:"BeforeReceivePipeline,omitempty"`
	BeforeSendPipeline    bool `json:"BeforeSendPipeline,omitempty"`
}

// MessagePropertyTracking Message Property Tracking
type MessagePropertyTracking struct {
	AfterReceivePipeline  bool `json:"AfterReceivePipeline,omitempty"`
	AfterSendPipeline     bool `json:"AfterSendPipeline,omitempty"`
	BeforeReceivePipeline bool `json:"BeforeReceivePipeline,omitempty"`
	BeforeSendPipeline    bool `json:"BeforeSendPipeline,omitempty"`
}

// Orchestration Model representing Orchestration
type Orchestration struct {
	AnalyticsEnabled      bool                          `json:"AnalyticsEnabled,omitempty"`
	ApplicationName       string                        `json:"ApplicationName,omitempty"`
	AssemblyName          string                        `json:"AssemblyName,omitempty"`
	Description           string                        `json:"Description,omitempty"`
	FullName              string                        `json:"FullName,omitempty"`
	Host                  string                        `json:"Host,omitempty"`
	ImplementedRoles      []string                      `json:"ImplementedRoles,omitempty"`
	InboundPorts          []OrchestrationInboundPort    `json:"InboundPorts,omitempty"`
	InvokedOrchestrations []string                      `json:"InvokedOrchestrations,omitempty"`
	OutboundPorts         []OrchestrationOutboundPort   `json:"OutboundPorts,omitempty"`
	Status                string                        `json:"Status,omitempty"`
	Tracking              *OrchestrationTrackingOptions `json:"Tracking,omitempty"`
	UsedRoles             []string                      `json:"UsedRoles,omitempty"`
}

// OrchestrationInboundPort Model representing an Inbound Port
type OrchestrationInboundPort struct {
	Binding     string `json:"Binding,omitempty"`
	Name        string `json:"Name,omitempty"`
	PortType    string `json:"PortType,omitempty"`
	ReceivePort string `json:"ReceivePort,omitempty"`
}

// OrchestrationOutboundPort Model representing an Outbound Port
type OrchestrationOutboundPort struct {
	Binding       string `json:"Binding,omitempty"`
	Name          string `json:"Name,omitempty"`
	PortType      string `json:"PortType,omitempty"`
	SendPort      string `json:"SendPort,omitempty"`
	SendPortGroup string `json:"SendPortGroup,omitempty"`
}

// OrchestrationTrackingOptions Model representing Tracking Options for Orchestration
type OrchestrationTrackingOptions struct {
	InboundMessageBody                 bool `json:"InboundMessageBody,omitempty"`
	MessageSendReceive                 bool `json:"MessageSendReceive,omitempty"`
	OrchestartionEvents                bool `json:"OrchestartionEvents,omitempty"`
	OutboundMessageBody                bool `json:"OutboundMessageBody,omitempty"`
	ServiceStartEnd                    bool `json:"ServiceStartEnd,omitempty"`
	TrackPropertiesForIncomingMessages bool `json:"TrackPropertiesForIncomingMessages,omitempty"`
	TrackPropertiesForOutgoingMessages bool `json:"TrackPropertiesForOutgoingMessages,omitempty"`
}

// ReceiveLocation Receive location model
type ReceiveLocation struct {
	Address             string                   `json:"Address,omitempty"`
	CustomData          string                   `json:"CustomData,omitempty"`
	Description         string                   `json:"Description,omitempty"`
	Enable              bool                     `json:"Enable,omitempty"`
	EncryptionCert      *CertificateInfo         `json:"EncryptionCert,omitempty"`
	FragmentMessages    string                   `json:"FragmentMessages,omitempty"`
	IsPrimary           bool                     `json:"IsPrimary,omitempty"`
	Name                string                   `json:"Name,omitempty"`
	PublicAddress       string                   `json:"PublicAddress,omitempty"`
	ReceiveHandler      string                   `json:"ReceiveHandler,omitempty"`
	ReceivePipeline     string                   `json:"ReceivePipeline,omitempty"`
	ReceivePipelineData string                   `json:"ReceivePipelineData,omitempty"`
	ReceivePortName     string                   `json:"ReceivePortName,omitempty"`
	Schedule            *ReceiveLocationSchedule `json:"Schedule,omitempty"`
	SendPipeline        string                   `json:"SendPipeline,omitempty"`
	SendPipelineData    string                   `json:"SendPipelineData,omitempty"`
	TransportType       string                   `json:"TransportType,omitempty"`
	TransportTypeData   string                   `json:"TransportTypeData,omitempty"`
}

// ReceiveLocationSchedule Service window
type ReceiveLocationSchedule struct {
	AutoAdjustToDaylightSaving bool   `json:"AutoAdjustToDaylightSaving,omitempty"`
	DaysOfWeek                 string `json:"DaysOfWeek,omitempty"`
	EndDate                    string `json:"EndDate,omitempty"`
	EndDateEnabled             bool   `json:"EndDateEnabled,omitempty"`
	FromTime                   string `json:"FromTime,omitempty"`
	LastDayOfMonth             bool   `json:"LastDayOfMonth,omitempty"`
	MonthDays                  string `json:"MonthDays,omitempty"`
	Months                     string `json:"Months,omitempty"`
	OrdinalDayOfWeek           string `json:"OrdinalDayOfWeek,omitempty"`
	OrdinalSchedule            string `json:"OrdinalSchedule,omitempty"`
	RecurFrom                  string `json:"RecurFrom,omitempty"`
	RecurInterval              int32  `json:"RecurInterval,omitempty"`
	RecurrenceSchType          string `json:"RecurrenceSchType,omitempty"`
	ScheduleIsOrdinal          bool   `json:"ScheduleIsOrdinal,omitempty"`
	ScheduleTimeZone           string `json:"ScheduleTimeZone,omitempty"`
	ServiceWindowEnabled       bool   `json:"ServiceWindowEnabled,omitempty"`
	StartDate                  string `json:"StartDate,omitempty"`
	StartDateEnabled           bool   `json:"StartDateEnabled,omitempty"`
	ToTime                     string `json:"ToTime,omitempty"`
}

// ReceivePort Model for Receive port
type ReceivePort struct {
	AnalyticsEnabled       bool      `json:"AnalyticsEnabled,omitempty"`
	ApplicationName        string    `json:"ApplicationName,omitempty"`
	CustomData             string    `json:"CustomData,omitempty"`
	Description            string    `json:"Description,omitempty"`
	InboundTransforms      []string  `json:"InboundTransforms,omitempty"`
	IsTwoWay               bool      `json:"IsTwoWay,omitempty"`
	Name                   string    `json:"Name,omitempty"`
	OutboundTransforms     []string  `json:"OutboundTransforms,omitempty"`
	PrimaryReceiveLocation string    `json:"PrimaryReceiveLocation,omitempty"`
	ReceiveLocations       []string  `json:"ReceiveLocations,omitempty"`
	Tracking               *Tracking `json:"Tracking,omitempty"`
}

// Schedule Service window
type Schedule struct {
	FromTime             string `json:"FromTime,omitempty"`
	ServiceWindowEnabled bool   `json:"ServiceWindowEnabled,omitempty"`
	ToTime               string `json:"ToTime,omitempty"`
}

// SendPort Model for Send port
type SendPort struct {
	AnalyticsEnabled     bool             `json:"AnalyticsEnabled,omitempty"`
	ApplicationName      string           `json:"ApplicationName,omitempty"`
	CustomData           string           `json:"CustomData,omitempty"`
	Description          string           `json:"Description,omitempty"`
	EncryptionCert       *CertificateInfo `json:"EncryptionCert,omitempty"`
	Filter               string           `json:"Filter,omitempty"`
	InboundTransforms    []string         `json:"InboundTransforms,omitempty"`
	IsDynamic            bool             `json:"IsDynamic,omitempty"`
	IsTwoWay             bool             `json:"IsTwoWay,omitempty"`
	Name                 string           `json:"Name,omitempty"`
	OrderPerAddress      bool             `json:"OrderPerAddress,omitempty"`
	OrderedDelivery      bool             `json:"OrderedDelivery,omitempty"`
	OutboundTransforms   []string         `json:"OutboundTransforms,omitempty"`
	PrimaryTransport     *TransportInfo   `json:"PrimaryTransport,omitempty"`
	Priority             int32            `json:"Priority,omitempty"`
	ReceivePipeline      string           `json:"ReceivePipeline,omitempty"`
	ReceivePipelineData  string           `json:"ReceivePipelineData,omitempty"`
	RouteFailedMessage   bool             `json:"RouteFailedMessage,omitempty"`
	SecondaryTransport   *TransportInfo   `json:"SecondaryTransport,omitempty"`
	SendPipeline         string           `json:"SendPipeline,omitempty"`
	SendPipelineData     string           `json:"SendPipelineData,omitempty"`
	Status               string           `json:"Status,omitempty"`
	StopSendingOnFailure bool             `json:"StopSendingOnFailure,omitempty"`
	Tracking             *Tracking        `json:"Tracking,omitempty"`
}

// SendPortGroup Model for SendPortGroup
type SendPortGroup struct {
	ApplicationName string   `json:"ApplicationName,omitempty"`
	CustomData      string   `json:"CustomData,omitempty"`
	Description     string   `json:"Description,omitempty"`
	Filter          string   `json:"Filter,omitempty"`
	Name            string   `json:"Name,omitempty"`
	SendPorts       []string `json:"SendPorts,omitempty"`
	Status          string   `json:"Status,omitempty"`
}

// Tracking Port tracking details
type Tracking struct {
	Body     *MessageBodyTracking     `json:"Body,omitempty"`
	Property *MessagePropertyTracking `json:"Property,omitempty"`
}

// TransportInfo The transport info.
type TransportInfo struct {
	Address           string    `json:"Address,omitempty"`
	OrderedDelivery   bool      `json:"OrderedDelivery,omitempty"`
	RetryCount        int32     `json:"RetryCount,omitempty"`
	RetryInterval     int32     `json:"RetryInterval,omitempty"`
	Schedule          *Schedule `json:"Schedule,omitempty"`
	SendHandler       string    `json:"SendHandler,omitempty"`
	TransportType     string    `json:"TransportType,omitempty"`
	TransportTypeData string    `json:"TransportTypeData,omitempty"`
}

// Message The message.
type Message struct {
	Id                     *uuid.UUID `json:"Id,omitempty"`
	Adapter                string     `json:"Adapter,omitempty"`
	CreationTime           string     `json:"CreationTime,omitempty"`
	HostName               string     `json:"HostName,omitempty"`
	MessageBoxDB           string     `json:"MessageBoxDB,omitempty"`
	MessageBoxServer       string     `json:"MessageBoxServer,omitempty"`
	MessageID              *uuid.UUID `json:"MessageID,omitempty"`
	MessageType            string     `json:"MessageType,omitempty"`
	OriginatorPartyName    string     `json:"OriginatorPartyName,omitempty"`
	OriginatorSecurityName string     `json:"OriginatorSecurityName,omitempty"`
	RetryCount             int32      `json:"RetryCount,omitempty"`
	ServiceClass           string     `json:"ServiceClass,omitempty"`
	ServiceInstanceID      *uuid.UUID `json:"ServiceInstanceID,omitempty"`
	ServiceName            string     `json:"ServiceName,omitempty"`
	ServiceStatus          string     `json:"ServiceStatus,omitempty"`
	ServiceTypeID          *uuid.UUID `json:"ServiceTypeID,omitempty"`
	Status                 string     `json:"Status,omitempty"`
	Submitter              string     `json:"Submitter,omitempty"`
	URI                    string     `json:"URI,omitempty"`
}
