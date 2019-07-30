package main

import (
	"fmt"
	_ "github.com/dirkarnez/side-effects/dialects/mssql"
	// _ "github.com/dirkarnez/side-effects/dialects/mysql"
	"github.com/dirkarnez/side-effects/loader"
)

func main() {
	fmt.Println(loader.Dialect)
}