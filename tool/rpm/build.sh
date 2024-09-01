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

$COMMAND