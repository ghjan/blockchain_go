export CENTRAL_NODE=$(cat central_node.log)
export WALLET_1=$(cat wallet1.log)
export WALLET_2=$(cat wallet2.log)
export WALLET_3=$(cat wallet3.log)
export MINER_WALLET=$(cat wallet4.log)
blockchain_go getbalance -address $CENTRAL_NODE
blockchain_go getbalance -address $WALLET_1
blockchain_go getbalance -address $WALLET_2
blockchain_go getbalance -address $WALLET_3
blockchain_go getbalance -address $MINER_WALLET