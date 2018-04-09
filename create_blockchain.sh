#!/bin/bash
export NODE_ID=3000
check_results=$(blockchain_go createwallet)
echo $check_results
chain_result=$(blockchain_go createblockchain -address "$check_results")
export CENTREAL_NODE=$(echo ${chain_result% *})
cp blockchain_3000.db blockchain_genesis.db
