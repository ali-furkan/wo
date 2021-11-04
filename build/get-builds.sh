#!/bin/sh

set -eu

darwin=(amd64 arm)
freebsd=(386 amd64 arm)
linux=(386 amd64 arm arm64)
windows=(386 amd64 arm)

builds=""
get_build_names() {
    for arch in $@
    do
        if [ $arch != $1 ]
        then
            builds+="wo_$1_${arch}.zip "
        fi
    done
}

get_build_names darwin ${darwin[@]}
get_build_names freebsd ${freebsd[@]}
get_build_names linux ${linux[@]}
get_build_names windows ${windows[@]}

echo $builds