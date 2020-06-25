package datasource

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hongyukeji/easy-go/utils"
	"github.com/jinzhu/gorm"
)

var Eloquent *gorm.DB

func init() {
	var (
		Host      = getEnv("DB_HOST", "127.0.0.1")
		Port      = getEnv("DB_PORT", "3306")
		Database  = getEnv("DB_DATABASE", "easy_go")
		Username  = getEnv("DB_USERNAME", "root")
		Password  = getEnv("DB_PASSWORD", "123456")
		Charset   = getEnv("DB_CHARSET", "utf8")
		Collation = getEnv("DB_COLLATION", "utf8_general_ci")
	)
	mysqlConfig := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&collation=%s&parseTime=True&loc=Local", Username, Password, Host, Port, Database, Charset, Collation)

	var err error
	Eloquent, err = gorm.Open("mysql", mysqlConfig)
	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}
	if Eloquent.Error != nil {
		fmt.Printf("datasource error %v", Eloquent.Error)
	}

	// orm "github.com/hongyukeji/easy-go/datasource"
	// Models "github.com/hongyukeji/easy-go/models"
	// orm.Eloquent.AutoMigrate(&Models.User{}, &Models.Article{})
	// Eloquent.AutoMigrate(&Models.User{})
	//addr := iris.Application.ConfigurationReadOnly("").GetOther()["AppAddr"].(string)
	//fmt.Printf("mysql connect error %v", addr)
}

func getEnv(key string, def string) string {
	return utils.GetEnv(key, def)
}
