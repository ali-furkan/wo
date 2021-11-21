#!/bin/bash

set -eu

ex_dir="$(dirname "$0")"
. "$ex_dir/platform-arch.sh"

package-binary() {
    echo "📜 Bundling '$2' binary"

    mkdir release
    cd release && mkdir bin && cd ..

    sudo cp -f ./README.md ./release/README.md
    sudo cp -f ./bin/$1 ./release/bin/$2

    cd release
    sudo zip -r ../$2.zip ./
    cd ..
    sudo rm -rf release
}

echo "📦 Preaparing Releases"

for p in ${platforms[@]}
do
    archs=$(get_platform_arch $p)
    for a in ${archs[@]}
    do
        if [ $p != "windows" ]
        then
            package-binary "$p/wo_${p}_$a" "wo_${p}_$a"
        else
            package-binary "$p/wo_${p}_$a.exe" "wo_${p}_$a"
        fi
    done
done

# Packing all built packages & Removes bin folder
sudo zip -r ./wo_all_platforms.zip ./bin

sudo rm -fr bin

echo "✅ Successfully Bundled Release files"
