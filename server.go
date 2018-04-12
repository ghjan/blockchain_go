package main

import (
	"bytes"
	"encoding/gob"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"strings"
	"time"
)

const PROTOCOL = "tcp"
const NODE_VERSION = 1
const COMMAND_LENGTH = 12
const THRESHOLD = 2
const RETRY = 3
const INVERVAL_RETRY = 3

//#NODE_ADDRESS_BIND和nodeAddressExternal可能不一样，前者仅仅用于tcp socket绑定，后者用于外部通信地址（比如运行在docker container中时）
var NODE_ADDRESS_BIND, nodeAddressExternal string
var miningAddress string

const CENTRAL_NODE = "www.davidzhang.xin:3000"

var beMainNode = false
var CENTRAL_NODE_IP = ""
var knownNodes = []string{CENTRAL_NODE}
var blocksInTransit = [][]byte{}
var mempool = make(map[string]Transaction)
var ipSelf = ""

type addr struct {
	AddrList []string
}

type block struct {
	AddrFrom string
	Block    []byte
}

type getblocks struct {
	AddrFrom string
}

type getdata struct {
	AddrFrom string
	Type     string
	ID       []byte
}

type inv struct {
	AddrFrom string
	Type     string
	Items    [][]byte
}

type tx struct {
	AddFrom     string
	Transaction []byte
}

type verzion struct {
	Version    int
	BestHeight int
	AddrFrom   string
}

func commandToBytes(command string) []byte {
	var bytes [COMMAND_LENGTH]byte

	for i, c := range command {
		bytes[i] = byte(c)
	}

	return bytes[:]
}

func bytesToCommand(bytes []byte) string {
	var command []byte

	for _, b := range bytes {
		if b != 0x0 {
			command = append(command, b)
		}
	}

	return fmt.Sprintf("%s", command)
}

func extractCommand(request []byte) []byte {
	return request[:COMMAND_LENGTH]
}

func requestBlocks() {
	for _, node := range knownNodes {
		sendGetBlocks(node, nodeAddressExternal)
	}
}

func sendAddr(address string, nodeAddress string) {
	nodes := addr{knownNodes}
	nodes.AddrList = append(nodes.AddrList, nodeAddress)
	payload := gobEncode(nodes)
	request := append(commandToBytes("addr"), payload...)

	sendData(address, request)
}

func sendBlock(addr string, b *Block, nodeAddress string) {
	data := block{nodeAddress, b.Serialize()}
	payload := gobEncode(data)
	request := append(commandToBytes("block"), payload...)

	sendData(addr, request)
}

func sendData(addr string, data []byte) error {
	conn, err := net.Dial(PROTOCOL, addr)
	if err != nil {

		fmt.Printf("sendData fail, %s is not available.\n", addr)
		var updatedNodes []string

		for _, node := range knownNodes {
			if node != addr {
				updatedNodes = append(updatedNodes, node)
			}
		}

		knownNodes = updatedNodes

		return err
	}
	defer conn.Close()

	_, err = io.Copy(conn, bytes.NewReader(data))
	if err != nil {
		log.Panic(err)
	}
	return err
}

func sendInv(address, kind string, items [][]byte, nodeAddress string) {
	inventory := inv{nodeAddress, kind, items}
	payload := gobEncode(inventory)
	fmt.Println("to sendInv")
	request := append(commandToBytes("inv"), payload...)

	err := error(nil)
	for i := 0; i <= RETRY; i++ {
		err = sendData(address, request)
		if err == nil {
			break
		} else {
			time.Sleep(time.Duration(INVERVAL_RETRY*(1+i)) * time.Second)
		}
	}
	if err != nil {
		fmt.Println("sendInv fail!!!")
	}
}

func sendGetBlocks(address string, nodeAddress string) {
	payload := gobEncode(getblocks{nodeAddress})
	request := append(commandToBytes("getblocks"), payload...)

	sendData(address, request)
}

func sendGetData(address, kind string, id []byte, nodeAddress string) {
	payload := gobEncode(getdata{nodeAddress, kind, id})
	request := append(commandToBytes("getdata"), payload...)

	sendData(address, request)
}

func sendTx(addr string, tnx *Transaction, nodeAddress string) {
	data := tx{nodeAddress, tnx.Serialize()}
	payload := gobEncode(data)
	request := append(commandToBytes("tx"), payload...)

	sendData(addr, request)
}

func sendVersion(addr string, bc *Blockchain, nodeAddress string) {
	bestHeight := bc.GetBestHeight()
	payload := gobEncode(verzion{NODE_VERSION, bestHeight, nodeAddress})

	request := append(commandToBytes("version"), payload...)

	sendData(addr, request)
}

func handleAddr(request []byte) {
	var buff bytes.Buffer
	var payload addr

	buff.Write(request[COMMAND_LENGTH:])
	dec := gob.NewDecoder(&buff)
	err := dec.Decode(&payload)
	if err != nil {
		log.Panic(err)
	}

	knownNodes = append(knownNodes, payload.AddrList...)
	fmt.Printf("There are %d known nodes now!\n", len(knownNodes))
	requestBlocks()
}

