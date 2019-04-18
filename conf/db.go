package conf

const DriverName = "mysql"

type DbConf struct {
	Host     string
	Port     int
	UserName string
	Password string
	DbName   string
}

var MasterDbConf DbConf = DbConf{
	Host:     "127.0.0.1",
	Port:     3306,
	UserName: "root",
	Password: "root",
	DbName:   "test_db",
}
