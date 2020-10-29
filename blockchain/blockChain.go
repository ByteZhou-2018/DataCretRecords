package blockchain

import (
	"errors"
	"github.com/boltdb/bolt"
	"math/big"
)

const BLOCKCHAINDB = "blockchain.db"
const LASTHASH = "lastHash"
const BUCKET_NAME = "blocks"

var CHAIN *BlockChain

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
	//打开本地保存的区块链
	db, err := bolt.Open(BLOCKCHAINDB, 0600, nil)
	if err != nil {
		panic(err.Error())
	}
	bc := BlockChain{
		BoltDB: db,
	}
	//2、判断本地是否已经存在了这条区块链（通过判断 blocks桶是否为空）
	db.Update(func(tx *bolt.Tx) error {
		backut := tx.Bucket([]byte(BUCKET_NAME))
		if backut == nil { //没有桶，创建桶
			genesis := CreatGenesisBlock() //创建创世区块
			backut, err = tx.CreateBucket([]byte(BUCKET_NAME))
			if err != nil {
				panic(err.Error())
			}

			blocksBucket := tx.Bucket([]byte(BUCKET_NAME))
			if blocksBucket == nil {
				return err
			}
			blocksBucket.Put(genesis.Hash, genesis.Serialize()) //新建区块的key value
			blocksBucket.Put([]byte(LASTHASH), genesis.Hash)    //最新区块的 hash值
			bc.LastHash = genesis.Hash

			//bc.LastHash = backut.Get([]byte(LASTHASH))
		} else {
			bc.LastHash = backut.Get([]byte(LASTHASH)) //更新一下区块链的最新区块的哈希值
		}
		return nil
	})
	CHAIN = &bc
	return bc

}
func (bc *BlockChain) SaveData(data []byte) (Block, error) {
	//1、从文件中读取到最新的区块
	db := bc.BoltDB
	var lastBlock *Block
	//error的自定义
	var err error
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAME))
		if bucket == nil {
			err = errors.New("读取区块链数据失败")
			//panic("读取区块链数据失败")
			return err
		}
		//lastHash := bucket.Get([]byte(LAST_HASH))
		lastBlockBytes := bucket.Get(bc.LastHash)
		//反序列化
		lastBlock, _ = Deserialize(lastBlockBytes)
		return nil
	})

	//新建一个区块
	newBlock := NewBlock(lastBlock.Height+1, lastBlock.Hash, data)
	//把新区块存到文件中
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAME))
		//序列化后的区块数据
		blockBytes := newBlock.Serialize()
		//fmt.Println("保存数据到区块，序列化后的区块数据：", blockBytes)
		//把新创建的区块存入到boltdb数据库中
		//fmt.Printf("保存数据到区块，区块的hash值是:%x\n", newBlock.Hash)
		bucket.Put(newBlock.Hash, blockBytes)
		//更新LASTHASH对应的值，更新为最新存储的区块的hash值
		bucket.Put([]byte(LASTHASH), newBlock.Hash)
		bc.LastHash = newBlock.Hash //将区块链实例的LASTHASH值更新为最新区块的HASH
		return nil
	})
	//返回值语句，newBlock，err，其中err可能包含错误信息
	return newBlock, err
}
//func (bc *BlockChain) SaveData(data []byte) (Block, error) { //用户传入需要保存的数据 data
//	//1、读取最新的区块
//	db := bc.BoltDB
//	var err error
//	var lastBlock *Block
//	db.View(func(tx *bolt.Tx) error {
//		blocks := tx.Bucket([]byte(BUCKET_NAME))
//		lastBlockBytes := blocks.Get(bc.LastHash)
//		lastBlock, _ = Deserialize(lastBlockBytes)
//		return nil
//	})
//	//新建一个区块
//	newBlock := NewBlock(lastBlock.Height+1, lastBlock.PrevHash, data)
//	//打开本地区块的数据库
//
//	db.Update(func(tx *bolt.Tx) error {
//		blocksBucket := tx.Bucket([]byte(BUCKET_NAME))
//		if blocksBucket == nil {
//			return err
//		}
//		blocksBucket.Put(newBlock.Hash, newBlock.Serialize()) //新建区块的key value
//		blocksBucket.Put([]byte(LASTHASH), newBlock.Hash)     //最新区块的 hash值
//		bc.LastHash = lastBlock.Hash
//		return err
//	})
//	//把区块链结构体的那个LasHash更新一下
//	return newBlock, err
//}
func (bc BlockChain) QueryBlockByHeight(height int64) (*Block, error) {
	var err error
	if height < 0 {
		err = errors.New("您输入高度小于零，查询失败。。。")
		return nil, err
	}
	db := bc.BoltDB
	var eachBlock *Block
	err = db.View(func(tx *bolt.Tx) error {
		blocks := tx.Bucket([]byte(BUCKET_NAME))
		eachHash := bc.LastHash
		for {
			blockBytes := blocks.Get(eachHash)
			eachBlock, _ = Deserialize(blockBytes)
			if eachBlock.Height < height {
				err = errors.New("您输入高度大于区块链的高度，查询失败。。。")
				return err
			}
			if eachBlock.Height == height {
				break
			}
			eachHash = eachBlock.PrevHash
		}
		return err
	})
	return eachBlock, err
}

