package main

import (
	"fmt"
	"github.com/lizhixin1992/test/commons"
)

func main() {
	//cond := &conditions.UserCondition{Name: "111",Age:11,Addrs:"3333",Page:1,Size:2}
	//d := dao.NewUserDao()
	//d.SetCondition(cond)
	fmt.Println(commons.GetNowUnix())
}
