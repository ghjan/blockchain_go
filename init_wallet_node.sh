#!/bin/bash
export NODE_ID=3001

if [ -e "blockchain_genesis.db" ]; then
    cp blockchain_genesis.db blockchain_3001.db
else
    echo "blockchain_genesis.db is not existed, please run create_blockchain.sh firstly!"
    exit 1
fi

if [  -z "$WALLET_1" -o ! -e "wallet1.log" ]; then
    w_results=$(blockchain_go createwallet)
    export WALLET_1=${w_results//Your new address: /}
    echo $WALLET_1 >wallet1.log
fi

if [  -z "$WALLET_2" -o ! -e "wallet2.log" ]; then
    w_results=$(blockchain_go createwallet)
    export WALLET_2=${w_results//Your new address: /}
    echo $WALLET_2 >wallet2.log
fi

if [  -z "$WALLET_3" -o ! -e "wallet3.log" ]; then
    w_results=$(blockchain_go createwallet)
    export WALLET_3=${w_results//Your new address: /}
    echo $WALLET_3 >wallet3.log
fi
