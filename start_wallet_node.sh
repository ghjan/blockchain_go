export NODE_ID=3001
if [ ! -e "blockchain_genesis.db" ]; then
    echo "blockchain_genesis.db is not existed, please run create_blockchain.sh firstly!"
    exit 1
elif [ ! -e "blockchain_3001.db" ]; then
    echo "blockchain_3001.db is not existed, copy from blockchain_genesis.db firstly"
    cp blockchain_genesis.db blockchain_3001.db
fi
echo "to startnode (wallet)"
blockchain_go startnode
