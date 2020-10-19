package utils

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"io/ioutil"
)
//对字符串 data 进行 MD5 哈希 返回 data 的md5哈希值
func MD5HashString(data string)string  {
	hashMd5 := md5.New()
	hashMd5.Write([]byte(data))
	bytes:=	hashMd5.Sum(nil)
	return hex.EncodeToString(bytes)
}

func MD5HashReader(reader io.Reader) (string,error)  {
	md5Hash := md5.New()
	readerBytes,err := ioutil.ReadAll(reader)
	if err != nil{
		return "",err
	}
	md5Hash.Write(readerBytes)
	bytes := md5Hash.Sum(nil)
	return hex.EncodeToString(bytes),nil
}
//读取 io 流中的数据，并对数据进行哈希计算，返回sha256哈希值和error
func SHA256HashReader(reader io.Reader) (string,error) {
	hashSha256 := sha256.New()
	readerBytes,err := ioutil.ReadAll(reader)
	if err != nil {
		return "",err
	}
	hashSha256.Write(readerBytes)
	bytes := hashSha256.Sum(nil)
	return  hex.EncodeToString(bytes),err
}