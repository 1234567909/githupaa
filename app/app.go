package app

import (
	"github.com/1234567909/githupaa/config"
	"github.com/1234567909/githupaa/mysql"
)

func Init(serviceName string, apps ...string) error {
	var err error
	err = config.GetClient()
	if err != nil {
		return err
	}
	for _, val := range apps {
		switch val {
		case "mysql":
			err = mysql.InitMysql(serviceName)
			if err != nil {
				panic(err)
			}
		}
	}
	return err
}
