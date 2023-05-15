# golog

![image](https://img.shields.io/badge/Docker-2CA5E0?style=for-the-badge&logo=docker&logoColor=white)
![image](https://img.shields.io/badge/Grafana-F2F4F9?style=for-the-badge&logo=grafana&logoColor=orange&labelColor=F2F4F9)
![image](https://img.shields.io/badge/grafana_loki-F2F4F9?style=for-the-badge&labelColor=2C3239&logo=data:image/png;base64,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)

A loki-based logging stack example with Go.

The docker-compose creates 4 services:

- `promtail` is the agent, responsible for gathering logs and sending them to Loki.
- `loki` is the main server, responsible for storing logs and processing queries.
- `grafana` for querying and displaying the logs.
- `app` is the go basic app that uses `fiber` as http engine and `zap` as logger.

## Usage

Start docker-compose services:

```sh
make local.docker.up.build
```

## Example 1

Access Grafana on http://localhost:3000
- go to explore tab on the left
- select container as label filter
- we can see an info log from the Go app

### Grafana explorer view :

<p align="center"><img src="examples/example1.png" alt="Loki Logo"></p>

## Example 2

The Go app is running on http://localhost:8000

In the [tests](/app/tests.go) file, there are handlers to trigger several logs levels:

```sh
curl localhost:8000/tests/debug     // trigger a debug log
curl localhost:8000/tests/info      // trigger an info log
curl localhost:8000/tests/warn      // trigger a warning log
curl localhost:8000/tests/error     // trigger an error log
curl localhost:8000/tests/various   // trigger all levels several times
```

### Grafana explorer view :

<p align="center"><img src="examples/example2.png" alt="Loki Logo"></p>