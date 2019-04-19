package dao

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/lizhixin1992/test/models"
	"github.com/lizhixin1992/test/models/conditions"
	"strconv"
)

type UserDao struct {
	engine *xorm.Engine
}

func NewUserDao(engine *xorm.Engine) *UserDao {
	return &UserDao{
		engine: engine,
	}
}

func (d *UserDao) Save(data *models.User) error {
	_, err := d.engine.Insert(data)
	return err
}

func (d *UserDao) Update(data *models.User) error {
	_, err := d.engine.Id(data.Id).Update(data)
	return err
}

func (d *UserDao) Delete(id int) error {
	_, err := d.engine.Id(id).Delete(new(models.User))
	return err
}

//func (d *UserDao) DeleteNot(id int) error{
//	_,err := d.engine.Id(id).Update()
//}

func (d *UserDao) GetById(id int) *models.User {
	data := &models.User{Id: id}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		data.Id = 0
		return data
	}
}

func (d *UserDao) GetAll() []models.User {
	dataList := make([]models.User, 0)
	err := d.engine.Find(&dataList)
	if err != nil {
		return dataList
	} else {
		return dataList
	}
}

//func (d *UserDao) GetListByCondition(condition *conditions.UserCondition) []models.User {
//	dataList := make([]models.User, 0)
//	commons.SetLimitSize(condition)
//	//err := d.engine.Where("").Limit(condition.StartRow,condition.EndRow).Find(&dataList)
//	d.engine.Query(setCondition(condition))
//
//}

func setCondition(cond *conditions.UserCondition) (sql string) {
	sql = "select id,name,age,addres,createTime,updateTime where"
	switch value := interface{}(cond.Name).(type) {
	case string:
		sql += " name = " + value
		fmt.Println("string", value)
	case int:
		fmt.Println("int", value)
	default:
		fmt.Println("unknown", value)
	}
	sql += " limit " + strconv.Itoa(cond.StartRow) + "," + strconv.Itoa(cond.EndRow)
	return sql
}