func handleBlock(request []byte, bc *Blockchain) {
	var buff bytes.Buffer
	var payload block

	buff.Write(request[COMMAND_LENGTH:])
	dec := gob.NewDecoder(&buff)
	err := dec.Decode(&payload)
	if err != nil {
		log.Panic(err)
	}

	blockData := payload.Block
	block := DeserializeBlock(blockData)

	fmt.Println("Recevied a new block!")
	bc.AddBlock(block)

	fmt.Printf("Added block %x\n", block.Hash)

	if len(blocksInTransit) > 0 {
		blockHash := blocksInTransit[0]
		sendGetData(payload.AddrFrom, "block", blockHash, nodeAddressExternal)

		blocksInTransit = blocksInTransit[1:]
	} else {
		UTXOSet := UTXOSet{bc}
		UTXOSet.Reindex()
	}
}

func handleInv(request []byte, bc *Blockchain) {
	var buff bytes.Buffer
	var payload inv

	buff.Write(request[COMMAND_LENGTH:])
	dec := gob.NewDecoder(&buff)
	err := dec.Decode(&payload)
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("Recevied inventory with %d %s\n", len(payload.Items), payload.Type)

	if payload.Type == "block" {
		blocksInTransit = payload.Items

		blockHash := payload.Items[0]
		sendGetData(payload.AddrFrom, "block", blockHash, nodeAddressExternal)

		newInTransit := [][]byte{}
		for _, b := range blocksInTransit {
			if bytes.Compare(b, blockHash) != 0 {
				newInTransit = append(newInTransit, b)
			}
		}
		blocksInTransit = newInTransit
	}

	if payload.Type == "tx" {
		txID := payload.Items[0]

		if mempool[hex.EncodeToString(txID)].ID == nil {
			sendGetData(payload.AddrFrom, "tx", txID, nodeAddressExternal)
		}
	}
}

func handleGetBlocks(request []byte, bc *Blockchain) {
	var buff bytes.Buffer
	var payload getblocks

	buff.Write(request[COMMAND_LENGTH:])
	dec := gob.NewDecoder(&buff)
	err := dec.Decode(&payload)
	if err != nil {
		log.Panic(err)
	}

	blocks := bc.GetBlockHashes()
	sendInv(payload.AddrFrom, "block", blocks, nodeAddressExternal)
}

func handleGetData(request []byte, bc *Blockchain) {
	var buff bytes.Buffer
	var payload getdata

	buff.Write(request[COMMAND_LENGTH:])
	dec := gob.NewDecoder(&buff)
	err := dec.Decode(&payload)
	if err != nil {
		log.Panic(err)
	}

	if payload.Type == "block" {
		block, err := bc.GetBlock([]byte(payload.ID))
		if err != nil {
			return
		}

		sendBlock(payload.AddrFrom, &block, nodeAddressExternal)
	}

	if payload.Type == "tx" {
		txID := hex.EncodeToString(payload.ID)
		tx := mempool[txID]

		sendTx(payload.AddrFrom, &tx, nodeAddressExternal)
		// delete(mempool, txID)
	}
}

func handleVersion(request []byte, bc *Blockchain) {
	var buff bytes.Buffer
	var payload verzion

	buff.Write(request[COMMAND_LENGTH:])
	dec := gob.NewDecoder(&buff)
	err := dec.Decode(&payload)
	if err != nil {
		log.Panic(err)
	}

	myBestHeight := bc.GetBestHeight()
	foreignerBestHeight := payload.BestHeight

	if myBestHeight < foreignerBestHeight {
		fmt.Printf("handleVersion, sendGetBlocks to:%s\n", payload.AddrFrom)
		sendGetBlocks(payload.AddrFrom, nodeAddressExternal)
	} else if myBestHeight > foreignerBestHeight {
		fmt.Printf("handleVersion, sendVersion to:%s\n", payload.AddrFrom)
		sendVersion(payload.AddrFrom, bc, nodeAddressExternal)
	}

	// sendAddr(payload.AddrFrom)
	if !nodeIsKnown(payload.AddrFrom) {
		knownNodes = append(knownNodes, payload.AddrFrom)
	}
}

