export CENTRAL_NODE=$(cat central_node.log)
if [ -z "$CENTRAL_NODE" ]; then
    export CENTRAL_NODE=$(cat central_node.log)
fi

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
blockchain_go getbalance -address $CENTRAL_NODE
blockchain_go getbalance -address $WALLET_1
blockchain_go getbalance -address $WALLET_2
blockchain_go getbalance -address $WALLET_3
blockchain_go getbalance -address $MINER_WALLET
