export CENTRAL_NODE=$(cat central_node.log)
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

if [ -z "$MINER_WALLET" -a -e "wallet4.log" ]; then
    export MINER_WALLET=$(cat wallet4.log)
fi

if [ -n "$CENTRAL_NODE"  -a -e "central_node.log" ]; then
    echo "balance for CENTRAL_NODE"
    blockchain_go getbalance -address $CENTRAL_NODE
fi
if [ -n "$WALLET_1" -a -e "wallet1.log" ]; then
    echo "balance for WALLET_1"
    blockchain_go getbalance -address $WALLET_1
fi
if [ -n "$WALLET_2" -a -e "wallet2.log" ]; then
    echo "balance for WALLET_2"
    blockchain_go getbalance -address $WALLET_2
fi
if [ -n "$WALLET_3" -a -e "wallet3.log" ]; then
    echo "balance for WALLET_3"
    blockchain_go getbalance -address $WALLET_3
fi
if [ -n "$MINER_WALLET" -a -e "wallet4.log" ]; then
    echo "balance for MINER_WALLET"
    blockchain_go getbalance -address $MINER_WALLET
fi
