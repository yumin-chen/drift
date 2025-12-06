#!/bin/bash
GOOS=js GOARCH=wasm go build -o drift.wasm .
