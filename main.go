package main

import (
	_ "github.com/zboyco/goutlis/encry"
	_ "github.com/zboyco/goutlis/list"
	_ "github.com/zboyco/goutlis/queue"
	_ "github.com/zboyco/goutlis/redis"
)

func main() {
	// 库下不增加import的话go get 不能 get
}
