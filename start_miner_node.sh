export NODE_ID=3002
if [ -z "$MINER_WALLET" ]; then
    if [ -e "wallet4.log" -a -e "blockchain_3002.db" ]; then
        export MINER_WALLET=$(cat wallet4.log)
        echo "wallet4.log and blockchain_3002.db are already existed, just to restore MINER_WALLET:" $MINER_WALLET
     else:
         echo "miner node is not initialized! Please run ./init_miner_node.sh firstly!"
         exit 1
     fi
fi
if [ -n "$MINER_WALLET"  -a -e "blockchain_3002.db" ]; then
    echo "to startnode(miner)"
    blockchain_go startnode -miner $MINER_WALLET
else:
     echo "blockchain_3002.db is not exsited, so miner node is not initialized! Please run ./init_miner_node.sh firstly!"
     exit 2
fi
