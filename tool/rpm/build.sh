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
    docker build -f ./tool/rpm/Dockerfile --rm . 
}

$COMMAND