package main

import (
	controller "Quizz/Controller"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	r := gin.Default()
	controller.SetupItemRouter(r)
	r.Run(":8080")

}
