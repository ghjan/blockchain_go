#!/bin/bash
export NODE_ID=3000

if [ -z "$WALLET_1" ]; then
    export WALLET_1=$(cat wallet1.log)
fi

if [ -z "$WALLET_2" ]; then
    export WALLET_2=$(cat wallet2.log)
fi

if [ -z "$WALLET_3" ]; then
    export WALLET_3=$(cat wallet3.log)
fi
export CENTRAL_NODE=$(cat central_node.log)
blockchain_go send -from $CENTRAL_NODE -to $WALLET_1 -amount 10 -mine
blockchain_go send -from $CENTRAL_NODE -to $WALLET_2 -amount 10 -mine
