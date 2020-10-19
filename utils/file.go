package utils

import (
	"io"
	"os"
)
func OpenFile(filePath string)(io.Reader,error) {
	file,err :=os.OpenFile(filePath,os.O_CREATE|os.O_RDWR,os.ModePerm)
	if err != nil{
		return nil,err
	}
	return file,nil
}
func SaveFile(fileName string,file io.Reader) (int64,error) {
	savefile, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, 777)
	if err != nil {

		return -1,err
	}
	length,err := io.Copy(savefile,file)
	if err != nil {
		return  -1,err
	}
	return  length,nil
}
