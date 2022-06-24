package main

import (
	"github.com/Singh555/mycms/common/db"
	"github.com/Singh555/mycms/pkg/customers"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("../common/envs/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	//dbUrl := viper.Get("DB_URL").(string)
	dbDriver := viper.Get("DB_DRIVER").(string)
	dbHost := viper.Get("DB_HOST").(string)
	dbPort := viper.Get("DB_PORT").(string)
	dbUserName := viper.Get("DB_USERNAME").(string)
	dbPassword := viper.Get("DB_PASSWORD").(string)
	dbDatabase := viper.Get("DB_DATABASE").(string)
	dbUrl := dbDriver + "://" + dbUserName + ":" + dbPassword + "@" + dbHost + ":" + dbPort + "/" + dbDatabase
	r := gin.Default()
	h := db.Init(dbUrl)
	r.Static("/satic_files/", "../")
	customers.RegisterRoutes(r, h)
	// register more routes here

	r.Run(port)
}
