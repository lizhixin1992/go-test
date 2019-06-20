package controller

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/lizhixin1992/test/commons"
	"github.com/lizhixin1992/test/models"
	"github.com/lizhixin1992/test/services"
	"log"
)

type UserController struct {
	Ctx     iris.Context
	Service services.UserService
}

////测试es返回数据格式
//func (c *UserController) Get() mvc.Result {
//	query := elastic.NewMatchQuery("passage", "elk rocks")
//	result := commons.MatchQuery("book1", "english", query, 0, 10)
//
//	return commons.SetResponseSuccessData(result)
//}

//localhost:8080/
func (c *UserController) Get() mvc.Result {
	dataList := c.Service.GetAll()
	return commons.SetResponseSuccessData(dataList)
}

//localhost:8080/{id}
func (c *UserController) GetBy(id int) mvc.Result {
	data := c.Service.GetById(id)
	return commons.SetResponseSuccessData(data)
}

//localhost:8080/save
func (c *UserController) PostSave() mvc.Result {
	info := models.User{}
	err := c.Ctx.ReadJSON(&info)
	if err != nil {
		log.Fatal(err)
		return commons.SetResponseFail()
	} else {
		date := int(commons.GetNowUnix())
		info.CreateTime = date
		info.UpdateTime = date
		c.Service.Save(&info)
		return commons.SetResponseSuccess()
	}
}

func (c *UserController) BeforeActivation(b mvc.BeforeActivation) {
	fmt.Println("******************* before *********************")
	//fmt.Println("Cache", "****************", commons.Cache.Get("test1"))
	fmt.Println("CacheCluster", "****************", commons.CacheCluster.Get("test"))

	cp := commons.CachePool.Get()
	v, _ := redis.String(cp.Do("GET", "redisUtil"))
	defer cp.Close()
	fmt.Println("CacheCluster", "****************", v)
}

func (c *UserController) AfterActivation(b mvc.AfterActivation) {
	fmt.Println("******************* after *********************")
}
