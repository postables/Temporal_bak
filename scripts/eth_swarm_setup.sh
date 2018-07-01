#! /bin/bash

# Used to setup a swarm connection

ENV="dev"
DATADIR="/home/solidity/swarm_data"

case "$ENV" in 

    "dev")
        echo "enter password for eth account"
        read -r password
        echo "$password" > "$DATADIR/password"
        geth --dev account new --datadir="$DATADIR" --password "$DATADIR/password" 2>&1 | tee "$DATADIR/account_creation.log"
        ETH_ADDRESS=$(grep '^Address' "$DATADIR/account_creation.log" | awk '{print $2}' | tr -d '{}')
        swarm --datadir="$DATADIR" --maxpeers 0 --bzzaccount "$ETH_ADDRESS" --password "$DATADIR/password"
        exit 0
        ;;
    "prod")
        echo "prod"
        exit 0
        ;;
    *)
        echo "invalid environment"
        echo "valid environments: 'dev' 'prod'"
        exit 1
        ;;
esac