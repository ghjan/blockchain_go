#!/bin/bash
export NODE_ID=3000
export WALLET_1=$(cat wallet1.log)
export WALLET_2=$(cat wallet2.log)
export WALLET_3=$(cat wallet3.log)
blockchain_go send -from $CENTRAL_NODE -to $WALLET_1 -amount 10 -mine
blockchain_go send -from $CENTRAL_NODE -to $WALLET_2 -amount 10 -mine
