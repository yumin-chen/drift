#!/bin/bash

GOOS=js GOARCH=wasm go build -o drift.wasm drift_wasm.go
