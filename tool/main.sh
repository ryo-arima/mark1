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
    TAG_NAME="v${VERSION}"
    ARCH=$(uname -m)
    REPO="ryo-arima/mark1"
    FILE_PATH="./tool/*.deb"
    FILE_NAME=$(basename $FILE_PATH)
    ASSET_ID=$(gh release view $TAG_NAME --json assets --jq ".assets | map(select(.name == \"$(basename $FILE_PATH)\")) | .[0].id" -R $REPO)
    if gh release list | grep -q "$TAG_NAME"; then
        echo "Release exists"
        if [ -n "$ASSET_ID" ]; then
          echo "Asset exists"
          gh release delete-asset $TAG_NAME $FILE_NAME -y
          gh release upload $TAG_NAME $FILE_PATH -R $REPO
        else
          echo "Asset does not exist"
          gh release upload $TAG_NAME $FILE_PATH -R $REPO
        fi
    else
        echo "Release does not exist"                  
        gh release create $TAG_NAME $FILE_PATH --title "$TAG_NAME" --notes "$TAG_NAME" --prerelease
    fi
}

function build-rpm(){
    VERSION=$(cat ./VERSION)
    DATETIME=$(date -d "2024-09-01" +"%a %b %d %Y")
    ARCH=$(uname -m)
    BASE_DIR=mark1-${ARCH}-${VERSION}
    rpmdev-setuptree
    echo $HOME
    ls -al ${HOME}
    pwd
    eval "echo \"$(cat ./tool/rpm/mark1.spec.template)\"" > ${HOME}/rpmbuild/SPECS/mark1.spec
    cd ./tool/rpm/ && \
    mkdir -p ${BASE_DIR} && \
    cp -r ../../bin ${BASE_DIR} && \
    mkdir -p ${BASE_DIR}/etc && \
    cp ../../etc/main.yaml.template ${BASE_DIR}/etc/main.yaml && \
    mkdir -p ${BASE_DIR}/systemd/system && \
    cp ../systemd/mark1.service ${BASE_DIR}/systemd/system && \
    tar -czvf mark1-${ARCH}-${VERSION}.tar.gz ${BASE_DIR} && \
    cp mark1-${ARCH}-${VERSION}.tar.gz ${HOME}/rpmbuild/SOURCES && \
    cd ../..
    rpmbuild --define '_build_id_links none' --define 'debug_package %{nil}' -ba ${HOME}/rpmbuild/SPECS/mark1.spec
}

function push-rpm(){
    VERSION=$(cat ./VERSION)
    TAG_NAME="v${VERSION}"
    ARCH=$(uname -m)
    REPO="ryo-arima/mark1"
    FILE_PATH="${HOME}/rpmbuild/RPMS/${ARCH}/mark1-${VERSION}-${ARCH}.${ARCH}.rpm"
    FILE_NAME=$(basename $FILE_PATH)
    ASSET_ID=$(gh release view $TAG_NAME --json assets --jq ".assets | map(select(.name == \"$(basename $FILE_PATH)\")) | .[0].id" -R $REPO)
    if gh release list | grep -q "$TAG_NAME"; then
        echo "Release exists"
        if [ -n "$ASSET_ID" ]; then
          echo "Asset exists"
          gh release delete-asset $TAG_NAME $FILE_NAME -y
          gh release upload $TAG_NAME $FILE_PATH -R $REPO
        else
          echo "Asset does not exist"
          gh release upload $TAG_NAME $FILE_PATH -R $REPO
        fi
    else
        echo "Release does not exist"                  
        gh release create $TAG_NAME $FILE_PATH --title "$TAG_NAME" --notes "$TAG_NAME" --prerelease
    fi
}

$COMMAND