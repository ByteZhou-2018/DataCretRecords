package blockchain

import "time"

//定义区块结构体  用于表示区块。。。。。。
type Block struct {
	Height int64  //表示区块的高度
	TimeStamp int64 //时间戳
	PrevHash []byte //上一个区块的hash值
	Data []byte //区块存储的字段
	Hash []byte //当前区块的hash值
	Version string //版本号
}
//创建一个新的区块
func NewBlock(height int64,preHash []byte,data []byte)Block{
	block := Block{
		Height:    height,
		TimeStamp: time.Now().Unix(),
		PrevHash:  preHash,
		Data:      data,
		Version:   "0x01",
	}
	//block.Hash =
	return block
}
//创建创世区块
func CreatGenesisBlock()Block  {
	return NewBlock(0,[]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0},nil)
}
//1、哈希值
//2、挖矿竞争新区块的创造权
