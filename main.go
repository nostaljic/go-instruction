package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

type Member struct {
	Name   string
	Age    int
	Active bool
}

func main() {

	newServer().Run()
}
func newServer() *gin.Engine {
	r := gin.Default()
	r.GET("", helloAccountHandler)
	r.GET("/db", dbMap)
	r.GET("/:name", helloAccountHandler)
	r.POST("/add", helloAccountHandler)
	return r
}
func helloAccountHandler(c *gin.Context) {

	mem := Member{"Alex", 10, true}
	c.JSON(http.StatusOK, mem)
}
func dbMap(c *gin.Context) {
	db, err := sql.Open("mysql", "root:1234@tcp(localhost:3306)/gotest")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	var version string
	db.QueryRow("SELECT VERSION()").Scan(&version)
	fmt.Println("Connected to:", version)
	c.JSON(http.StatusOK, version)
}
