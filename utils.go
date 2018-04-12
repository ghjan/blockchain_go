package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
)

// IntToHex converts an int64 to a byte array
func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}

// ReverseBytes reverses a byte array
func ReverseBytes(data []byte) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}

func GetPubkeyhashFromAddress(address string) []byte {
	pubKeyHash := Base58Decode([]byte(address))
	pubKeyHash = pubKeyHash[1: len(pubKeyHash)-4]
	return pubKeyHash
}

func getIP() (string, error) {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Println(err)
	}
	ip := ""
	for _, address := range addrs {

		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = ipnet.IP.String()
				break
			}

		}
	}
	return ip, err
}

//获取外部ip地址
func get_external() (string, error) {
	resp, err := http.Get("http://myexternalip.com/raw")
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Stderr.WriteString("\n")
		return "", err
	}
	defer resp.Body.Close()
	//io.Copy(os.Stdout, resp.Body)
	bb, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		return string(bb[:]), nil
	} else {
		return "", err
	}
}

//获取内部ip地址
func get_internal() ([]string, error) {
	var ips []string
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Oops:" + err.Error())
		return ips, err
	}
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ips = append(ips, ipnet.IP.String())
			}
		}
	}
	return ips, err
}

func lookupHostIP(domainName string) ([]string, error) {
	ns, err := net.LookupHost(domainName)
	return ns, err
}
