package blockchain

import (
	"DataCertPhone/utils"
	"bytes"
	"fmt"
	"time"
)
var TimeTotals []int64
/*		system given difficulty: how many zeros are in front of the []byte
//[]byte 256位   系统给定的难度:  前多少位数字为为零?       难度与总矿工的数量有关. :目前自己有多少个协程在运行so
与上一个周期矿工挖出的一个平均时间有关			golang don't know how many goroutines his has running right now
//    系统给定的值 [1,0,0,0,0 ....  0,0,0]		我们矿工挖到的[]byte要小于系统给定的这个值

2016个区块出块总用时
*/
func Minner(sysBytes []byte) (int64,error) {//sysBytes 为 系统给的[]byte 返回一个挖矿的时间和一个error
	timeStart := time.Now().UnixNano()
	fmt.Println(timeStart)
	var i int64 = 0
	for  {
		i++
		byteI,err :=  utils.Int64ToByte(i)
		if err != nil {
			fmt.Println(err.Error())
			return -1,err
		}
		hashByteI,err := utils.SHA256HashByte(byteI)
		if err != nil {
			fmt.Println(err.Error())
			return -1,err

		}
		if CompareBytes(hashByteI,sysBytes){//当hashByteI < sysBytess 时 返回true
			timeEnd := time.Now().UnixNano()
			fmt.Println(timeEnd)

			timeTotal := timeEnd - timeStart
			//fmt.Println("花费了我",timeTotal,"秒")
			fmt.Println("我找到这个数学题的解决方案啦! ")
			return timeTotal,nil
		}
		fmt.Printf("第 %v 次寻找hash值\n",i)
	}
}
func GetSySBytes()  {
	a := []byte{}
	for i := 1;i<=255 ;i++  {
		a = append(a, 0)
	}
	a = append(a, 1)

	fmt.Println()
}
func CompareBytes(a, b []byte)bool {
	if bytes.Compare(a, b) < 0 {//a小于b
		//fmt.Println("a less than b")
		return true
	} else if bytes.Compare(a, b) > 0 {//a大于b
		//fmt.Println("a greater than b")
		return false

	} else if bytes.Compare(a, b) == 0 {//a与b相等
		//fmt.Println("a equals b")
		return false

	}
	return false
}