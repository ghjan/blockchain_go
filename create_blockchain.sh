#!/bin/bash
export NODE_ID=3000
check_results=$(blockchain_go createwallet)
echo $check_results
result=$(blockchain_go createblockchain -address "$check_results")
chain_result=$(echo ${result% *})
chain_result=$(echo ${chain_result##* })
chain_result=$(echo ${chain_result##*\r})
export CENTREAL_NODE=$chain_result
cp blockchain_3000.db blockchain_genesis.db
