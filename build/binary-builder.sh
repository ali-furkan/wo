#!/bin/bash

set -eu

ex_dir="$(dirname "$0")"

. "$ex_dir/platform-arch.sh"

build_binary() {
    echo "📑 Wo is building binaries of '$1'"
    for arch in $@
    do
        if [ $1 != $arch ]
        then
            echo " - Building $1/$arch"
            local file="./bin/$1/wo_${1}_$arch"
            if [ $1 = "windows" ]
            then
                file+=".exe"
            fi
            GOOS=$1 GOARCH=$arch go build -o $file
        fi
    done
}

for p in ${platforms[@]}
do
    archs=$(get_platform_arch $p)
    build_binary $p ${archs[@]} 
done
