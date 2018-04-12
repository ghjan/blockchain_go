#!/bin/bash
export NODE_ID=3000
if [ -z "$CENTRAL_NODE" ]; then
    if [ -e "central_node.log" -a -e "blockchain_3000.db" ]; then
        export CENTRAL_NODE=$(cat central_node.log)
        echo "central_node.log and blockchain_3000.db are already existed, just to restore CENTRAL_NODE:" $CENTRAL_NODE
    else:
         echo "blockchain is not created! Please run ./create_blockchain.sh firstly!"
         exit 1
     fi
fi

if [ -n "$CENTRAL_NODE"  -a -e "blockchain_3000.db" ]; then
    echo "to startnode(main)"
    blockchain_go startnode
else
     echo "blockchain is not created! Please run ./create_blockchain.sh firstly!"
     exit 2
fi
