package models

import (
	"DataCertPhone/Hash"
	"DataCertPhone/db_mysql"
	"fmt"
)

type User struct {
	Id       int    `form:"id"`
	Phone    string `form:"phone"`
	Password string `form:"password"`
}


func (u User) AddUser() (int64, error) {

	u.Password = Hash.HASH(u.Password, "md5", false)
	result, err := db_mysql.Db.Exec("insert into user_info(phone,password)"+
		"value (?,?)", u.Phone, u.Password)
	//result, err := Db.Exec("insert into user_info(phone,password,)"+
	//	"values(?,?)",u.Phone,u.Password,)
	if err != nil {
		return -1, err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return -1, err
	}
	return rows, nil
}
func (u User) LoginUser() (*User, error) {
	u.Password = Hash.HASH(u.Password, "md5", false)

	row := db_mysql.Db.QueryRow("select phone from user_info where phone = ? and password = ?",
		u.Phone, u.Password)

	err := row.Scan(&u.Phone)

	if err != nil {
		return nil, err
	}
	return &u, err

	//var user_db User
	//
	//err := rows.Scan(&user_db.Phone,&user_db.Password)
	//fmt.Println(user_db)
	//if u.Phone ==user_db.Phone && u.Password == user_db.Password {
	//	//	return true ,nil
	//	//}
}
func QueryUserId(Phone string) (int, error) { //返回一个userid属性的值 和error
	//fmt.Println(Phone)
	row := db_mysql.Db.QueryRow("select id from user_info where phone = ?", Phone)
	var userId int
	err := row.Scan(&userId)
	if err != nil {
		return -1, err
	}

	return userId, nil
}
func QueryUserByPhone(phone string) (*User, error) {//仅用于home处理
	fmt.Println("phone 为",phone)
	row := db_mysql.Db.QueryRow("select id from user_info where phone = ?",phone)
	var user User
	err := row.Scan(&user.Id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}