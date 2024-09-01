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

function build-deb-local(){
    ./tool/deb/build.sh run_on_local_container
}

function build-rpm-local(){
    ./tool/rpm/build.sh run_on_local_container
}

function build-container-up(){
    docker-compose -f ./tool/docker-compose.yaml up --build --force-recreate -d
}

function build-deb(){
    VERSION=$(cat ./VERSION)
    DATETIME=$(date "+%a, %d %b %Y %H:%M:%S +0900" --date="2024-09-01 12:00:00")
    ARCH=$(uname -m)
    mkdir -p ./tool/deb/debian
    cp -r ./bin ./tool/deb/debian
    eval "echo \"$(cat ./tool/deb/changelog.template)\"" > ./tool/deb/debian/changelog
    eval "echo \"$(cat ./tool/deb/control.template)\"" > ./tool/deb/debian/control
    eval "echo \"$(cat ./tool/deb/copyright.template)\"" > ./tool/deb/debian/copyright
    cd ./tool/deb/ && dpkg-buildpackage -us -uc
}

$COMMAND