package config

import (
	"fmt"

	"github.com/gomodul/envy"
	_ "github.com/go-sql-driver/mysql"

)

type databaseItem struct {
	Drivername     string
	Datasourcename string
}

type database struct {
	MySQL databaseItem
}

var DB = database{
	MySQL: databaseItem{
		Drivername: envy.Get("MYSQL_DRIVER", "mysql"),
		Datasourcename: fmt.Sprintf(
			"%v:%v@tcp(%v:%v)/%v?parseTime=true&charset=utf8",
			envy.Get("mysql_user", "root"),
			envy.Get("mysql_pass", "abcd"),
			envy.Get("mysql_addr", "127.0.0.1"),
			envy.Get("mysql_port", "3306"),
			envy.Get("mysql_name", "bitly"),
		),
	},
}
