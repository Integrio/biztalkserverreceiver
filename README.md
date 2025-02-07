# biztalkserverreceiver
BizTalk Server Receiver

## 1. **Suspended Instances**  

**Metric Type**: Count (sum)  
**Description**: Tracks the number of suspended instances in BizTalk Server.

### Possible Values:

- Count of suspended instances.

### Metric Example:

```json
{
  "name": "biztalk.suspended_instances",
  "description": "Count of suspended instances in BizTalk Server.",
  "unit": "1",
  "sum": {
    "dataPoints": [
      {
        "attributes": [
          {
            "key": "suspended_instances.status",
            "value": { "stringValue": "Suspended" }
          }
        ],
        "startTimeUnixNano": "1738920605382705506",
        "timeUnixNano": "1738920606390452840",
        "asInt": "5"  # Count of suspended instances
      }
    ]
  }
}
```

---

## 2. **Suspended Messages**

**Metric Type**: Count (sum)  
**Description**: Tracks the number of suspended messages in BizTalk Server.

### Possible Values:

- Count of suspended messages.

### Metric Example:

```json
{
  "name": "biztalk.suspended_messages",
  "description": "Count of suspended messages in BizTalk Server.",
  "unit": "1",
  "sum": {
    "dataPoints": [
      {
        "attributes": [
          {
            "key": "suspended_messages.status",
            "value": { "stringValue": "Suspended" }
          }
        ],
        "startTimeUnixNano": "1738920605382705506",
        "timeUnixNano": "1738920606390452840",
        "asInt": "3"  # Count of suspended messages
      }
    ]
  }
}
```

---

## 3. **Orchestrations**  

**Metric Type**: Status (gauge)  
**Description**: Tracks the status of orchestrations in BizTalk Server.

### Possible Values:

- `Enlisted`
- `Unenlisted`
- `Started`

### Metric Example:

```json
{
  "name": "biztalk.orchestrations_status",
  "description": "Status of orchestrations in BizTalk Server.",
  "unit": "1",
  "gauge": {
    "dataPoints": [
      {
        "attributes": [
          {
            "key": "orchestration.status",
            "value": { "stringValue": "Started" }
          }
        ],
        "startTimeUnixNano": "1738920605382705506",
        "timeUnixNano": "1738920606390452840",
        "asInt": "1"  # Started (1), Unenlisted (0), Enlisted (2)
      }
    ]
  }
}
```

---

## 4. **Receive Locations**  

**Metric Type**: Enabled (gauge)  
**Description**: Tracks whether a receive location is enabled.

### Possible Values:

- `True` (Enabled)
- `False` (Disabled)

### Metric Example:

```json
{
  "name": "biztalk.receive_locations_enabled",
  "description": "Enable status of receive locations in BizTalk Server.",
  "unit": "1",
  "gauge": {
    "dataPoints": [
      {
        "attributes": [
          {
            "key": "receive_location.name",
            "value": { "stringValue": "ReceiveLocation1" }
          },
          {
            "key": "receive_location.enabled",
            "value": { "boolValue": true }  # true for enabled, false for disabled
          }
        ],
        "startTimeUnixNano": "1738920605382705506",
        "timeUnixNano": "1738920606390452840",
        "asInt": "1"  # Enabled = 1, Disabled = 0
      }
    ]
  }
}
```

---

## 5. **Send Ports**  

**Metric Type**: Status (gauge)  
**Description**: Tracks the status of send ports in BizTalk Server.

### Possible Values:

- `Bound`
- `Stopped`
- `Started`

### Metric Example:
```json
{
  "name": "biztalk.send_ports_status",
  "description": "Status of send ports in BizTalk Server.",
  "unit": "1",
  "gauge": {
    "dataPoints": [
      {
        "attributes": [
          {
            "key": "send_port.name",
            "value": { "stringValue": "SendPort1" }
          },
          {
            "key": "send_port.status",
            "value": { "stringValue": "Started" }
          }
        ],
        "startTimeUnixNano": "1738920605382705506",
        "timeUnixNano": "1738920606390452840",
        "asInt": "1"  # Bound (0), Stopped (2), Started (1)
      }
    ]
  }
}
```

---

## 6. **Send Port Groups**  

**Metric Type**: Status (gauge)  
**Description**: Tracks the status of send port groups in BizTalk Server.

### Possible Values:

- `Bound`
- `Stopped`
- `Started`

### Metric Example:

```json
{
  "name": "biztalk.sendport_groups_status",
  "description": "Status of send port groups in BizTalk Server.",
  "unit": "1",
  "gauge": {
    "dataPoints": [
      {
        "attributes": [
          {
            "key": "send_port_group.name",
            "value": { "stringValue": "SendPortGroup1" }
          },
          {
            "key": "send_port_group.status",
            "value": { "stringValue": "Bound" }
          }
        ],
        "startTimeUnixNano": "1738920605382705506",
        "timeUnixNano": "1738920606390452840",
        "asInt": "0"  # Bound (0), Stopped (2), Started (1)
      }
    ]
  }
}
```

---

## 7. **Host Instances**  

**Metric Type**: Status (gauge)  
**Description**: Tracks the status of host instances in BizTalk Server.

### Possible Values:

- `Stopped`
- `Unknown`
- `Running`

### Metric Example:

```json
{
  "name": "biztalk.host_instances_status",
  "description": "Status of host instances in BizTalk Server.",
  "unit": "1",
  "gauge": {
    "dataPoints": [
      {
        "attributes": [
          {
            "key": "host_instance.name",
            "value": { "stringValue": "HostInstance1" }
          },
          {
            "key": "host_instance.status",
            "value": { "stringValue": "Running" }
          }
        ],
        "startTimeUnixNano": "1738920605382705506",
        "timeUnixNano": "1738920606390452840",
        "asInt": "2"  # Stopped (0), Unknown (1), Running (2)
      }
    ]
  }
}
```

