#!/bin/bash
export NODE_ID=3000
if [ -z "$CENTRAL_NODE"  -a ! -e "blockchain_3000.db" ]; then
    check_results=$(blockchain_go createwallet)
    export CENTRAL_NODE=${check_results//Your new address: /}
    echo $CENTRAL_NODE >central_node.log
    blockchain_go createblockchain -address "$CENTRAL_NODE"
    cp blockchain_3000.db blockchain_genesis.db
else
     echo "blockchain is already existed! CENTRAL_NODE is:" $CENTRAL_NODE
fi
