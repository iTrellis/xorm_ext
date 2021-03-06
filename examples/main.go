/*
Copyright © 2019 Henry Huang <hhh@rutcode.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/

package main

import (
	"fmt"

	"github.com/iTrellis/xorm_ext"
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
