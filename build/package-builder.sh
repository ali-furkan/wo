#!/bin/sh

set -eu

mkdir releases

Release() {
    echo "Creating '$2'"
    
    mkdir release
    cd release && mkdir bin && cd ..

    sudo cp -f ./README.md ./release/README.md
    sudo cp -f ./bin/$1 ./release/bin/$2

    cd release
    sudo zip -r ../$2.zip ./
    cd ..
    sudo rm -fr release
}

echo "📦 Preaparing Releases"

# Darwin
Release darwin/wo_amd64 wo_darwin_amd64
Release darwin/wo_arm64 wo_darwin_arm64

#Linux
Release linux/wo_386 wo_linux_386
Release linux/wo_amd64 wo_linux_amd64
Release linux/wo_arm wo_linux_arm
Release linux/wo_arm64 wo_linux_arm64

# FreeBSD
Release freebsd/wo_386 wo_freebsd_386
Release freebsd/wo_amd64 wo_freebsd_amd64
Release freebsd/wo_arm wo_freebsd_arm
Release freebsd/wo_arm64 wo_freebsd_arm64

# Windows
Release win/wo_386.exe wo_win_386
Release win/wo_amd64.exe wo_win_amd64
Release win/wo_arm.exe wo_win_arm

# Pack all built packages & Removes bin folder
sudo zip -r ./wo_all_platforms.zip ./bin

sudo rm -fr bin

echo "✅ Successfully Created Release files"