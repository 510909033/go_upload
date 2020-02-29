#!/bin/sh

env=$1
echo "env=$env"

if [[ $env == "dev" ]];then
    export ucm_remote=172.24.19.102:9602
    export go_env_file=env_dev
    make dev
fi

if [[ $env == "online" ]];then
    export go_env_file=env_online
    make online
fi

echo "over"
