#!/bin/bash
export NODE_ID=3000
check_results=$(blockchain_go createwallet)
echo $check_results
blockchain_go createblockchain -address "$check_results" > temp.txt
sed 's/Done!//g' temp.txt >temp2.txt
sed 's/ //g' temp2.txt >temp3.txt
sed 's/\r/,/g' temp3.txt >temp4.txt

chain_result=${result//Done/}
chain_result=${chain_result// /}
chain_result=${result//\\r/,}
chain_result=${arr[${#chain_result[@]}-1]}
export CENTREAL_NODE=$chain_result
cp blockchain_3000.db blockchain_genesis.db
tail
