package utils

import (
	"bytes"
	"encoding/binary"
)

//将一个int64类型转化为[]byte类型
func Int64ToByte(num int64)([]byte,error){
	//Buffer : 缓冲区
	buff := new(bytes.Buffer)//通过new实例化一个缓冲区
	//buff.Write()   		//通过列的通用方法写入数据
	//buff.Bytes()			//通过Bytes方法从缓冲区获取数据
	/*
	两种排列方式:
		大端位序排列: binary.BigEndian
		小端位序排列: binary.LitterEndian
	 */
	err := binary.Write(buff,binary.LittleEndian,num)
	if err != nil {
		return nil,err
	}
	return buff.Bytes(),nil
}
//将字符串转换为字节切片
func StringToBytes(data string)[]byte  {
	return []byte(data)
}
func IntToBytes(data int)([]byte,error){
	buffer := new(bytes.Buffer)
	err :=binary.Write(buffer,binary.LittleEndian,data)
	if err != nil {
		return nil,err
	}
	return buffer.Bytes(),nil
}
//func BytesToInt(data []byte)int{
//	buff := new(bytes.Buffer)
//	binary.Write(buff,binary.LittleEndian,data)
//
//}