func handleTx(request []byte, bc *Blockchain) {
	var buff bytes.Buffer
	var payload tx

	buff.Write(request[COMMAND_LENGTH:])
	dec := gob.NewDecoder(&buff)
	err := dec.Decode(&payload)
	if err != nil {
		log.Panic(err)
	}

	txData := payload.Transaction
	tx := DeserializeTransaction(txData)
	if bc.VerifyTransaction(&tx) {
		mempool[hex.EncodeToString(tx.ID)] = tx
	} else {
		fmt.Println("Received transaction is invalid! Waiting for new ones..")
	}

	if beMainNode {
		for _, node := range knownNodes {
			if node != CENTRAL_NODE && node != nodeAddressExternal && node != payload.AddFrom {
				sendInv(node, "tx", [][]byte{tx.ID}, nodeAddressExternal)
			}
		}
	} else {
		if len(mempool) >= THRESHOLD && len(miningAddress) > 0 {
		MineTransactions:
			var txs []*Transaction

			for id := range mempool {
				tx := mempool[id]
				if bc.VerifyTransaction(&tx) {
					txs = append(txs, &tx)
				}
			}

			if len(txs) == 0 {
				fmt.Println("All transactions are invalid! Waiting for new ones...")
				return
			}

			cbTx := NewCoinbaseTX(miningAddress, "")
			txs = append(txs, cbTx)

			newBlock := bc.MineBlock(txs)
			UTXOSet := UTXOSet{bc}
			UTXOSet.Reindex()

			fmt.Println("New block is mined!")

			for _, tx := range txs {
				txID := hex.EncodeToString(tx.ID)
				delete(mempool, txID)
			}

			for _, node := range knownNodes {
				if node != nodeAddressExternal {
					sendInv(node, "block", [][]byte{newBlock.Hash}, nodeAddressExternal)
				}
			}

			if len(mempool) > 0 {
				goto MineTransactions
			}
		} else if len(miningAddress) > 0 {
			fmt.Printf("len(mempool):%d, less than THRESHOLD:%d, so Waiting for new ones...", len(mempool), THRESHOLD)
		}
	}
}
func isMainNode(nodeAddress string) bool {
	if CENTRAL_NODE_IP == "" {
		lookupMainNodeIP()
	}
	if CENTRAL_NODE_IP != "" {
		port := strings.Split(CENTRAL_NODE, ":")[1]
		return nodeAddress == CENTRAL_NODE_IP+":"+port
	} else {
		fmt.Println("can not process ip+port for CENTRAOL_NODE, so use nodeAddress == knownNodes[0] instead!")
		return nodeAddress == knownNodes[0]
	}
}

func lookupMainNodeIP() (string, error) {
	ips, err := lookupHostIP(strings.Split(CENTRAL_NODE, ":")[0])
	if err != nil {
		log.Panic(err)
	}
	if len(ips) > 0 {
		CENTRAL_NODE_IP = ips[0]
	} else {
		fmt.Println("can not find an ip address for the CENTRAL_NODE")
	}
	return CENTRAL_NODE_IP, err
}

func handleConnection(conn net.Conn, bc *Blockchain) {
	request, err := ioutil.ReadAll(conn)
	if err != nil {
		log.Panic(err)
	}
	command := bytesToCommand(request[:COMMAND_LENGTH])
	fmt.Printf("Received %s command\n", command)

	switch command {
	case "addr":
		handleAddr(request)
	case "block":
		handleBlock(request, bc)
	case "inv":
		handleInv(request, bc)
	case "getblocks":
		handleGetBlocks(request, bc)
	case "getdata":
		handleGetData(request, bc)
	case "tx":
		handleTx(request, bc)
	case "version":
		handleVersion(request, bc)
	default:
		fmt.Println("Unknown command!")
	}

	conn.Close()
}

// StartServer starts a node
func StartServer(nodeID, minerAddress string) {
	ln, err := listenMe(nodeID, minerAddress)
	if err != nil {
		log.Panic(err)
	}
	defer ln.Close()
	acceptLoop(nodeID, ln)
}

func listenMe(nodeID, minerAddress string) (net.Listener, error) {
	var ips []string
	if ipSelf == "" {
		ipEx, err := get_external()
		if err != nil {
			fmt.Println("can not get correct external ip address, so use interval address instead!")
			//fmt.Println(err)
			log.Panic(err)
		}
		ipInternal, err := getIP()
		if err != nil {
			log.Panic(err)
		}
		ips = append(ips, ipInternal)
		ips = append(ips, ipEx)
		nodeAddressExternal = fmt.Sprintf(ipEx+":%s", nodeID)
		beMainNode = isMainNode(nodeAddressExternal)
	}

	for _, ip := range ips {
		nodeAddress := fmt.Sprintf(ip+":%s", nodeID)
		miningAddress = minerAddress
		ln, err := net.Listen(PROTOCOL, nodeAddress)
		if err == nil {
			NODE_ADDRESS_BIND = nodeAddress
			ipSelf = ip
			return ln, err
		}
	}
	return net.Listener(nil), errors.New("fail")
}

func acceptLoop(nodeID string, ln net.Listener) {
	bc := NewBlockchain(nodeID)
	if !beMainNode {
		fmt.Printf("sendVersion from %s to %s\n", nodeAddressExternal, knownNodes[0])
		sendVersion(knownNodes[0], bc, nodeAddressExternal)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Panic(err)
		}
		go handleConnection(conn, bc)
	}
}

func gobEncode(data interface{}) []byte {
	var buff bytes.Buffer

	enc := gob.NewEncoder(&buff)
	err := enc.Encode(data)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}

func nodeIsKnown(addr string) bool {
	for _, node := range knownNodes {
		if node == addr {
			return true
		}
	}

	return false
}

//TODO addr message
// Provide information on known nodes of the network. Non-advertised nodes should be forgotten after typically 3 hours
