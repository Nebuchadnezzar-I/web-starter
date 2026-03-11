#!/usr/bin/env bash

set -e

echo "Starting project..."

templ generate
go run ./cmd/web
