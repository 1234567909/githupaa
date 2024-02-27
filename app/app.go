package app

import "github.com/1234567909/githupaa/mysql"

func Init(apps ...string) error {
	var err error
	for _, val := range apps {
		switch val {
		case val:
			err = mysql.InitMysql()
		}
	}
	return err
}
