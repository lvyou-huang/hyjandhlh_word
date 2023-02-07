package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"test/dao"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		phoneoremail, err := c.Cookie("phoneoremail")
		fmt.Printf("phoneoremail : %s \n", phoneoremail)
		if err != nil {
			c.Abort()
			c.JSON(400, gin.H{
				"msg": "cookie验证失败1",
			})
			return
		}

		Db, err := dao.OpenDb()
		if err != nil {
			log.Println(err)
			c.JSON(500, gin.H{
				"msg": "服务器错误",
			})
			return
		}
		row := Db.QueryRow("select phoneoremail from user where phoneoremail = ?", phoneoremail)
		var phoneoremail2 string
		err = row.Scan(&phoneoremail2)
		if err != nil {
			log.Println(err)
			c.JSON(500, gin.H{
				"msg": "cookie验证失败，数据库中不存在该用户",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
