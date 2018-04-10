export NODE_ID=3002
export MINER_WALLET=$(cat wallet4.log)
blockchain_go startnode -miner $MINER_WALLET
