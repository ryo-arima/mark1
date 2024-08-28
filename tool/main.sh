#!/bin/bash

COMMAND="$1"
ARGs="$@"
PARAMs=(${ARGs// / })

function build_server(){
    go build -o bin/server cmd/server/main.go
}

function build_client_anonymous(){
    go build -o bin/anonymous-client cmd/client/anonymous/main.go
}

function build_client_app(){
    go build -o bin/app-client cmd/client/app/main.go
}

function build_client_admin(){
    go build -o bin/admin-client cmd/client/admin/main.go
}

function build_client(){
    build_client_anonymous
    build_client_app
    build_client_admin
}

function build(){
    build_server
    build_client
}

$COMMAND