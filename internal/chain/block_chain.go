package chain

import (
	"fmt"
	"log"
)

type BlockChain struct {
	Blocks []*Block
}

func (c *BlockChain) Println() {
	for _, block := range c.Blocks {
		block.Println()
		fmt.Println()
	}
}

func NewBlockChain() *BlockChain {
	root := GenerateRoot()
	return &BlockChain{
		Blocks: []*Block{&root},
	}
}

func (c *BlockChain) SendData(data string) {
	pre := c.Blocks[len(c.Blocks)-1]
	next := GenerateNext(*pre, data)
	c.Add(&next)
}

func (c *BlockChain) Add(next *Block) {
	if isValid(*next, *c.Blocks[len(c.Blocks)-1]) {
		c.Blocks = append(c.Blocks, next)
	} else {
		log.Fatalln("invalid block")
	}
}

func isValid(newBlock Block, oldBlock Block) bool {
	if newBlock.Index-1 != oldBlock.Index {
		return false
	}
	if newBlock.PreBlockHash != oldBlock.Hash {
		return false
	}
	if calcHash(newBlock) != newBlock.Hash {
		return false
	}
	return true
}