//该方法用于遍历区块链  查出所有区块，并返回
func (bc BlockChain) QueryAllBlocks() ([]*Block, error) {
	var blocks = make([]*Block, 0) //blocks 是一个容器，用于存放查询到的区块
	var err error
	db := bc.BoltDB
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAME))
		if bucket == nil {
			err = errors.New("查询区块链数据失败!")
			return err
		}
		//bucket存在
		eachHash := bc.LastHash
		eachBig := new(big.Int)
		zeroBig := big.NewInt(0) //默认值零的大整数
		for {
			//根据区块的hash值获取对应的区块
			eachBlockBytes := bucket.Get(eachHash)
			//反序列化操作
			eachBlock, _ := Deserialize(eachBlockBytes)
			//将遍历到每一个区块放入到切片容器当中
			blocks = append(blocks, eachBlock)

			eachBig.SetBytes(eachBlock.PrevHash)
			if eachBig.Cmp(zeroBig) == 0 { //找到了创世区块
				break //跳出循环
			}
			//不满足条件，没有找到创世区块
			eachHash = eachBlock.PrevHash
		}
		return nil
	})

	return blocks, err
}

//eachHash := bc.LastHash
//eachBig := new(big.Int)
//zero := big.NewInt(0)//默认值零的大整数
//for{
//	eachBlockBytes := bucket.Get(eachHash)
//	eachBlock,err := Deserialize(eachBlockBytes)
//	//遍历到的每一个区块放入到容器中
//	if err !=nil {
//		return err
//	}
//	blocks = append(blocks, eachBlock)
//	eachBig.SetBytes(eachBlock.PrevHash)
//	if eachBig.Cmp(zero) == 0 {
//		break
//	}
//	eachHash = eachBlock.PrevHash
//}
//return err

////查找指定的区块信息返回
//func (bc BlockChain) QueryBlock(HashKey []byte) Block {
//	var block *Block
//	db := bc.BoltDB
//	db.View(func(tx *bolt.Tx) error {
//		blocks := tx.Bucket([]byte(BUCKET_NAME))
//		if blocks == nil {
//			panic("桶是空的")
//		}
//		thisBlockBytes := blocks.Get(HashKey)
//		block, _ = Deserialize(thisBlockBytes)
//		return nil
//	})
//	return *block
//}

