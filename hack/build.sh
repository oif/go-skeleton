#!/bin/sh
TARGET_NAME=skeleton
BUILD_DIR=./_build
VERSION=`grep -i "^VER" CHANGELOG | tail -1 | cut -d " " -f2`
BUILD_DATE=`date -u +'%Y-%m-%dT%H:%M:%SZ'`
GIT_REVISION=`git rev-parse --short HEAD`
REPO=github.com/oif/go-skeleton
SCRIPT_PWD=$(dirname ${BASH_SOURCE})

BUILD_PLATFORMS=("windows/amd64" "darwin/amd64" "linux/amd64")

buildBoredBinary() {
    RELEASE_NAME=$3
    CGO_ENABLED=0 GOOS="$1" GOARCH="$2" go build -i \
    -o $BUILD_DIR/$RELEASE_NAME $REPO/cmd

    if [ $? -ne 0 ]; then
        echo "Failed to build $RELEASE_NAME"
        exit 1
    fi

    echo "Build $RELEASE_NAME, OS is $1, Arch is $2"
}

rm -rf $BUILD_DIR

for platform in "${BUILD_PLATFORMS[@]}"
do
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}
    output_name=$TARGET_NAME'-'$VERSION'-'$GOOS'-'$GOARCH

    if [ $GOOS = "windows" ]; then
        output_name+='.exe'
    fi
    buildBoredBinary $GOOS $GOARCH $output_name
done
