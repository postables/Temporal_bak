#! /bin/bash

# FOR PRODUCTION PLEASE USE CUSTOM KEYS
MODE="$2"

if [[ "$2" == "" ]]; then
    MODE="dev"
fi

if [[ "$MODE" != "dev" ]]; then
    if [[ "$MINIO_ACCESS_KEY" == "" ]]; then
        echo "MINIO_ACCESS_KEY environment variable empty"
        exit 1
    fi

    if [[ "$MINIO_SECRET_KEY" == "" ]]; then
        echo "MINIO_SECRET_KEY environment variable empty"
        exit 1
    fi
else
    MINIO_ACCESS_KEY="C03T49S17RP0APEZDK6M"
    MINIO_SECRET_KEY="q4I9t2MN/6bAgLkbF6uyS7jtQrXuNARcyrm2vvNA"
fi

if [[ "$DATA_DIR" == "" ]]; then    
    # set the defaults
    DATA_DIR="/mini_data"
fi

if [[ "$CONFIG_DIR" == "" ]]; then
    # set the defaults
    CONFIG_DIR="/mini_config"
fi


case "$1" in 

    pull-stable)
        sudo docker pull minio/minio
        ;;
    run-stable)
        docker run -p 9000:9000 --name minio1 -v "$DATA_DIR":/data -v "$CONFIG_DIR":/root/.minio minio/minio server /data
        ;;
    run-stable-custom)
        sudo docker run -p 9000:9000 --name minio1 \
        -e "MINIO_ACCESS_KEY=$MINIO_ACCESS_KEY" \
        -e "MINIO_SECRET_KEY=$MINIO_SECRET_KEY" \
        -v /mnt/data:"$DATA_DIR" \
        -v /mnt/config:"$CONFIG_DIR" \
        minio/minio server "$DATA_DIR"
        ;;
    *)
        echo "Invalid usage"
        echo "./build.sh [pull-stable|run-stable|run-stable-custom]"
        exit 1
        ;;

esac