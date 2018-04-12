#!/bin/bash
rm -rf *.db
rm -rf *.dat
rm -rf *.log

wallets=$(env |grep WALL |awk 'BEGIN {FS="="} {print $1}')
for w in $wallets
    do
	echo $w
        unset $w
    done

nodes=$(env |grep NODE |awk 'BEGIN {FS="="} {print $1}')
for node in $nodes
    do
	echo $node
        unset $node
    done
