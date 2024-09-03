#!/bin/bash

COMMAND="$1"
ARGs="$@"
PARAMs=(${ARGs// / })


function build_server(){
    go build -v -o bin/server cmd/server/main.go
}

function build_client_root(){
    go build -v -o bin/mark1-ctl cmd/client/root/main.go
}

function build_client_anonymous(){
    go build -v -o bin/anonymous-client cmd/client/anonymous/main.go
}

function build_client_app(){
    go build -v -o bin/app-client cmd/client/app/main.go
}

function build_client_admin(){
    go build -v -o bin/admin-client cmd/client/admin/main.go
}

function build_client(){
    build_client_root
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
    cp -r ./bin ./tool/deb
    cp -r ./tool/systemd ./tool/deb
    mkdir -p ./tool/deb/etc
    cp etc/main.yaml.template ./tool/deb/etc/main.yaml
    eval "echo \"$(cat ./tool/deb/changelog.template)\"" > ./tool/deb/debian/changelog
    eval "echo \"$(cat ./tool/deb/control.template)\"" > ./tool/deb/debian/control
    eval "echo \"$(cat ./tool/deb/copyright.template)\"" > ./tool/deb/debian/copyright
    eval "echo \"$(cat ./tool/deb/rules.template)\"" > ./tool/deb/debian/rules
    cp ./tool/deb/mark1.install ./tool/deb/debian/mark1.install
    chmod +x ./tool/deb/debian/rules

    cd ./tool/deb/ && \
    fakeroot dpkg-buildpackage -us -uc && \
    cd ../..
}

function push-deb(){
    VERSION=$(cat ./VERSION)
    RELEASE="debian-${VERSION}"
    if gh release list | grep -q "v$RELEASE"; then
        gh release delete v$RELEASE -y
        gh release create v$RELEASE ./tool/*.deb --title "v$RELEASE" --notes "v$RELEASE" --prerelease
    else                    
        gh release create v$RELEASE ./tool/*.deb --title "v$RELEASE" --notes "v$RELEASE" --prerelease
    fi
}

$COMMAND