package conf

import (
	"fmt"
	"github.com/pelletier/go-toml"
)

var GlobalConf = New()

func New() *toml.Tree {
	tree, err := toml.LoadFile("./conf/GlobalConf.toml")
	if err != nil {
		fmt.Println("loadfile GlobalConf fail", err.Error())
	}
	return tree
}
