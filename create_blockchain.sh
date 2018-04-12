#!/bin/bash
export NODE_ID=3000
if [ ! -e "blockchain_3000.db" ]; then
    if [ -z "$CENTRAL_NODE" ]; then
        check_results=$(blockchain_go createwallet)
        export CENTRAL_NODE=${check_results//Your new address: /}
        echo $CENTRAL_NODE >central_node.log
    fi
    blockchain_go createblockchain -address "$CENTRAL_NODE"
    cp blockchain_3000.db blockchain_genesis.db
else
     echo "blockchain is already existed! CENTRAL_NODE is:" $CENTRAL_NODE
     exit 1
fi
