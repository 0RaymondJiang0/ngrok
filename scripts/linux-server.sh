#!/bin/bash
GOBIN= GOOS=linux GOARCH=amd64 make release-server

# GOBIN= GOOS=linux GOARCH=amd64 GODEBUG="x509ignoreCN=0" make release-server