#!/bin/bash

COMMAND="$1"
ARGs="$@"
PARAMs=(${ARGs// / })

function setup(){
    echo "Setup"
}

function run(){
    setup
}

$COMMAND