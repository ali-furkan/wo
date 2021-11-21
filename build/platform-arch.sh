#!/bin/bash

platforms=(darwin freebsd linux windows)
darwin=(amd64 arm64)
freebsd=(386 amd64 arm arm64)
linux=(386 amd64 arm arm64)
windows=(386 amd64 arm arm64)

get_platform_arch() {
    local archs=""
    local platform=()

    case $1 in
        "darwin") platform=${darwin[@]}
        ;;
        "freebsd") platform=${freebsd[@]}
        ;;
        "linux") platform=${linux[@]}
        ;;
        "windows") platform=${windows[@]}
        ;;
        *) 
            echo "invalid platform name"
            exit 1
            ;;
    esac

    for arch in ${platform[@]}
    do
        archs+="$arch "
    done 

    echo "$archs"
}

get_all() {
    local res=""

    for platform in ${platforms[@]}
    do
        local archs=$(get_platform_arch $platform)
        for a in ${archs[@]}
        do
            res+="$platform/$a "
        done
    done

    echo "$res"
}

main() {
    if [ -z $1 ]
    then
        get_all
        exit 0
    fi

    get_platform_arch $1
}

if [ $(basename $0) = "platform-arch.sh" ]
then
    main "${@}"
fi
