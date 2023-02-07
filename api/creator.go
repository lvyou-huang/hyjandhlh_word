package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"test/dao"
	"test/handler"
)

func Creator() {
	Router := gin.Default()
	Router.POST("/editor/drafts/new", func(c *gin.Context) {
		data, _ := c.GetRawData()
		var m map[string]any
		_ = json.Unmarshal(data, &m)
		article_title := m["article_title"]
		article_content := m["article_content"]
		category := m["category"]
		label := m["label"]
		abstract := m["abstract"]
		date := m["date"]
		column1 := m["column1"]
		column2 := m["column2"]
		column3 := m["column3"]
		phoneoremail, err := c.Cookie("phoneoremail")
		if err != nil {
			log.Println(err)
			c.JSON(400, gin.H{
				"msg": "cookie过期",
			})
			return
		}
		Db, err := dao.OpenDb()
		if err != nil {
			log.Println(err)
			c.JSON(500, gin.H{
				"msg": "服务器错误",
			})
		}

		result, err := Db.Exec("insert into article (article_title,article_content,date,phoneoremail,status,category,label,abstract,column1,column2,column3) value (?,?,?,?,?,?,?,?,?,?,?)", article_title, article_content, date, phoneoremail, "auditng", category, label, abstract, column1, column2, column3)
		log.Println(result)
		if err != nil {
			log.Println(err)
			c.JSON(500, gin.H{
				"msg": "无法上传文章",
			})
		}
		c.JSON(200, gin.H{
			"msg": "正在审核中",
		})
		c.Redirect(300, "http://localhost:8080/published")
	})
	Router.GET("/published", func(c *gin.Context) {

	})
	Creator := Router.Group("/creator", handler.Auth())
	{
		Creator.GET("/home", func(c *gin.Context) {

		})
		Content := Creator.Group("/content")
		{
			Article := Content.Group("/article")
			{
				Article.POST("/essays", func(c *gin.Context) {
					status := c.Query("status")
					if status == "" {
						status = "all"
					}

				})
				Article.GET("/drafts", func(c *gin.Context) {

				})
			}
			Content.POST("/column", func(c *gin.Context) {
				column_title := c.Query("column_title")
				column_introduction := c.Query("column_introduction")
				phoneoremail, err := c.Cookie("phoneoremail")
				if err != nil {
					log.Println(err)
					c.JSON(400, gin.H{
						"msg": "cookie过期",
					})
					return
				}
				if column_title == "" {
					log.Println("标题不能为空")
					c.JSON(417, gin.H{
						"msg": "标题不能为空",
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
				result, err := Db.Exec("insert into column (column_title,column_introduction,phoneoremail,id)value (?,?,?,?)", column_title, column_introduction, phoneoremail, nil)
				log.Println(result)
				if err != nil {
					log.Println(err)
					c.JSON(500, gin.H{
						"msg": "无法创建新专栏",
					})
					return
				}
				c.JSON(200, gin.H{
					"msg": "创建成功",
				})
				log.Println("创建成功")
			})

		}

	}
	Router.POST("/comment", func(c *gin.Context) {

	})
	Router.Run()
}
