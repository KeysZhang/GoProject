# 实验五  xorm构建数据库服务

[TOC]

</br>

##一. xorm概述

xorm是一个简单而强大的Go语言ORM库. 通过它可以使数据库操作非常简便。xorm的目标并不是让你完全不去学习SQL，我们认为SQL并不会为ORM所替代，但是ORM将可以解决绝大部分的简单SQL需求。xorm支持两种风格的混用。

</br>

## 二. xorm使用

#### 1. 加载驱动：

```
import (
	"fmt"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	_ "github.com/go-sql-driver/mysql"
)
```

`_ "github.com/go-sql-driver/mysql"` 启动包在 init() 阶段，通过调用` xorm.NewEngine(····)`注册驱动到应用中。

```
var myegine *xorm.Engine

func init() {
	egine, err := xorm.NewEngine("mysql", "root:zzm15331411@tcp(localhost:3306)/test?charset=utf8")
	if err != nil {
		panic(err)
	}
	var user UserInfo
	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, "golang_xorm_")
	egine.SetTableMapper(tbMapper)
	has, _ := egine.IsTableExist(&user)
	if (!has){
		egine.CreateTables(&user)
	}
	myegine = egine
}
```

`tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, "golang_xorm_")`可以创建一个在SnakeMapper的基础上在命名中添加统一的前缀，最终生成的表格的名字是golang_xorm_user_info

创建表使用`engine.CreateTables()`，参数为一个或多个空的对应Struct的指针。

`engine.IsTableExist()`判断表是否为空，参数和CreateTables相同

####2. 访问 CRUD与服务

```
package entities

//UserInfoAtomicService .
type UserInfoAtomicService struct{}

//UserInfoService .
var UserInfoService = UserInfoAtomicService{}

// Save .
func (*UserInfoAtomicService) Save(u *UserInfo) error {
	_, err := myegine.Insert(u)
	checkErr(err)
	return nil
}

// FindAll .
func (*UserInfoAtomicService) FindAll() []UserInfo {
	everyone := make([]UserInfo,0)
	err := myegine.Find(&everyone)
	checkErr(err)
	return everyone
}

// FindByID .
func (*UserInfoAtomicService) FindByID(id int) *UserInfo {
	users := make([]UserInfo, 0)
	err := myegine.Id(id).Find(&users)
	checkErr(err)
	return &users[0]
}
```

`myegine.Insert(u)`将结构体数据插入数据库中

查询多条数据使用`Find`方法，Find方法的第一个参数为`slice`的指针或`Map`指针，即为查询后返回的结果。

#### 3. 数据服务测试

```
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
```

.*test.go文件可用于集成测试，测试这三个模块结合起来使用是否可以正常运行，输入命令`go test userinfo_test.go`，测试结果如下图：

![img](实验截图/1.png)

</br>

## 三. database/sql 与 orm 实现的异同

####编程效率

orm框架ORM使所有的数据表都按照统一的标准精确地映射成结构体，使系统在代码层面保持准确统一，提供了很多高效便捷的接口，可以使用结构体自动对数据库进行增删改查，而database/sql则必须通过sql语言对数据库进行增删改查。使用orm代码量大量减少，编程效率显著提高，减少了重复的劳动，比database/sql优越

#### 程序结构

使用orm实现数据库操作，结构简单易懂，因为orm已经提供了现成的接口可以自动进行增删改查，而database/sql则需要自定义各种接口，结构较为复杂，使用起来难度较高

#### 服务性能

orm 是用的反射技术、牺牲性能获得易用性，在处理多表联查、where条件复杂之类的查询时，ORM的语法会变得复杂且猥琐，而且越是功能强大的ORM越是消耗内存，因为一个ORM Object会带有很多成员变量和成员函数，因此性能明显比database/sql低很多

</br>

## 四. ab测试

以FindAll为例

xorm:

![img](实验截图/2.png)

database/sql

![img](实验截图/3.png)

两张表对比，可以看出orm每秒请求个数（299.93）比database/sql（233.87）多一点，说明orm在处理少量数据时比database/sql高效一点