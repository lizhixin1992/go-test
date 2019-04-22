package models

//设置json:""，当在使用json转换的时候会按设置的名称转换

type User struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Age        int    `json:"age"`
	Addrs      string `json:"addrs"`
	CreateTime int    `json:"create_time"`
	UpdateTime int    `json:"update_time"`
}
