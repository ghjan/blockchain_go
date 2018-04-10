package main

import (
	"fmt"
	"log"
	"strings"
)

func (cli *CLI) createBlockchain(address, nodeID string) {
	if strings.Count(address,"")>36 || strings.Contains(address, ":") {
		address = strings.Replace(strings.Split(address, ":")[1], " ", "", -1)
		address = strings.Replace(address, "\n", "", -1)
	}
	if !ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}
	bc := CreateBlockchain(address, nodeID)
	defer bc.db.Close()

	UTXOSet := UTXOSet{bc}
	UTXOSet.Reindex()

	fmt.Println("Done!")
}
