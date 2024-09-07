#!/bin/bash

COMMAND="$1"
ARGs="$@"
PARAMs=(${ARGs// / })


function test(){
  go test -cover -v ./test/...
}

$COMMAND