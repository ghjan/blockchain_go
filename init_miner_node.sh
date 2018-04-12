#!/bin/bash
export NODE_ID=3002
if [ -z "$MINER_WALLET" ]; then
    if [ -e "wallet4.log" -a -e "blockchain_3002.db" ]; then
        export MINER_WALLET=$(cat wallet4.log)
        echo "wallet4.log and blockchain_3002.db are already existed, just to restore MINER_WALLET:" $MINER_WALLET
    else
        cp blockchain_genesis.db blockchain_3002.db
        w_results=$(blockchain_go createwallet)
        export MINER_WALLET=${w_results//Your new address: /}
        echo $MINER_WALLET >wallet4.log
    fi
else
     echo "MINER_WALLET is already existed:" $MINER_WALLET
fi
