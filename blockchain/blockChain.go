package blockchain

import (
	"bytes"
	"github.com/boltdb/bolt"
)

const BLOCKCHAIN = "blockchain.db"
const LASTHASH = "lastHash"
const BLOCKS = "blocks"

//区块链结构体的定义	,代表的是整个区块链，他需要以下几个功能,方法。
//1、将新区块与已有区块连接
//2、查询某个已有区块的数据和信息
//3、遍历区块信息

type BlockChain struct {
	LastHash []byte //表示区块链中最新区块的哈希，用于查找最新的区块内容

	BoltDB *bolt.DB //区块链中操作区块数据文件的数据库操作对象
}

//实例化一条区块链
func NewBolckChain() BlockChain {
	//创世区块
	genesis := CreatGenesisBlock()
	//创建区块链保存的文件
	db, err := bolt.Open(BLOCKCHAIN, 0600, nil)
	if err != nil {
		panic(err.Error())
	}
	bc := BlockChain{
		LastHash: genesis.Hash,
		BoltDB:   db,
	}
	//把创世区块保存到数据库文件中去
	db.Update(func(tx *bolt.Tx) error {
		err := AddBlock(tx, genesis)
		if err != nil {
			return err
		}
		return nil
	})
	return bc
}
func (bc BlockChain) SaveData(data []byte) { //用户传入需要保存的数据 data
	//1、读取最新的区块
	db := bc.BoltDB
	var lastBlock *Block
	db.View(func(tx *bolt.Tx) error {
		LastHash := tx.Bucket([]byte(BLOCKS))
		lastBlockBytes := LastHash.Get([]byte(LASTHASH))
		//
		lastBlock, _ = Deserialize(lastBlockBytes)

		return nil
	})

	//新建一个区块
	newBlock := NewBlock(lastBlock.Height+1, bc.LastHash, data)
	//打开本地区块的数据库
	db, err := bolt.Open(BLOCKCHAIN, 0600, nil)
	if err != nil {
		panic(err.Error())
	}
	db.Update(func(tx *bolt.Tx) error {
		AddBlock(tx, newBlock)
		return nil
	})
}
func (bc BlockChain) QueryBlock(HashKey []byte) Block {
	var block *Block
	db, err := bolt.Open(BLOCKCHAIN, 0600, nil)
	if err != nil {
		panic(err.Error())
	}
	db.View(func(tx *bolt.Tx) error {
		blocks := tx.Bucket([]byte(BLOCKS))
		if blocks == nil {
			panic(err.Error())
		}
		thisBlockBytes := blocks.Get(HashKey)
		block, _ = Deserialize(thisBlockBytes)
		return nil
	})
	return *block
}

//遍历所有区块信息，返回一个区块切片
func (bc BlockChain) Each() []*Block {
	var Blocks []*Block
	db, err := bolt.Open(BLOCKCHAIN, 0600, nil)
	if err != nil {
		panic(err.Error())
	}
	db.View(func(tx *bolt.Tx) error {

		blocks := tx.Bucket([]byte(BLOCKS))
		if blocks == nil {
			panic(err.Error())
		}
		//ThisBlockBytes := blocks.Get(bc.LastHash)
		//ThisBlock,_ := Deserialize(ThisBlockBytes)
		//Blocks = append(Blocks, ThisBlock)
		//LastBlockBytes := blocks.Get(ThisBlock.PrevHash)
		//LastBlock,_ := Deserialize(LastBlockBytes)
		//
		for {
			ThisBlockBytes := blocks.Get(bc.LastHash)
			ThisBlock, _ := Deserialize(ThisBlockBytes)
			Blocks = append(Blocks, ThisBlock)
			bc.LastHash = ThisBlock.PrevHash
			if bytes.Compare(bc.LastHash, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}) == 0 {
				ThisBlockBytes := blocks.Get(bc.LastHash)
				ThisBlock, _ := Deserialize(ThisBlockBytes)
				Blocks = append(Blocks, ThisBlock)
				break
			}

		}
		return nil
	})
	return Blocks
}

func AddBlock(tx *bolt.Tx, block Block) error {
	blocksBucket, err := tx.CreateBucket([]byte(BLOCKS))
	if err != nil {
		panic(err.Error())
	}
	blocksBucket.Put(block.Hash, block.Serialize()) //新建区块的key value

	blocksBucket.Put([]byte(LASTHASH), block.Hash) //最新区块的 hash值
	return nil
}

//更新最新区块的哈希值记录
//lastHash,err := tx.CreateBucket([]byte("lastHash"))不需要新建一个桶
//if err != nil {
//	panic(err.Error())
//}
