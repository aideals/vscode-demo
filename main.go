package main

import "fmt"

func main() {

	bc := NewBlockChain()

	bc.AddBlock("hello 航头")
	bc.AddBlock("再见，航头")

	for i, block := range bc.Blocks {
		fmt.Printf("=========== 区块高度 : %d ===========\n", i)

		fmt.Printf("前区块哈希值:%x\n", block.PrevBlockHash)
		fmt.Printf("当前区块哈希值:%x\n", block.Hash)
		fmt.Printf("区块数据:%x\n", block.Data)
	}
}
