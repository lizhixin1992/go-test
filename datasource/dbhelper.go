package datasource

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/lizhixin1992/test/conf"
	"log"
	"sync"
	"time"
)

var (
	masterEngine *xorm.Engine
	lock         sync.Mutex
)

func InstanceMaster() *xorm.Engine {
	if masterEngine != nil {
		return masterEngine
	}

	lock.Lock()
	defer lock.Unlock()

	//再次判断，确保不重复创建
	if masterEngine != nil {
		return masterEngine
	}

	c := conf.MasterDbConf
	driverSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		c.UserName, c.Password, c.Host, c.Port, c.DbName)
	engine, err := xorm.NewEngine(conf.DriverName, driverSource)
	if err != nil {
		log.Fatal("dbhelper.DbInstanceMaster,", err)
		return nil
	}
	//Debug模式，打印全部的SQL语句
	engine.ShowSQL(true)
	//设置时区
	engine.SetTZLocation(conf.SysTimeLocation)

	//性能优化的时候才考虑，加上本机的SQL缓存，maxElementSize是缓存的struct个数
	cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
	//手动设置缓存时间
	cacher.Expired = 60 * time.Second
	engine.SetDefaultCacher(cacher)

	masterEngine = engine
	return masterEngine
}
