export NODE_ID=3001
if [ -z "$WALLET_1" ]; then
    w_results=$(blockchain_go createwallet)
    export WALLET_1=${w_results//Your new address: /}
    echo $WALLET_1 >wallet1.log
fi

if [ -z "$WALLET_2" ]; then
    w_results=$(blockchain_go createwallet)
    export WALLET_2=${w_results//Your new address: /}
    echo $WALLET_2 >wallet2.log
fi

if [ -z "$WALLET_3" ]; then
    w_results=$(blockchain_go createwallet)
    export WALLET_3=${w_results//Your new address: /}
    echo $WALLET_3 >wallet3.log
fi

cp blockchain_genesis.db blockchain_3001.db
