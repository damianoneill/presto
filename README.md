# README

Skeleton OAS server implementation for MEF Presto.

## Generate

```bash
swagger generate server -f presto-nrp.yaml -A presto
```

## Deps

```bash
go get -u -f ./...
```

## Install

```bash
go install github.com/damianoneill/presto/...
```

## Run

```bash
$ presto-server --port=8080
2018/11/14 18:44:36 Serving presto at http://127.0.0.1:8080
```

## Test

```bash
$ curl http://127.0.0.1:8080/restconf/data/context
<string>operation tapi_common.GetDataContext has not yet been implemented</string>
```

## Docker

Build image 

```bash
docker build -t presto-server .
```

Run image

```bash
docker run -p 8080:8080 --name presto presto-server
```