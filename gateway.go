package main

import (
	"fmt"

	"github.com/jinshikoike/gateway/api/models"
	"github.com/jinshikoike/gateway/config"
	"github.com/jinshikoike/gateway/database"
)

func main() {
	db, err := database.Connection(config.DBUser(), config.DBPass(), config.DBHost(), config.DBPort(), config.DBName())
	defer db.Close()
	if err != nil {
		panic(err)
	}

	var gateway models.Gateway
	db.Table("gateways").Select("gateways.path, gateways.id, gateways.host_id, hosts.id, hosts.name").Joins("left join hosts on gateways.host_id = hosts.id").Where("gateways.path = ?", "/tests").Where("hosts.name = ?", "koike.co.jp").First(&gateway)
	db.Model(gateway).Related(&gateway.Host)
	db.Model(gateway).Related(&gateway.Clapton)
	fmt.Println(gateway.Host.Name)
	fmt.Println(gateway.Clapton.UUID)
	fmt.Println(gateway.Host.ID)
	fmt.Println(gateway.Path)
}
