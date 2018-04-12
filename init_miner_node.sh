#!/bin/bash
export NODE_ID=3002
if [ -e "blockchain_3002.db" -a -n "$MINER_WALLET" ]; then
     echo "MINER_WALLET is already existed:" $MINER_WALLET
     exit 1
elif [ -e "wallet4.log" -a -e "blockchain_3002.db" ]; then
        export MINER_WALLET=$(cat wallet4.log)
        echo "wallet4.log and blockchain_3002.db are already existed, just to restore MINER_WALLET:" $MINER_WALLET
else
    cp blockchain_genesis.db blockchain_3002.db
    w_results=$(blockchain_go createwallet)
    export MINER_WALLET=${w_results//Your new address: /}
    echo $MINER_WALLET >wallet4.log
fi
