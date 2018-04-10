#!/bin/bash
export NODE_ID=3000
check_results=$(blockchain_go createwallet)
CENTRAL_NODE=${check_results//Your new address: /}
blockchain_go createblockchain -address "$CENTRAL_NODE"
:<<!
> temp.log
sed -e 's/\r/\n/g;s/Done!//g;' temp.log >temp2.log
tail -n 3 temp2.log >temp3.log
chain_result=$(cat temp3.log |sed '/^$/d')
!
cp blockchain_3000.db blockchain_genesis.db
