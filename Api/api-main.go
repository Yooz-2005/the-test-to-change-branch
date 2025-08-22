package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello, World!")

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello, World!"})
		fmt.Println("hello,我是Zachary分支")
		fmt.Println("什么情况 我重新来一次")
	})
}
