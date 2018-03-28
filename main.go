package main

import (
	"fmt"
	"strconv"
)

func main() {
	cliRun()
}

func cliRun() {
	//bc := NewBlockchain()
	//defer bc.db.Close()

	cli := CLI{}
	cli.Run()
}
