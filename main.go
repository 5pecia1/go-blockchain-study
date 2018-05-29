package main

import (
	"fmt"
)

func main() {
	bc := NewBlockchain()

	bc.Addblock("A한테 B가 100BTC 보냄")
	bc.Addblock("C -> D 100BTC")
	bc.Addblock("C -> D 100BTC")

	for _, block := range bc.blocks {
		fmt.Println("Prev ", block.PrevBlockHash)
		fmt.Println("Data ", block.Data)
		fmt.Println("Hash ", block.Hash)
		fmt.Println()
	}
}
