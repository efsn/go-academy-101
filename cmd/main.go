package main

import (
	"github.com/efsn/go-workshop-101/internal/chain"
	"github.com/efsn/go-workshop-101/internal/rpc"
)

func main() {
	//handler.Route()
	//log.Println("listener: started: listen on: 4000")
	//err := http.ListenAndServe(":4000", nil)
	//if err != nil {
	//	log.Fatalln(err)
	//}

	//blockChain := chain.NewBlockChain()
	//blockChain.SendData("send 1 btc to oho")
	//blockChain.SendData("send 2 btc to oho")
	//blockChain.SendData("send 3 btc to oho")
	//blockChain.Println()

	rpc.Start(chain.NewBlockChain())
}
