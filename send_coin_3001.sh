#!/bin/bash
if [ -z "$WALLET_1" ]; then
    export WALLET_1=$(cat wallet1.log)
fi

if [ -z "$WALLET_2" ]; then
    export WALLET_2=$(cat wallet2.log)
fi

if [ -z "$WALLET_3" ]; then
    export WALLET_3=$(cat wallet3.log)
fi

if [ -z "$MINER_WALLET" ]; then
    export MINER_WALLET=$(cat wallet4.log)
fi
blockchain_go send -from $WALLET_1 -to $WALLET_3 -amount 1
blockchain_go send -from $WALLET_2 -to $MINER_WALLET -amount 2
