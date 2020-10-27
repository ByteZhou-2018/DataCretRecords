package blockchain

import (
	"DataCertPhone/utils"
	"bytes"
	"fmt"
	"math/big"
)

type ProofOfWork struct {
	Target *big.Int
	Block  Block
}

var Difficulty  uint= 10
func NewPoW(block Block) ProofOfWork {
	t := big.NewInt(1)
	t.Lsh(t, 255-Difficulty)
	pow := ProofOfWork{
		Target: t,
		Block:  block,
	}
	return pow
}
func (p ProofOfWork) Run() ([]byte, int64) {
	var nonce int64 = 0
	BlockBytes, _ := BlockToByte(p.Block)
	var blockHash []byte
	for {
		fmt.Println("当前尝试的nonce值", nonce)
		nonceBytes, _ := utils.Int64ToByte(nonce)
		blockBytes := bytes.Join([][]byte{
			BlockBytes, nonceBytes,
		}, []byte{})
		blockHash, _ = utils.SHA256HashByte(blockBytes)
		target := p.Target
		hashBig := new(big.Int)
		hashBig = hashBig.SetBytes(blockHash)
		if hashBig.Cmp(target) == -1 {
			//找到了这个值
			break
		}
		nonce++
	}
	return blockHash, nonce
}
