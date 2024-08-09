#!/bin/sh

if [ "$1" != "" ]
then
    terraform-docs markdown ./build --output-file ./docs/$1.md
else
    echo "Please provide filename."
    exit 1
fi