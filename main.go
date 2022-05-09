package main

import (
	"aydin-tracker/aydinins-backend/pkg/common/db"
	"aydin-tracker/aydinins-backend/pkg/controller"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)

	dbUrl := viper.Get("DB_URL").(string)

	r := gin.Default()

	h := db.Init(dbUrl)

	controller.RouteListener(r, h)
	r.Run(port)
}
