#!/bin/bash
export NODE_ID=3000
check_results=$(blockchain_go createwallet)
check_results=${check_results//Your new address: /}
blockchain_go createblockchain -address "$check_results" > temp.log
sed -e 's/\r/\n/g;s/Done!//g;' temp.log >temp2.log
tail -n 3 temp2.log >temp3.log
chain_result=$(cat temp3.log |sed '/^$/d')
export CENTRAL_NODE=$chain_result
cp blockchain_3000.db blockchain_genesis.db
