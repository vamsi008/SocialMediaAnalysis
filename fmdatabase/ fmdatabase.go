package fmdatabase

import (
	"database/sql"
	"fmt"
	"sync"

	"github.com/BurntSushi/toml"
	_ "github.com/go-sql-driver/mysql"
)

var (
	// DBCon is the connection handle
	// for the database
	instance *sql.DB
	once     sync.Once
)

/*type singleton struct {
}*/

func GetInstance() *sql.DB {
	once.Do(func() {
		var err error
		instance, err = sql.Open("mysql", "root:pramati123@/fullmetal")
		/*instance.SetMaxOpenConns(500)
		instance.SetMaxIdleConns(10)
		instance.SetConnMaxLifetime(1)*/
		if err != nil {
			fmt.Println("Into the error while creaing the db connection")
			panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
		}
	})
	return instance
}

type tomlConfig struct {
	Title string
	DB    database `toml:"database"`
}

type database struct {
	DbServer string `toml:"db_server"`
	User     string
	Password string `toml:"connection_max"`
}

func init() {

	fmt.Println("Initializing data base connection")

	var config tomlConfig

	//dev_server := viper.GetString("development.server")
	if _, err := toml.DecodeFile("/usr/share/fullmetal/config.toml", &config); err != nil {
		fmt.Println(err)
		return
	}

	/*var err error
	DBCon, err = sql.Open()

	if err != nil {
		fmt.Println("Into the error while creaing the db connection")
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}*/

}
