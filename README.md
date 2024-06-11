# hackathon-dash-moi-ca

## Generate GO CLI

```
go build -o dashmoica
```

## Jsonnet directory

### Install Grafonnet

```console
jb install github.com/grafana/grafonnet/gen/grafonnet-latest@main
```

### Create a new dashboard

```jsonnet
// dashboard.jsonnet
local grafonnet = import 'github.com/grafana/grafonnet/gen/grafonnet-latest/main.libsonnet';

grafonnet.dashboard.new('My Dashboard')
```

### Generate the dashboard

```console
jsonnet -J vendor dashboard.jsonnet
```