//遍历所有区块信息，返回一个区块切片
//func (bc BlockChain) Each() (map[string]Block, error) {
//	db := bc.BoltDB
//	var allBlock = make(map[string]Block, 0)
//	var err error
//	err = db.View(func(tx *bolt.Tx) error {
//		//开启一个事务
//		//tx, err := db.Begin(true)
//		//if err != nil {
//		//	return err
//		//}
//		//defer tx.Rollback()
//		//tx.Commit()
//
//		blocks := tx.Bucket([]byte(BUCKET_NAME))
//		if blocks == nil {
//			err = errors.New("读取数据失败，区块链不存在！")
//			return err
//		}
//		//ForEach遍历
//		//blocks.ForEach(func(k, v []byte) error {
//		//		//	//fmt.Printf("key=%x, value=%v\n", k, v)
//		//		//	//key := hex.EncodeToString(k)
//		//		//	//thisblock,_ := Deserialize(v)
//		//		//	if k != nil {
//		//		//		//key := hex.EncodeToString(k)
//		//		//		//allBlock[key] = *thisblock
//		//		//		fmt.Printf("%x\n", k)
//		//		//	}
//		//		//
//		//		//	return nil
//		//		//})
//
//		//bucket.Cursor方法遍历
//		c := blocks.Cursor()
//		for k, v := c.First(); k != nil; k, v = c.Next() {
//			//fmt.Printf("key=%s, value=%s\n", k, v)
//			thisBlock,_ :=Deserialize(v)
//			if thisBlock != nil {
//
//				key := hex.EncodeToString(k)
//				allBlock[key] = *thisBlock
//			}
//		}
//
//		//eachHash := []byte(LASTHASH)
//		//for {
//		//	//ThisBlockBytes := blocks.Get(bc.LastHash)
//		//	ThisBlockBytes := blocks.Get(eachHash) //hash值
//		//	//fmt.Printf("区块的hash值：%x\n", ThisBlockHash)
//		//	//ThisBlockBytes := blocks.Get(ThisBlockHash) //区块的[]byte
//		//	//fmt.Println(len(ThisBlockBytes))
//		//	ThisBlock, err := Deserialize(ThisBlockBytes)
//		//	if err != nil {
//		//		fmt.Println(err.Error())
//		//		break
//		//	}
//		//	//blocksArray = append(blocksArray, *ThisBlock)
//		//	//bc.LastHash = ThisBlock.PrevHash
//		//	fmt.Println(ThisBlock == nil)
//		//	if bytes.Compare(ThisBlock.PrevHash, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}) == 0 {
//		//		//ThisBlockBytes := blocks.Get(bc.LastHash)
//		//		//ThisBlock, _ := Deserialize(ThisBlockBytes)
//		//		//Blocks = append(Blocks, ThisBlock)
//		//		break
//		//	}
//		//	//把拿出来的区块 中的 prevHash ，作为key重新去查 上一个区块的[]byte
//		//	eachHash = ThisBlock.PrevHash
//		//}
//		return nil
//	})
//	return allBlock, err
//}
//
////func AddGenesisBlock(tx *bolt.Tx, block Block,bc BlockChain) error {
////		blocksBucket, err := tx.CreateBucket([]byte(BUCKETBLOCKS))
////		if err != nil {
////			panic(err.Error())
////		}
////		blocksBucket.Put(block.Hash, block.Serialize()) //新建区块的key value
////		blocksBucket.Put([]byte(LASTHASH), block.Hash) //最新区块的 hash值
////		bc.LastHash = block.Hash//更新这个实例化的区块链的lastHash
////
////
////	return nil
////
////}
//func AddNewBlock(tx *bolt.Tx, block Block, bc *BlockChain) {
//	blocksBucket := tx.Bucket([]byte(BUCKET_NAME))
//	if blocksBucket == nil {
//		panic("桶是空的")
//	}
//	blocksBucket.Put(block.Hash, block.Serialize()) //新建区块的key value
//	blocksBucket.Put([]byte(LASTHASH), block.Hash)  //最新区块的 hash值
//	bc.LastHash = block.Hash
//}

//更新最新区块的哈希值记录
//lastHash,err := tx.CreateBucket([]byte("lastHash"))不需要新建一个桶
//if err != nil {
//	panic(err.Error())
//}
//ThisBlockBytes := blocks.Get(bc.LastHash)
//ThisBlock,_ := Deserialize(ThisBlockBytes)
//Blocks = append(Blocks, ThisBlock)
//LastBlockBytes := blocks.Get(ThisBlock.PrevHash)
//LastBlock,_ := Deserialize(LastBlockBytes)
//
////创世区块
//genesis := CreatGenesisBlock()
////创建区块链保存的文件
//
//bc := BlockChain{
//	LastHash: genesis.Hash,
//	BoltDB:   db,
//}
////把创世区块保存到数据库文件中去
//db.Update(func(tx *bolt.Tx) error {
//	err := AddGenesisBlock(tx, genesis)
//	if err != nil {
//		return err
//	}
//	return nil
//})
//db.Update(func(tx *bolt.Tx) error {
//buckt := tx.Bucket([]byte(BUCKETBLOCKS))//假设有桶
//if buckt == nil {
//		genesis := CreatGenesisBlock()
//
//	AddGenesisBlock(tx,genesis)
//buckt,err := tx.CreateBucket([]byte(BUCKETBLOCKS))
//if err == nil {
//	panic(err.Error())
//}
//lastHash := buckt.Get([]byte(BUCKETBLOCKS))
//if len(lastHash) == 0{//桶中没有lasthash记录，需要创世区块，并保存
//	genesis := CreatGenesisBlock()
//	AddGenesisBlock(tx,genesis)
//	bc = BlockChain{
//		LastHash: genesis.Hash,
//		BoltDB:   db,
//	}
//}else {
//lastHsah1 := buckt.Get([]byte(lastHash))
//bc = BlockChain{
//LastHash: lastHsah1,
//BoltDB:   db,
//}
//}}
//
//return nil
//})
//return bc

//var lastBlock *Block
//db.View(func(tx *bolt.Tx) error {
//	LastHash := tx.Bucket([]byte(BUCKETBLOCKS))
//	lastBlockBytes := LastHash.Get([]byte(LASTHASH))
//
//	//
//	lastBlock, _ = Deserialize(lastBlockBytes)
//
//	return nil
//})
//db, err := bolt.Open(BLOCKCHAINDB, 0600, nil)
//db,err := bolt.Open(BLOCKCHAINDB,0777,nil)
//fmt.Println("emmm......1")
//
//if err != nil {
//	fmt.Println("emmm......2")
//
//	fmt.Println("打开db文件失败")
//	return nil,err
//}
