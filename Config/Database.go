package Config

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

type Dbconfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func Hello() {
	fmt.Println("hello")
}

func BuidDbConfig() *Dbconfig {
	dbConfig := Dbconfig{
		Host:     "0.0.0.0",
		Port:     3306,
		User:     "dt_admin",
		DBName:   "todo",
		Password: "BackSpace2;",
	}
	return &dbConfig
}

func DbUrl(dbConfig *Dbconfig) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}
