#!/bin/bash
export NODE_ID=3000
if [ -n "$CENTRAL_NODE" ]; then
    blockchain_go startnode
fi