package main

import (
	"test/api"
	"test/dao"
)

func main() {
	dao.OpenDb()
	api.User()

}
