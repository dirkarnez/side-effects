package mysql

import "github.com/dirkarnez/side-effects/loader"

func init() {
	loader.Dialect = "mysql"
}