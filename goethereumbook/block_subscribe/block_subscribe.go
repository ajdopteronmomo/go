package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main()  {
	//订阅新区块
	client,err:=ethclient.Dial("wss://mainnet.infura.io/ws/v3/ff8dd216d182458ab681895ebbdca51f")
	if err !=nil{
		log.Fatal(err)
	}

	headers:=make(chan *types.Header)
	sub,err:=client.SubscribeNewHead(context.Background(),headers)
	if err !=nil{
		log.Fatal(err)
	}

	for {
		select {
		case err:=<-sub.Err():
			log.Fatal(err)
		case header:=<-headers:
			fmt.Println(header.Hash().Hex()) 

			block,err:=client.BlockByHash(context.Background(),header.Hash())
			if err !=nil{
				log.Fatal(err)
			}

			fmt.Println(block.Hash().Hex())
			fmt.Println(block.Number().Uint64())
			fmt.Println(block.Time())
			fmt.Println(block.Nonce())
			fmt.Println(len(block.Transactions()))
		}
	}
}