export NODE_ID=3002
if [ -z "$MINER_WALLET" ]; then
    export MINER_WALLET=$(cat wallet4.log)
fi
blockchain_go startnode -miner $MINER_WALLET
