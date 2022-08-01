package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

// BlockChain is an array of pointers
//array of blocks to create blockchain
type BlockChain struct {
	blocks []*Block
}

// Block is a single unit in the blockchain
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	// This will join our previous block's relevant info with the new blocks
	hash := sha256.Sum256(info)
	//This performs the actual hashing algorithm
	b.Hash = hash[:]
	//If this ^ doesn't make sense, you can look up slice defaults
}

//outputs a pointer to a block
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}

	block.DeriveHash()
	return block
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, new)
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func InitBlockChain() *BlockChain {
	//return refernce to BlockChain, and inside an array of Bl;ocks with a call to Genesis funcntion
	return &BlockChain{[]*Block{Genesis()}}
}

func main() {
	chain := InitBlockChain()

	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second Block after Genesis")
	chain.AddBlock("Third Block after Genesis")

	for _, block := range chain.blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
	}
}
