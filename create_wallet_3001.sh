export NODE_ID=3001
export WALLET_1=$(blockchain_go createwallet)
export WALLET_2=$(blockchain_go createwallet)
export WALLET_3=$(blockchain_go createwallet)
blockchain_go send -from $CENTREAL_NODE -to $WALLET_1 -amount 10 -mine
blockchain_go send -from $CENTREAL_NODE -to $WALLET_2 -amount 10 -mine
