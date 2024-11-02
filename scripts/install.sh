#!/bin/bash

go get golang.org/x/mobile/bind
go install golang.org/x/mobile/cmd/gomobile@latest
gomobile init
