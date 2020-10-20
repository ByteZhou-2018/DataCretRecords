package blockchain

import (
	"DataCertPhone/utils"
	"bytes"
	"time"
)

//定义区块结构体  用于表示区块。。。。。。
type Block struct {
	Height    int64  //表示区块的高度
	TimeStamp int64  //时间戳
	PrevHash  []byte //上一个区块的hash值
	Data      []byte //区块存储的字段
	Hash      []byte //当前区块的hash值
	Version   string //版本号
}

var BlockChain []Block

//创建一个新的区块
func NewBlock(height int64, preHash []byte, data []byte)(*Block,error) {
	block := Block{
		Height:    height,
		TimeStamp: time.Now().Unix(),
		PrevHash:  preHash,
		Data:      data,
		Version:   "0x01",
	}
	blockBytes,err := BlockToByte(block)
	if err != nil {
		return nil, err
	}
	block.Hash, err = utils.SHA256HashByte(blockBytes)
	if err != nil {
		return nil, err
	}
	return &block,nil
}

//创建创世区块
func CreatGenesisBlock()( *Block,error) {
	newBlock,err :=  NewBlock(0, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, nil)
	if err != nil {
		return nil,err
	}
	return newBlock,nil
}



//对区块数据进行类型转换 并返回一个总的字节切片
func BlockToByte(block Block)([]byte,error){
	//1、将block结构体数据转换为[]byte类型
	heightBytes,err  := utils.Int64ToByte(block.Height)
	if err != nil {
		return nil,err
	}
	timeStampBytes,err :=utils.Int64ToByte(block.TimeStamp)
	if err != nil {
		return nil,err
	}
	versionBytes := utils.StringToBytes(block.Version)

	var blockBytes []byte
	//bytes.Join  bytes包下的join函数 拼接字节切片
	blockBytes  =  bytes.Join([][]byte{
		heightBytes,
		timeStampBytes,
		block.PrevHash,
		block.Data,
		versionBytes,
	},[]byte{})
	return blockBytes,nil
}





















//1、哈希值

//2、挖矿竞争新区块的创造权
/*
区块链 的裢 是否是通过数组切片来实现的？每一次添加都相当于一个调用类似于 append方法
第二个区块的值我们应该放哪些值？
*/

//func WaKaung(thisName string) {
//	go func() { //在零到 一千中去找一个数
//		rand.Seed(time.Now().UnixNano())
//		for {
//			numberRand := rand.Intn(10000)
//			if numberRand%10 == 0 && numberRand%5 == 0 {
//				fmt.Println(thisName, "胜出了")
//				var a chan bool
//				
//				break
//			}
//		}
//	}()
//
//}

func AddBlock(i int64, preHash []byte, data []byte, ) error{ //一条链上添加一个新的区块
	i = int64(len(BlockChain) - 1)
	newBlock,err := NewBlock(BlockChain[i].Height+1, BlockChain[i].Hash, data)
	if err != nil {
		return err
	}
	BlockChain = append(BlockChain, *newBlock)
	return nil
}
