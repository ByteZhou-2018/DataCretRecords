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
	CertTimeFormat string // 仅作为格式化展示使用的字段

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


//序列化  ：将数据从内存中的形式转换为一种可以持久化存储在硬盘上或者在网络上传输的形式。称之为序列化
//反序列		：将数据从文件中或网络中读取，然后转化到计算机内存中的过程
//只有进行序列化以后的对象才能进行传输
//序列化和反序列化有很多种方式：
//json: 序列化 ：json.Marshal 反序列化 ：json.UnMarshal
//xml: 序列化 ：xml.Marshal 反序列化 ：xml.UnMarshal

//blockJson,_ := json.Marshal(block0)
//blockXml,_ := xml.Marshal(block0)
//blockGob := block0.Serialize()
//fmt.Println("序列化以后的block",string(blockGob))
//fmt.Println("xml序列化以后的block",string(blockXml))
////blockAsn1,_ :=asn1.Marshal(block0)
//fmt.Println("ans1序列化后的block",string(blockAsn1))

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
	err := decoder.Decode(&block)//将decoder中的通过NewReader读取到数据通过Decode()方法解析成block类型
	if err != nil {
		return nil,err
	}
	return &block,nil
}