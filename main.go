package main

import (
	"github.com/liuhongdi/digv26/router"
)

func main() {
	//引入路由
	r := router.Router()
	//run
	r.Run(":8080")
}
