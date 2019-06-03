package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func init() {
	getConfig()
}

func getConfig() {
	viper.SetConfigName("config") // name of config file (without extension)

	viper.AddConfigPath("./conf/")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

func GetMysqlDSN() string {
	usr := viper.GetString("mysql.user")
	pwd := viper.GetString("mysql.password")
	host := viper.GetString("mysql.host")
	port := viper.GetString("mysql.port")
	db := viper.GetString("mysql.db")
	charset := viper.GetString("mysql.charset")
	//return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=Local", usr, pwd, host, port, db, charset)
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Asia%%2FShanghai", usr, pwd, host, port, db, charset)
}

func GetSMTPConfig() (server string, port int, user, pwd string) {
	server = viper.GetString("mail.smtp")
	port = viper.GetInt("mail.smtp-port")
	user = viper.GetString("mail.user")
	pwd = viper.GetString("mail.password")
	return
}

func GetServerURL() (url string) {
	url = viper.GetString("server.url")
	return
}
