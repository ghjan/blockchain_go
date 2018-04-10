export NODE_ID=3001
export CENTRAL_NODE=$(cat central_node.log)
blockchain_go getbalance -address $WALLET_1
blockchain_go getbalance -address $WALLET_2
blockchain_go getbalance -address $WALLET_3
blockchain_go getbalance -address $CENTRAL_NODE
export MINER_WALLET=$(cat wallet4.log)
blockchain_go getbalance -address $MINER_WALLET