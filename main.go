package main

func main() {
	cliRun()
}

func cliRun() {
	//bc := NewBlockchain()
	//defer bc.db.Close()

	cli := CLI{}
	cli.Run()
}
