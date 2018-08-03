#!/bin/bash

export NAMESPACE="default"
export DOMAIN="default.example.com"
export REGISTRY_CREDENTIALS="build-bot"
export REPOSITORY="jszroberto"
export SERVICE_ADDRESS="http://localhost:8080"

pushd cmd/kleincf
go build && ./kleincf
popd
