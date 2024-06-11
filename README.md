# hackathon-dash-moi-ca

## Install Grafonnet

```console
jb install github.com/grafana/grafonnet/gen/grafonnet-latest@main
```

## Create a new dashboard

```jsonnet
// dashboard.jsonnet
local grafonnet = import 'github.com/grafana/grafonnet/gen/grafonnet-latest/main.libsonnet';

grafonnet.dashboard.new('My Dashboard')
```

## Generate the dashboard

```console
cd jsonnet
jsonnet -J vendor dashboard.jsonnet
```

