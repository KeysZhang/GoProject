package entities

import (
	"fmt"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"

	_ "github.com/go-sql-driver/mysql"
)

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
	fmt.Println(has)
	if (!has){
		egine.CreateTables(&user)
	}
	myegine = egine
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
