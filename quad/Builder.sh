#!/bin/bash

# Build all Go files
go build quadA.go
go build quadB.go
go build quadC.go
go build quadD.go
go build quadE.go
go build quadchecker
mkdir quad
mv quad*.go main.go ./quad