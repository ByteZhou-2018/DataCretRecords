package blockchain

import (
	"DataCertPhone/utils"
	"bytes"
	"encoding/gob"
	"time"
)

type Block struct {
	Height    int64
	TimeStamp int64
	PrevHash  []byte
	Data      []byte
	Hash      []byte
	Nonce     int64
	Version   string
}

func NewBlock(height int64, prevHash, data []byte) Block {
	block := Block{
		Height:    height,
		TimeStamp: time.Now().Unix(),
		PrevHash:  prevHash,
		Data:      data,
		Version:   "0 x 01",
	}
	pow := NewPoW(block)

	block.Hash, block.Nonce = pow.Run()
	return block

}

func CreatGenesisBlock() Block {
	newBlock := NewBlock(0, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, nil)

	return newBlock
}



func BlockToByte(block Block) ([]byte, error) {
	//1、将block结构体数据转换为[]byte类型
	heightBytes, err := utils.Int64ToByte(block.Height)
	if err != nil {
		return nil, err
	}
	timeStampBytes, err := utils.Int64ToByte(block.TimeStamp)
	if err != nil {
		return nil, err
	}
	versionBytes := utils.StringToBytes(block.Version)

	var blockBytes []byte
	//bytes.Join  bytes包下的join函数 拼接字节切片
	blockBytes = bytes.Join([][]byte{
		heightBytes,
		timeStampBytes,
		block.PrevHash,
		block.Data,
		versionBytes,
	}, []byte{})
	//挖矿竞争，获得记账权
	return blockBytes, nil
}
//对区块进行序列化操作。。。
func (b Block) Serialize()([]byte) {
	buff := new(bytes.Buffer) //缓冲区

	encoder := gob.NewEncoder(buff)
	encoder.Encode(b)
	return buff.Bytes()

}
//对区块进行反序列化操作。。。
func Deserialize(data []byte) (*Block,error) {
	//buff := new(bytes.Buffer) //缓冲区
	//
	//encoder := gob.NewDecoder(buff)
	//block1 := encoder.Decode(data)
	//return
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(data))

	err := decoder.Decode(block)//将decoder中的通过NewReader读取到数据通过Decode()方法解析成block类型
	if err != nil {
		return nil,err
	}
	return &block,nil
}