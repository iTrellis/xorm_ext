// GNU GPL v3 License
// Copyright (c) 2019 github.com:iTrellis

package main

import (
	"fmt"

	"github.com/iTrellis/go-xorm_ext"
	"xorm.io/xorm"
)

var (
	engines map[string]*xorm.Engine
)

type Sample struct {
	ID   string `xorm:"id"`
	Name string `xorm:"Name"`
}

func (*Sample) TableName() string {
	return "sample"
}

func main() {
	var err error
	engines, err = xorm_ext.NewEnginesFromFile("../mysql.yaml")
	if err != nil {
		panic(err)
	}

	engine := engines[xorm_ext.DefaultDatabase]
	if engine == nil {
		panic(xorm_ext.ErrNotFoundXormEngine)
	}

	var ss []Sample
	err = engine.NewSession().Find(&ss)
	if err != nil {
		panic(err)
	}
	fmt.Println(ss)
}
