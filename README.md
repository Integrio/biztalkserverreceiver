
# OpenTelemetry Collector Biztalk Server Receiver

The `biztalkserverreceiver` enables the collection of metrics from Biztalk Server. For details about which metrics that are produced, see [documentation](./documentation.md).

## Configuration

Using no auth:
```yaml
biztalkserver:
  collection_interval: 1m
  endpoint: http://biztalkserver01.local
```

Using basic auth:
```yaml
biztalkserver:
  collection_interval: 1m
  endpoint: http://biztalkserver01.local
  username: ${env:USERNAME}
  password: ${env:PASSWORD}
  auth: basic
```

Using ntlm auth:
```yaml
biztalkserver:
  collection_interval: 1m
  endpoint: http://biztalkserver01.local
  username: ${env:USERNAME}
  password: ${env:PASSWORD}
  auth: ntlm
```


## Contributing
### Developing
Clone the repository\
Make changes to the receiver\
Run `make build-local` to build a local collector to ./collector\
TBC..