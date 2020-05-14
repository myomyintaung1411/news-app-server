package main

import (
	"huana/route"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"os"
	"huana/common"
	"net/http"
)

func main() {

	InitConfig()
	db := common.InitDB()
	defer db.Close()

	r := gin.Default()
	
	r = route.CollectRoute(r)
	r.Static("/images", "./images")
	r.StaticFS("/public", http.Dir("public"))
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run()) 
}
func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

// func upload(c *gin.Context) {
// 	file, header, err := c.Request.FormFile("file") 
// 	if err != nil {
// 	  c.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
// 	  return
// 	}
// 	filename := header.Filename
// 	out, err := os.Create("public/" + filename)
// 	if err != nil {
// 	  log.Fatal(err)
// 	}
// 	defer out.Close()
// 	_, err = io.Copy(out, file)
// 	if err != nil {
// 	  log.Fatal(err)
// 	}
// 	filepath := "http://localhost:3000/public/" + filename
// 	c.JSON(http.StatusOK, gin.H{"filepath": filepath})
//   }



