export NODE_ID=3001
w_results=$(blockchain_go createwallet)
export WALLET_1=${w_results//Your new address: /}
w_results=$(blockchain_go createwallet)
export WALLET_2=${w_results//Your new address: /}
w_results=$(blockchain_go createwallet)
export WALLET_3=${w_results//Your new address: /}
echo $WALLET_1 >wallet1.log
echo $WALLET_2 >wallet2.log
echo $WALLET_3 >wallet3.log
cp blockchain_genesis.db blockchain_3001.db
