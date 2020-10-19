package Hash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/md4"
	"golang.org/x/crypto/ripemd160"
	"hash"
)


func MD4(text string) string {
	//var hashInstance hash.Hash
	hashInstance := md4.New()
	arr, _ := hex.DecodeString(text)
	hashInstance.Write(arr)
	bytes := hashInstance.Sum(nil)
	return fmt.Sprintf("%x", bytes)
}
func HASH(text string, hexType string, isHex bool) string {
	var hashInstance hash.Hash
	//通过传入的hexType值  来switch一下自己要hash的方法。
	switch hexType {
	case "md4":
		hashInstance = md4.New()
	case "md5":
		hashInstance = md5.New()
	case "sha1":
		hashInstance = sha1.New()
	case "sha256":
		hashInstance = sha256.New()
	case "sha512":
		hashInstance = sha512.New()
	case "ripemd160":
		hashInstance = ripemd160.New()
	}
	if isHex { //如果传入的text 为十六进制，先hex。DecodeString()解码为arry再 write
		arr, _ := hex.DecodeString(text)
		hashInstance.Write(arr)

	} else { //如果传入的text 为十进制，直接用[]byte()强转array write
		hashInstance.Write([]byte(text))
	}
	bytes := hashInstance.Sum(nil)
	return fmt.Sprintf("%x\n", bytes)
}
func SHA256Double(text string, isHex bool) []byte {
	hashInstance := sha1.New()
	if isHex {
		arr, _ := HexStringToBytes(text)
		hashInstance.Write(arr)
	} else {
		hashInstance.Write([]byte(text))
	}
	bytes := hashInstance.Sum(nil)
	//不用再 new object ,使用Reset()重置原本已经定义过的object.节省内存,提升性能.
	hashInstance.Reset()
	hashInstance.Write(bytes)
	bytes = hashInstance.Sum(nil)
	return bytes
}
func SHA256DoubleString(text string, isHex bool) string {
	bytes := SHA256Double(text, isHex)
	return fmt.Sprintf("%f\n", bytes)

}
//将字节数组转换成十六进制字符串： []byte -->string
func BytesToHexString(arr []byte) string {
	return hex.EncodeToString(arr)
}

//将十六进制的字符串转换成字节数组 hex string ->[]byte
func HexStringToBytes(s string) ([]byte, error) {
	arr, err := hex.DecodeString(s)
	return arr, err
}

//十六进制字符串大端和小端颠倒
func ReverseHexString(hexStr string) string {
	arr, _ := hex.DecodeString(hexStr)
	ReverseBytes(arr)
	return hex.EncodeToString(arr)
}

//字节数组大端和小端颠倒
func ReverseBytes(data []byte) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}