package main

import (
	"encoding/json"
	"fmt"
	"github.com/lizhixin1992/go-test/models"
)

func main() {
	info := models.User{
		Id:         1,
		Name:       "1111",
		Age:        12,
		Addrs:      "sssss",
		CreateTime: 213131213,
		UpdateTime: 123213123,
	}
	//fmt.Println(info)
	str, _ := json.Marshal(info)
	fmt.Println("%s\n", str)
	fmt.Printf("%s\n", str)

	fmt.Println("*******************************************************")

	str0, _ := json.MarshalIndent(info, "", " ")
	fmt.Println("%s\n", str0)
	fmt.Printf("%s\n", str0)
}
