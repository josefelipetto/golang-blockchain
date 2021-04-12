package blockchain

type BlockChain struct {
	Blocks []*Block
}

type Block struct {
	Hash []byte
	Data []byte
	PrevHash []byte
	Nonce int
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash, 0}
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}



func (c *BlockChain) AddBlock(data string)  {
	prevBlock := c.Blocks[len(c.Blocks) - 1]
	newBlock := CreateBlock(data, prevBlock.Hash)
	c.Blocks = append(c.Blocks, newBlock)
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func InitBlockChain()  *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}
