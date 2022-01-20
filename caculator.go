package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func main() {
	r := gin.Default()
	r.POST("/calculator", func(c *gin.Context) {
		x, _ := strconv.Atoi(c.PostForm("x"))
		y, _ := strconv.Atoi(c.PostForm("y"))
		operator := c.PostForm("operator")

		var result = ""
		switch operator {
		case "+":
			result = fmt.Sprint(x + y)
		case "-":
			result = fmt.Sprint(x - y)
		case "*":
			result = fmt.Sprint(x * y)
		case "/":
			result = fmt.Sprint(x / y)
		default:
			result = "operator error!"
		}

		c.JSON(200, gin.H{
			"result": result,
		})

	})
	r.Run(":8000")
}
