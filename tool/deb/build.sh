#!/bin/bash

COMMAND="$1"
ARGs="$@"
PARAMs=(${ARGs// / })

function setup(){
    echo "Setup"
}

function test(){
    echo "Test"
}

function run(){
    setup
    test
}

function run_on_local_container(){
    echo "Run on local container"
    ARCH=$(uname -m)
    if [[ "$ARCH" == *arm* ]]; then
        docker build -f ./tool/deb/Dockerfile --rm --build-arg ARCH='arm' . 
    else
        docker build -f ./tool/deb/Dockerfile --rm --build-arg ARCH='base' . 
    fi
}

$COMMAND