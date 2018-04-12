#!/bin/bash
export NODE_ID=3000

if [ -z "$CENTRAL_NODE" -a -e "central_node.log" ]; then
    export CENTRAL_NODE=$(cat central_node.log)
fi

if [ -z "$WALLET_1" -a -e "wallet1.log" ]; then
    export WALLET_1=$(cat wallet1.log)
fi

if [ -z "$WALLET_2" -a -e "wallet2.log" ]; then
    export WALLET_2=$(cat wallet2.log)
fi

if [ -z "$WALLET_3" -a -e "wallet3.log" ]; then
    export WALLET_3=$(cat wallet3.log)
fi

if [ -n "$CENTRAL_NODE" -a -n "$WALLET_1" ]; then
    blockchain_go send -from $CENTRAL_NODE -to $WALLET_1 -amount 10 -mine
fi
if [ -n "$CENTRAL_NODE" -a -n "$WALLET_2" ]; then
    blockchain_go send -from $CENTRAL_NODE -to $WALLET_2 -amount 10 -mine
fi
