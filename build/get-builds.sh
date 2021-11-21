#!/bin/bash

set -eu

ex_dir="$(dirname "$0")"
. "$ex_dir/platform-arch.sh"

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

for p in ${platforms[@]}
do
    archs=$(get_platform_arch $p)
    get_build_names $p ${archs[@]}
done

echo $builds