package blockchain

import (
	"DataCertPhone/utils"
	"bytes"
	"fmt"
	"github.com/boltdb/bolt"
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
func (bc BlockChain) Each2() ([]*Block, error) {
	db := bc.BoltDB
	var Blocks []*Block




	db.Update(func(tx *bolt.Tx) error {
		blocks := tx.Bucket([]byte(BUCKETBLOCKS))
		if blocks == nil {
			panic("获取桶对象失败！")
		}

		for {
			//ThisBlockBytes := blocks.Get(bc.LastHash)
			ThisBlockHash := blocks.Get(bc.LastHash) //hash值
			//fmt.Printf("区块的hash值：%x\n", ThisBlockHash)
			ThisBlockBytes := blocks.Get(ThisBlockHash) //区块的[]byte
			fmt.Println(len(ThisBlockBytes))
			ThisBlock, err := Deserialize(ThisBlockBytes)
			if err != nil {
				fmt.Println(err.Error())
				break
			}
			Blocks = append(Blocks, ThisBlock)
			//bc.LastHash = ThisBlock.PrevHash
			//fmt.Println(ThisBlock == nil)
			if bytes.Compare(ThisBlock.PrevHash, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}) == 0 {
				//ThisBlockBytes := blocks.Get(bc.LastHash)
				//ThisBlock, _ := Deserialize(ThisBlockBytes)
				//Blocks = append(Blocks, ThisBlock)
				break
			}
			//把拿出来的区块 中的 prevHash ，作为key重新去查 上一个区块的[]byte
			bc.LastHash = ThisBlock.PrevHash
		}
		return nil
	})
	return Blocks, nil
}