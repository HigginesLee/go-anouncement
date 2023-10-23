package main

import (
	"go_announce/service"
	"go_announce/utils"
)

func main() {
	utils.DBInit()
	service.DouDianAnounceMent()
}
