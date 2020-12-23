package services

import (
	"github.com/lizhixin1992/go-test/dao"
	"github.com/lizhixin1992/go-test/datasource"
	"github.com/lizhixin1992/go-test/models"
)

type UserService interface {
	Save(data *models.User) error
	Update(data *models.User) error
	Delete(id int) error
	GetById(id int) *models.User
	GetAll() []models.User
}

type userService struct {
	dao *dao.UserDao
}

func NewUserservice() *userService {
	return &userService{
		dao: dao.NewUserDao(datasource.InstanceMaster()),
	}
}

func (u *userService) Save(data *models.User) error {
	return u.dao.Save(data)
}

func (u *userService) Update(data *models.User) error {
	return u.dao.Update(data)
}

func (u *userService) Delete(id int) error {
	return u.dao.Delete(id)
}

//func (d *UserDao) DeleteNot(id int) error{
//	_,err := d.engine.Id(id).Update()
//}

func (u *userService) GetById(id int) *models.User {
	return u.dao.GetById(id)
}

func (u *userService) GetAll() []models.User {
	return u.dao.GetAll()
}
