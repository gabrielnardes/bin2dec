#!/bin/bash

podman run -v $PWD/src/:/go/src/bin2dec/cmd/api/src -v $PWD/../../pkg/:/go/src/bin2dec/pkg -v $PWD/../../go.mod:/go/src/bin2dec/go.mod  -p 8080:8080 --rm  bin2dec
