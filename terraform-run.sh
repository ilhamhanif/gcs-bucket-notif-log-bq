#!/bin/sh

# Initialization
if [ "$1" == "init" ]
then

    terraform -chdir=build init $@
    echo "Terraform Initialized."

# Building Block
elif [ "$1" == "build" ]
then

    if [ "$3" != "" ]
    then
        envVariableFile=$3.tfvars
    else # Error Handler
        echo "Please state environment variable .tfvars file name (without extension)."
        exit 1
    fi

    if [ "$2" == "sync" ] # Sync
    then
        echo "Syncing current state infrastructure with target."
        terraform -chdir=build apply -var-file $envVariableFile -auto-approve -refresh-only
    elif [ "$2" == "deploy" ] # Deploy
    then
        echo "Deploying infrastructure."
        terraform -chdir=build apply -var-file $envVariableFile -auto-approve
    else # Error Handler
        echo "Please define command [sync/deploy]."
        exit 1
    fi

else # Handler

    echo "Please define command [init/build]."
    exit 1

fi