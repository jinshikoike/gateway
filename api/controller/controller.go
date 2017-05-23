package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/jinshikoike/gateway/api/models"
	"github.com/jinshikoike/gateway/config"
	"github.com/jinshikoike/gateway/database"
	"github.com/labstack/echo"
)

// Caller call clapton runner api

// TODO: refactoring
func Caller(c echo.Context) error {
	// TODO: pool db connection
	db, err := database.Connection(config.DBUser(), config.DBPass(), config.DBHost(), config.DBPort(), config.DBName())
	defer db.Close()
	if err != nil {
		panic(err)
	}

	// TODO: support parameter, formdata ... etc
	hostName := strings.Split(c.Request().Host, ":")[0]
	path := c.Request().URL.Path

	fmt.Println(hostName)
	fmt.Println(path)

	var gateway models.Gateway
	db.Table("gateways").Select("gateways.path, gateways.id, gateways.host_id, gateways.clapton_id, hosts.id, hosts.name").Joins("left join hosts on gateways.host_id = hosts.id").Where("gateways.path = ?", path).Where("hosts.name = ?", hostName).First(&gateway)
	fmt.Println(gateway)
	if gateway.ID == 0 {
		return c.String(http.StatusNotFound, "Gateway Not Found")
	}
	if gateway.HostID == 0 {
		return c.String(http.StatusNotFound, "Host Not Found")
	}
	if gateway.ClaptonID == 0 {
		return c.String(http.StatusNotFound, "Host Not Found")
	}
	db.Model(gateway).Related(&gateway.Host)
	db.Model(gateway).Related(&gateway.Clapton)

	fmt.Println(gateway.Host.Name)
	fmt.Println(gateway.Host.ID)
	fmt.Println(gateway.Clapton.UUID)
	fmt.Println(gateway.Path)

	resp, err := http.Get("http://localhost:8080/launch/" + gateway.Clapton.UUID)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	return c.String(http.StatusOK, string(body))
}

// Test return test
func Test(c echo.Context) error {
	return c.String(http.StatusOK, "caller called")
}
