#!/bin/sh

if [ "$1" != "" ]
then
    if ! [ -d ./build/docs ]; then
        mkdir ./build/docs
    fi
    terraform-docs markdown ./build --output-file ./docs/$1.md

    if ! [ -d ./docs/md ]; then
        mkdir ./docs/md
    fi
    mv ./build/docs/$1.md ./docs/md/$1.md

    rm -rf ./build/docs

else
    echo "Please provide filename."
    exit 1
fi