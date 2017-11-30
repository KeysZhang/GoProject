package entities

import (
	"fmt"
	"time"
	"strconv"
	"testing"
	"math/rand"
)

func TestUserService(t *testing.T){

	ini_mes := "\nPlease input the number of operator you want:\n1.Insert\n2.FindAll\n3.FindById\n4.exit"
	fmt.Println(ini_mes)

	op := 1
	fmt.Print("operatorNum:")
	fmt.Print(op)

	for(op != 4){

		t := time.Now()
		t_str := t.Format("2006-01-02 15:04:05")
		prefix_message := "[Userinfo-test][Time:" + t_str + "]"

		//插入数据
		if (op == 1){
			username := "zhangzemian"
			departname := "javadevelop"
			fmt.Println("Uername:" + username + ", departname:" + departname)

			user := NewUserInfo(UserInfo{UserName:username, DepartName:departname})
			UserInfoService.Save(user)

			mes_userid := "UserID:" + strconv.Itoa(user.UID)
			mes_username := "Username:" + user.UserName
			mes_departname := "Departname:" + user.DepartName
			mes_time := "Createtime:" + user.CreateAt.Format("2006-01-02 15:04:05")
			message := prefix_message + "[Insert]" + mes_userid + "," + mes_username + "," + mes_departname + "," + mes_time
			fmt.Println(message)
		}
		
		//查找所有用户
		if(op == 2){
			query_all_userlist := UserInfoService.FindAll()

			if len(query_all_userlist) == 0{
				message := prefix_message + "[FindAll]There is not user in database"
				fmt.Println(message)
			} else {
				for i := 0; i < len(query_all_userlist); i++{
					user := query_all_userlist[i]
					mes_userid := "UserID:" + strconv.Itoa(user.UID)
					mes_username := "Username:" + user.UserName
					mes_departname := "Departname:" + user.DepartName
					mes_time := "Createtime:" + user.CreateAt.Format("2006-01-02 15:04:05")
					message := prefix_message + "[FindAll]" + mes_userid + "," + mes_username + "," + mes_departname + "," + mes_time
					fmt.Println(message)
				}
			}
		}

		//根据id查找用户
		if (op == 3){
			id_mes := "Please input the id of users"
			fmt.Println(id_mes)
			fmt.Print("UserID:")
			id := rand.Intn(3)			

			user := UserInfoService.FindByID(id)

			if (user == nil){
				message := prefix_message + "[FindByID]The user is not exist"
				fmt.Println(message)
			} else{
				mes_userid := "UserID:" + strconv.Itoa(user.UID)
				mes_username := "Username:" + user.UserName
				mes_departname := "Departname:" + user.DepartName
				mes_time := "Createtime:" + user.CreateAt.Format("2006-01-02 15:04:05")
				message := prefix_message + "[FindByID]" + mes_userid + "," + mes_username + "," + mes_departname + "," + mes_time
				fmt.Println(message)
			}
		}

		//退出
		if (op == 4){
			fmt.Println("Bye")
			break
		}
		fmt.Println(ini_mes)
		fmt.Print("operatorNum:")
		op = op + 1
		fmt.Print(op)
	}
	
}