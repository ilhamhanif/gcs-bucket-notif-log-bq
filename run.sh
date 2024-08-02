#!/bin/sh

if [ "$1" == "init" ]
then
    if [ "$2" == "" ]
    then
        echo "Initialize Terraform."
        terraform -chdir=build init
    else
        echo "Upgrading current Terraform infrastructure state."
        terraform -chdir=build init -upgrade
    fi
elif [ "$1" == "build" ]
then
    if [ "$3" == "" ]
    then
        echo "Please state environment variable .tfvars file"
        exit 1
    else
        envVariableFile=$3
    fi

    if [ "$2" == "sync" ]
    then
        echo "Syncing current state infrastructure with target."
        terraform -chdir=build apply -var-file $envVariableFile -auto-approve -refresh-only
    elif [ "$2" == "deploy" ]
    then
        echo "Deploying infrastructure."
        terraform -chdir=build apply -var-file $envVariableFile -auto-approve
    else
        echo "Please define command [sync/deploy]"
        exit 1
    fi
else
    echo "Please define command [init/build]"
    exit 1
fi