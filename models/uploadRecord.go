package models

import (
	"DataCertPhone/db_mysql"
	"fmt"
	//"time"
)

type UploadRecord struct {
	Id        int
	UserId    int
	FileName  string
	FileSize  int64
	FileCert  string
	FileTitle string
	CertTime  int64
	CerTimeFormat string
}




//func (u *UploadRecord)GetVaule()(*UploadRecord,error){
//	u.UserId ,err := u.QueryUserId()
//
//}
func (u *UploadRecord) AddFiles() (int64, error) {
	result, err := db_mysql.Db.Exec("INSERT INTO upload_record" +
		"(user_id,file_name,file_size,file_cert,file_title,cert_time)" +
		"value (?,?,?,?,?,?)",u.UserId,u.FileName,u.FileSize,u.FileCert,u.FileTitle,u.CertTime)
	if err != nil {
		fmt.Println("数据库存储文件出错，请重试！")
		return -1, err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return -1, err
	}
	return rows, nil
}
func QueryRecordsByUserId(userId int)([]UploadRecord,error)  {
	fmt.Println("当前登录用户的ID为:",userId)
	rows,err := db_mysql.Db.Query("select id,user_id,file_name,file_size,file_cert,file_title,cert_time from upload_record where user_id =?",userId)
	if err != nil {
		return nil,err
	}
	var records = make([]UploadRecord, 0)
	var record =UploadRecord{}
	for rows.Next()  {
		err = rows.Scan(&record.Id,&record.UserId,&record.FileName,&record.FileSize,&record.FileCert,&record.FileTitle,&record.CertTime)
		if err != nil {
			return nil,err
		}
		//t := time.Unix()
		//t := time.Unix(record.CertTime,0)
		// tStr := t.Format("2006年01月02日 15：04：05") // 2006/01/02 15：04：05
		records = append(records, record)
	}
	return records,nil
}