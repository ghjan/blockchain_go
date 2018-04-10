#!/bin/bash
export NODE_ID=3001
export WALLET_1=$(cat wallet1.log)
export WALLET_2=$(cat wallet2.log)
export WALLET_3=$(cat wallet3.log)
export WALLET_4=$(cat wallet4.log)
blockchain_go send -from $WALLET_1 -to $WALLET_3 -amount 10
# blockchain_go send -from $WALLET_2 -to $WALLET_4 -amount 10
