export NODE_ID=3002
cp blockchain_genesis.db blockchain_3002.db
w_results=$(blockchain_go createwallet)
export MINER_WALLET=${w_results//Your new address: /}
echo $MINER_WALLET >wallet4.log