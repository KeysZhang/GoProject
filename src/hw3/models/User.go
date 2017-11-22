/*
Create by zhangzemian on 2017/11/03
*/

package models

import(
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

//user's model
type User struct {
	Username string
	Password      string
}

//init for connection of database
func init(){
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:zzm15331411@/server_compute?charset=utf8")
	
}

//register
func RegisterUser(user User) string {
	o := orm.NewOrm()
	o.Using("server_compute")
	
	sql_instert := "INSERT INTO sc_users(USERNAME, PASSWORD) VALUES(?, ?)"

	if isEmpty(user) != ""{
		return isEmpty(user)
	}

	if isRegisted(user) != "" {
		return isRegisted(user)
	}
	
	_, insert_err := o.Raw(sql_instert, user.Username, user.Password).Exec()

	if insert_err != nil {
		return "the insert in database is wrong"
	}

	return ""
}


//check the user's message is right when login
func IsAllRight(user User) string {
	if(isEmpty(user) != "") {
		return isEmpty(user)
	}
	if isUserMessageRight(user) != "" {
		return isUserMessageRight(user)
	}
	return ""
}


//check the input is empty
func isEmpty(user User) string {
	if(user.Username == "" || user.Password == "" ) {
		return "the username and the password is not allowed empty\n"
	}
	return ""
}

//check the user is registed
func isRegisted(user User) string{
	o := orm.NewOrm()
	o.Using("service_compute")

	var users []User
	sql := "SELECT * FROM sc_users WHERE username = ?"
	res,err := o.Raw(sql, user.Username).QueryRows(&users)


	if err != nil {
		return "the query in database is wrong"
	}

	if res != 0{
		return "the username has been registed"
	}

	return ""
}

//check the user is registed
func isUserMessageRight(user User) string{
	o := orm.NewOrm()
	o.Using("service_compute")

	var users []User
	sql := "SELECT * FROM sc_users WHERE username = ? AND password = ?"
	res,err := o.Raw(sql, user.Username, user.Password).QueryRows(&users)


	if err != nil {
		return "the query in database is wrong"
	}

	if res == 0{
		return "the username or the password is wrong\n"
	}

	return ""
}