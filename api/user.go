package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"strings"
	"test/dao"
	"test/handler"
	"test/model"
)

func User() {
	Router := gin.Default()

	Router.LoadHTMLGlob("templates/*")
	Router.Static("/static", "./static")
	Router.Static("/fonts", "./fonts")
	//注册
	Router.GET("/signUp", func(c *gin.Context) {
		c.HTML(200, "reg.html", gin.H{})

	})

	Router.POST("/signUp", func(c *gin.Context) {

		data, _ := c.GetRawData()
		var m map[string]interface{}
		_ = json.Unmarshal(data, &m)
		phoneoremail := m["phoneoremail"]
		password := m["password"]

		Db, err := dao.OpenDb()
		if err != nil {
			log.Println(err)
			c.JSON(500, gin.H{
				"msg": "服务器错误",
			})
			return
		}
		result, err := Db.Exec("insert into user (id,password,phoneoremail) value (?,?,?)", nil, password, phoneoremail)

		log.Println(result)
		if err != nil {
			log.Println(err)
			c.JSON(500, gin.H{
				"msg": "手机或者邮箱已经存在",
			})
			return
		}
		result, err = Db.Exec("insert into userinfo (name,position,company,web,intruduce,phoneoremail,cover) value (?,?,?,?,?,?,?)", phoneoremail, "空", "空", "空", "空", phoneoremail, "data:image/png;base64,UklGRoYCAABXRUJQVlA4IHoCAADQEQCdASpkAGQAPpFEnUulo6KhohqqyLASCWMA1bYBP95XIi4EG2a8wHm3eg/eVd5ZX8U8FB6tFMO5eFgK0Jqjl7gd+OLhZCPv+ikSZkwJ4AC4sOU1dmfDQYBpEMMIWm6mPWnOntvsgOF27mjuoEA+WyMQlE+3fau51WNI/gwC3urjpaNLgaa+oRm3HMpblAZDeV/Og0jAAP7jd2ezomf9wdiu2LGAAERGuugcRU32y37bL/r7cVdYG+quDlEBLUsJje1U7jI8oqL13vFZI6oip4YkLI/Bw3SZrfPGUBjpYubbRfXilAg8KPs5/aCS51rbmvtosiWJ3AV3fqEr++cDp4PUD9oG6wXgeprFnvqeAjuBKBLWixjM5g6UD5f0/gD4UuWD/Tc4ST+Wrk9SyPUJzyOwOh5eXvUfpoPz/fJWyH7eZ5e15bGX10Y+5PvpU6YemOfW+797XUJs4IxkIRhSFB7BcX6MJPrelWYTmSEXcBmOjhOPkyB6QPtPD0ysSOzF1HPPApw8LW3gyiRgOQ+FWdKpLhXBwnzPdAR2ZznxoDirrn8bq+s4LOUWw+my3io/rHUPcGc5E3ks2IEtWIslPUQw0j3Uj2v4A61XavpfNKaelyRLq2zgqr7sa6ioIGyNPne9xiFdyQucfvbfSjD61fZx8E1R/F73tvXnxYuxmmZQQ6SXmyevM3vztSoTKiXGdITepuAQz1hQMh6kw5ujO1zGt6l5zgT+ymnJ32ga0ArbWeT716qBjClXJkubRKy+JQj1kG5geD/6uEvrQ/sztqrSw4g8OVrkCHFcol3VAwLzQmmJONIKVfcmAe1xKj0tkghEM1hFSAAA")
		result, err = Db.Exec("insert into accountinfo (phone,weixin,xinlang,github,phoneoremail) value (?,?,?,?,?)", 0, "空", "空", "空", phoneoremail)

		log.Println(result)
		if err != nil {
			log.Println(err)
			Db.Exec("delete from user where phoneoremail=?", phoneoremail)
			c.JSON(500, gin.H{
				"msg": "用户信息初始化失败",
			})
			return
		}
		c.JSON(200, gin.H{
			"msg": "注册成功",
		})
		log.Println("注册成功")
		Db.Close()
	})

	//登录
	Router.GET("/login", func(c *gin.Context) {
		c.HTML(200, "login.html", gin.H{})
	})
	Router.POST("/login", func(c *gin.Context) {
		//c.HTML(200, "login.html", gin.H{})

		data, _ := c.GetRawData()
		var m map[string]string
		_ = json.Unmarshal(data, &m)

		phoneoremail := m["phoneoremail"]
		log.Println(phoneoremail)
		password := m["password"]

		var user struct {
			Id           int
			Phoneoremail string
			Password     string
		}
		db, err := dao.OpenDb()
		if err != nil {
			log.Println(err)
			return
		}
		_, err = db.Query("select * from user where phoneoremail =? ", phoneoremail)

		if err == nil {

			Db, err := dao.OpenDb()
			if err != nil {
				log.Println(err)
				c.JSON(200, gin.H{
					"msg": "服务器错误1",
				})
				return
			}
			row := Db.QueryRow("select * from user where phoneoremail = ?", phoneoremail)
			if err != nil {
				log.Println(err)
				c.JSON(200, gin.H{
					"msg": "服务器错误2",
				})
				return
			}
			err = row.Scan(&user.Id, &user.Phoneoremail, &user.Password)
			if err != nil {
				log.Println(err)
				c.JSON(200, gin.H{
					"msg": "服务器错误3",
				})
				return
			}
			if user.Password == password {

				log.Println(strconv.Itoa(user.Id))
				c.SetCookie("user_id", strconv.Itoa(user.Id), 60*60*24, "/", "localhost", false, true)
				c.SetCookie("phoneoremail", phoneoremail, 60*60*24, "/", "localhost", false, true)

				c.JSON(200, gin.H{
					"msg":     "登录成功",
					"user_id": user.Id,
				})
				return
			} else {
				c.JSON(200, gin.H{
					"msg": "密码错误",
				})
				return
			}
		}
	})
	Router.GET("/sign_out", func(c *gin.Context) {
		c.SetCookie("phoneoremail", "1", 1, "/", "localhost", false, true)

		c.JSON(200, gin.H{
			"msg": "退出成功",
		})
		return

	})

	Router.POST("/getnotice", func(c *gin.Context) {
		data, _ := c.GetRawData()
		var m map[string]string
		_ = json.Unmarshal(data, &m)
		user_id := m["user_id"]
		db, err := dao.OpenDb()
		if err != nil {
			log.Println(err)
			return
		}
		myid, err := c.Cookie("user_id")
		if err != nil {
			log.Println(err)
			return
		}
		var follow int
		var noticeornot = 1
		row := db.QueryRow("select follow from attention  where followed=? and  follow=?", user_id, myid)
		err = row.Scan(&follow)
		if err != nil {
			noticeornot = 0
		}
		log.Println(noticeornot)
		c.JSON(200, gin.H{
			"noticeornot": noticeornot,
		})
	})

	Router.POST("/getuserid", func(c *gin.Context) {
		db, err := dao.OpenDb()
		if err != nil {
			log.Println(err)
			return
		}
		data, _ := c.GetRawData()
		var m map[string]string
		_ = json.Unmarshal(data, &m)
		phoneoremail := m["phoneoremail"]
		row := db.QueryRow("select id from user where phoneoremail=?", phoneoremail)
		var id int
		err = row.Scan(&id)
		if err != nil {
			log.Println(err)
			return
		}
		c.JSON(200, gin.H{
			"user_id": id,
		})
	})

	//获取个人信息
	Router.GET("/profile", func(c *gin.Context) {
		db, err := dao.OpenDb()
		if err != nil {
			log.Println(err)
			c.JSON(200, gin.H{
				"msg": "数据库错误1",
			})
			return
		}
		phoneoremail, err := c.Cookie("phoneoremail")
		if err != nil {
			log.Println(err)
			c.JSON(200, gin.H{
				"msg": "获取cookie失败",
			})
			return
		}
		var userinfo model.Userinfo
		row := db.QueryRow("select * from userinfo where phoneoremail = ?", phoneoremail)
		err = row.Scan(&userinfo.Name, &userinfo.Position, &userinfo.Company, &userinfo.Web, &userinfo.Introduce, &userinfo.Phoneoremail, &userinfo.Cover)
		if err != nil {
			log.Println(err)
			c.JSON(200, gin.H{
				"msg": "数据库错误2",
			})
			return
		}

		c.JSON(200, gin.H{
			"userinfo": userinfo,
		})
	})

	//获取账号信息
	Router.GET("/account", func(c *gin.Context) {
		db, err := dao.OpenDb()
		if err != nil {
			log.Println(err)
			c.JSON(200, gin.H{
				"msg": "数据库错误1",
			})
			return
		}
		phoneoremail, err := c.Cookie("phoneoremail")
		if err != nil {
			log.Println(err)
			c.JSON(200, gin.H{
				"msg": "获取cookie失败",
			})
			return
		}
		var accountinfo model.Accountinfo
		row := db.QueryRow("select * from accountinfo where phoneoremail = ?", phoneoremail)
		err = row.Scan(&accountinfo.Phone, &accountinfo.Weixin, &accountinfo.Xinlang, &accountinfo.Github, &accountinfo.Phoneoremail)
		if err != nil {
			log.Println(err)
			c.JSON(200, gin.H{
				"msg": "数据库错误2",
			})
			return
		}
		c.JSON(200, gin.H{
			"accountinfo": accountinfo,
		})
	})

	User := Router.Group("/user", handler.Auth())
	{ //个人详情页，动态
		User.GET("", func(c *gin.Context) {
			c.HTML(200, "person3.html", gin.H{})
		})
		User.POST("", func(c *gin.Context) {
			data, _ := c.GetRawData()
			var m map[string]string
			_ = json.Unmarshal(data, &m)
			id := m["user_id"]
			db, err := dao.OpenDb()
			if err != nil {
				log.Println(err)
				return
			}
			rows, err := db.Query("select * from article where author_id=? ", id)
			if err != nil {
				log.Println(err)
				return
			}

			articles := make([]model.Article, 0)
			//  迭代查询获取数据  必须调用
			for rows.Next() {
				// row.scan 必须按照先后顺序 &获取数据
				var article model.Article
				err := rows.Scan(&article.Article_title, &article.Article_content, &article.Date, &article.Category, &article.Label, &article.Column, &article.Like, &article.Id, &article.Author, &article.Author_id, &article.View, &article.Postorboil, &article.Cover, &article.Comment, &article.Collection)
				if err != nil {
					log.Println(err)
					return
				}

				articles = append(articles, article)
			}

			c.JSON(200, gin.H{
				"msg": articles,
			})

		})

		//个人详情页，文章
		User.GET("/posts", func(c *gin.Context) {
			c.HTML(200, "person-home-article6.html", gin.H{})
		})
		User.POST("/posts", func(c *gin.Context) {
			data, _ := c.GetRawData()
			var m map[string]string
			_ = json.Unmarshal(data, &m)
			id := m["user_id"]

			db, err := dao.OpenDb()
			if err != nil {
				log.Println(err)
				return
			}
			rows, err := db.Query("select * from article where author_id=? and postorboil=?", id, 0)
			if err != nil {
				log.Println(err)
				return
			}

			articles := make([]model.Article, 0)
			//  迭代查询获取数据  必须调用
			for rows.Next() {
				// row.scan 必须按照先后顺序 &获取数据
				var article model.Article
				err := rows.Scan(&article.Article_title, &article.Article_content, &article.Date, &article.Category, &article.Label, &article.Column, &article.Like, &article.Id, &article.Author, &article.Author_id, &article.View, &article.Postorboil, &article.Cover, &article.Comment, &article.Collection)
				if err != nil {
					log.Println(err)
					return
				}

				articles = append(articles, article)
			}

			c.JSON(200, gin.H{
				"msg": articles,
			})

		})
		User.GET("/columns", func(c *gin.Context) {
			c.HTML(200, "person-home-column4.html", gin.H{})
		})

		User.GET("/pins", func(c *gin.Context) {
			c.HTML(200, "person-home-boil4.html", gin.H{})
		})
		//沸点展示
		User.POST("/pins", func(c *gin.Context) {
			data, _ := c.GetRawData()
			var m map[string]string
			_ = json.Unmarshal(data, &m)
			user_id := m["user_id"]
			log.Println(user_id)
			db, err := dao.OpenDb()
			if err != nil {
				log.Println(err)
				return
			}
			rows, err := db.Query("select * from article where author_id=? and postorboil=?", user_id, 1)
			if err != nil {
				log.Println(err)
				return
			}

			articles := make([]model.Article, 0)
			//  迭代查询获取数据  必须调用
			for rows.Next() {

				var article model.Article
				err = rows.Scan(&article.Article_title, &article.Article_content, &article.Date, &article.Category, &article.Label, &article.Column, &article.Like, &article.Id, &article.Author, &article.Author_id, &article.View, &article.Postorboil, &article.Cover, &article.Comment, &article.Collection)
				if err != nil {
					log.Println(err)
					return
				}

				articles = append(articles, article)
			}
			log.Println(articles)
			c.JSON(200, gin.H{
				"msg": articles,
			})

		})
		User.GET("/collections", func(c *gin.Context) {
			c.HTML(200, "person-home-collect4.html", gin.H{})
		})
		//收藏集
		User.POST("/collections", func(c *gin.Context) {
			data, _ := c.GetRawData()
			var m map[string]string
			_ = json.Unmarshal(data, &m)
			user_id := m["user_id"]
			db, err := dao.OpenDb()
			if err != nil {
				log.Println(err)
				return
			}
			rows, err := db.Query("select * from article where id in (select article_id from collections where collector_id=?)", user_id)
			articles := make([]model.Article, 0)
			for rows.Next() {
				var article model.Article
				err = rows.Scan(&article.Article_title, &article.Article_content, &article.Date, &article.Category, &article.Label, &article.Column, &article.Like, &article.Id, &article.Author, &article.Author_id, &article.View, &article.Postorboil, &article.Cover, &article.Comment, &article.Collection)
				if err != nil {
					log.Println(err)
					return
				}
				articles = append(articles, article)
			}
			log.Println(articles)
			c.JSON(200, gin.H{
				"msg": articles,
			})
		})
		//关注的用户
		User.GET("/following", func(c *gin.Context) {
			c.HTML(200, "person-home-follow5.html", gin.H{})
		})
		User.POST("/following", func(c *gin.Context) {
			data, _ := c.GetRawData()
			var m map[string]string
			_ = json.Unmarshal(data, &m)
			user_id := m["user_id"]
			db, err := dao.OpenDb()
			if err != nil {
				log.Println(err)
				return
			}
			rows, err := db.Query("select followed from attention where follow=?", user_id)

			if err != nil {
				log.Println(err)
				return
			}
			userinfos := make([]model.Userinfo, 0)
			for rows.Next() {
				var followed int
				var phoneoremail string
				var userinfo model.Userinfo
				err := rows.Scan(&followed)
				if err != nil {
					log.Println(err)
					return
				}
				row := db.QueryRow("select phoneoremail from user where id=?", followed)
				err = row.Scan(&phoneoremail)
				if err != nil {
					log.Println(err)
					return
				}
				row = db.QueryRow("select * from userinfo where phoneoremail=?", phoneoremail)
				err = row.Scan(&userinfo.Name, &userinfo.Position, &userinfo.Company, &userinfo.Web, &userinfo.Introduce, &userinfo.Phoneoremail, &userinfo.Cover)
				if err != nil {
					log.Println(err)
					return
				}
				userinfos = append(userinfos, userinfo)
			}
			c.JSON(200, gin.H{
				"msg": userinfos,
			})
		})
		User.GET("/followers", func(c *gin.Context) {
			c.HTML(200, "person-home-befollowed3.html", gin.H{})
		})
		//关注者
		User.POST("/followers", func(c *gin.Context) {
			data, _ := c.GetRawData()
			var m map[string]string
			_ = json.Unmarshal(data, &m)
			user_id := m["user_id"]
			db, err := dao.OpenDb()
			if err != nil {
				log.Println(err)
				return
			}
			rows, err := db.Query("select follow from attention where followed=?", user_id)

			if err != nil {
				log.Println(err)
				return
			}
			userinfos := make([]model.Userinfo, 0)
			for rows.Next() {
				var follow int
				var phoneoremail string
				var userinfo model.Userinfo
				err := rows.Scan(&follow)
				if err != nil {
					log.Println(err)
					return
				}
				row := db.QueryRow("select phoneoremail from user where id=?", follow)
				err = row.Scan(&phoneoremail)
				if err != nil {
					log.Println(err)
					return
				}
				row = db.QueryRow("select * from userinfo where phoneoremail=?", phoneoremail)
				err = row.Scan(&userinfo.Name, &userinfo.Position, &userinfo.Company, &userinfo.Web, &userinfo.Introduce, &userinfo.Phoneoremail, &userinfo.Cover)
				if err != nil {
					log.Println(err)
					return
				}
				userinfos = append(userinfos, userinfo)
			}
			c.JSON(200, gin.H{
				"msg": userinfos,
			})
		})
		User.GET("/likes", func(c *gin.Context) {
			c.HTML(200, "person-home-like-article5.html", gin.H{})
		})
		//点赞的文章
		User.POST("/likes", func(c *gin.Context) {
			data, _ := c.GetRawData()
			var m map[string]string
			_ = json.Unmarshal(data, &m)
			user_id := m["user_id"]
			db, err := dao.OpenDb()
			if err != nil {
				log.Println(err)
				return
			}
			rows, err := db.Query("select * from article where id in (select liked_id from `like` where status='点赞' and type=0 and liker_id=?)", user_id)
			articles := make([]model.Article, 0)
			for rows.Next() {
				var article model.Article
				err := rows.Scan(&article.Article_title, &article.Article_content, &article.Date, &article.Category, &article.Label, &article.Column, &article.Like, &article.Id, &article.Author, &article.Author_id, &article.View, &article.Postorboil, &article.Cover, &article.Comment, &article.Collection)
				if err != nil {
					log.Println(err)
					return
				}

				articles = append(articles, article)
			}
			log.Println(articles)
			c.JSON(200, gin.H{
				"msg": articles,
			})

		})
		User.GET("/praise", func(c *gin.Context) {
			c.HTML(200, "person-home-like-boil5.html", gin.H{})
		})
		//点赞的沸点
		User.POST("/praise", func(c *gin.Context) {
			data, _ := c.GetRawData()
			var m map[string]string
			_ = json.Unmarshal(data, &m)
			user_id := m["user_id"]
			db, err := dao.OpenDb()
			if err != nil {
				log.Println(err)
				return
			}
			rows, err := db.Query("select liked_id from `like` where liker_id=? and type=?", user_id, 1)

			if err != nil {
				log.Println(err)
				return
			}
			userinfos := make([]model.Userinfo, 0)
			for rows.Next() {
				var liked_id int
				var phoneoremail string
				var userinfo model.Userinfo
				err := rows.Scan(&liked_id)
				if err != nil {
					log.Println(err)
					return
				}
				row := db.QueryRow("select phoneoremail from user where id=?", liked_id)
				err = row.Scan(&phoneoremail)
				if err != nil {
					log.Println(err)
					return
				}
				row = db.QueryRow("select * from userinfo where phoneoremail=?", phoneoremail)
				err = row.Scan(&userinfo.Name, &userinfo.Position, &userinfo.Company, &userinfo.Web, &userinfo.Introduce, &userinfo.Phoneoremail, &userinfo.Cover)
				if err != nil {
					log.Println(err)
					return
				}
				userinfos = append(userinfos, userinfo)
			}
			c.JSON(200, gin.H{
				"msg": userinfos,
			})

		})
		//设置
		Settings := User.Group("/settings")
		{

			//设置账号信息
			Settings.GET("/account", func(c *gin.Context) {
				c.HTML(200, "edit-account.html", gin.H{})
			})
			Settings.POST("/account", func(c *gin.Context) {
				data, _ := c.GetRawData()
				var m map[string]any
				_ = json.Unmarshal(data, &m)
				phone := m["phone"]
				weixin := m["weixin"]
				xinlang := m["xinlang"]
				github := m["github"]
				phoneoremail, err := c.Cookie("phoneoremail")

				if phone == nil {
					c.JSON(403, gin.H{
						"msg": "phone不能为空",
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
				result, err := Db.Exec("update accountinfo set phone=?,weixin=?,xinlang=?,github=? where phoneoremail=?", phone, weixin, xinlang, github, phoneoremail)
				log.Println(result)
				if err != nil {
					log.Println(err)
					c.JSON(500, gin.H{
						"msg": "服务器错误无法更新",
					})
					return
				}
				c.JSON(200, gin.H{
					"msg": "修改成功",
				})
			})
		}

	}
	//设置个人信息
	Router.GET("/user/settings/profile", func(c *gin.Context) {
		c.HTML(200, "edit-person.html", gin.H{})
	})
	Router.POST("/user/settings/profile", func(c *gin.Context) {

		data, _ := c.GetRawData()
		var m map[string]string
		_ = json.Unmarshal(data, &m)
		name := m["name"]
		position := m["position"]
		company := m["company"]
		web := m["web"]
		intruduce := m["introduce"]
		cover := m["cover"]
		phoneoremail, _ := c.Cookie("phoneoremail")
		if name == "" {
			c.JSON(403, gin.H{
				"msg": "name不能为空",
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
		result, err := Db.Exec("update userinfo set name=?,position=?,company=?,web=?,intruduce=?,cover=? where phoneoremail=?", name, position, company, web, intruduce, cover, phoneoremail)
		log.Println(result)
		if err != nil {
			log.Println(err)
			c.JSON(500, gin.H{
				"msg": "服务器错误无法更新",
			})
			return
		}
		c.JSON(200, gin.H{
			"msg": "修改成功",
		})
	})
	Router.GET("/published", func(c *gin.Context) {
		c.HTML(200, "publish_success2.html", gin.H{})
	})

	//创作者中心
	Creator := Router.Group("/creator", handler.Auth())
	{
		//创作者中心主页
		Creator.GET("/home", func(c *gin.Context) {
			c.HTML(200, "creator-home.html", gin.H{})
		})
		Creator.POST("/home", func(c *gin.Context) {
			data, _ := c.GetRawData()
			var m map[string]any
			_ = json.Unmarshal(data, &m)

			user_id, err := c.Cookie("user_id")
			if err != nil {
				log.Println(err)
				return
			}
			db, err := dao.OpenDb()
			if err != nil {
				log.Println(err)
				return
			}
			var followers_number, article_view_number, like_number, comment_number, collection_number = 0, 0, 0, 0, 0
			rows, err := db.Query("select date,status from attention where followed=?", user_id)
			if err != nil {
				log.Println(err)
				followers_number = 0
			}
			for rows.Next() {
				followers_number++
			}
			var view, like, comments, collection int
			rows, err = db.Query("select view,`like`,comments,collection from article where author_id=?", user_id)

			if err != nil {
				log.Println(err)
				return
			}
			for rows.Next() {
				err = rows.Scan(&view, &like, &comments, &collection)
				if err != nil {
					log.Println(err)
					return
				}
				article_view_number, like_number, comment_number, collection_number = article_view_number+view, like_number+like, comment_number+comments, collection_number+collection
			}

			c.JSON(200, gin.H{
				"followers_number":    followers_number,
				"article_view_number": article_view_number,
				"like_number":         like_number,
				"comment_number":      comment_number,
				"collection_number":   collection_number,
			})
		})
		Content := Creator.Group("/content")
		{
			//内容管理

			Article := Content.Group("/article")
			{
				//文章管理
				Article.GET("/essays", func(c *gin.Context) {
					c.HTML(200, "article_charge2.html", gin.H{})
				})
				Article.POST("/essays", func(c *gin.Context) {
					db, err := dao.OpenDb()
					if err != nil {
						log.Println(err)
						c.JSON(200, gin.H{
							"msg": "数据库错误",
						})
						return
					}
					phoneoremail, err := c.Cookie("phoneoremail")

					if err != nil {
						log.Println(err)
						c.JSON(200, gin.H{
							"msg": "cookie错误",
						})
						return
					}
					var user_id int
					row := db.QueryRow("select id from user where phoneoremail=?", phoneoremail)

					row.Scan(&user_id)
					rows, err := db.Query("select * from article where postorboil=? and author_id=?", 0, user_id)
					log.Println(rows)
					articles := make([]model.Article, 0)
					for rows.Next() {
						var article model.Article
						err := rows.Scan(&article.Article_title, &article.Article_content, &article.Date, &article.Category, &article.Label, &article.Column, &article.Like, &article.Id, &article.Author, &article.Author_id, &article.View, &article.Postorboil, &article.Cover, &article.Comment, &article.Collection)
						if err != nil {
							log.Println(err)
							return
						}
						articles = append(articles, article)
					}

					c.JSON(200, gin.H{
						"msg": articles,
					})
				})

			}
			//专栏管理
			Content.GET("/column", func(c *gin.Context) {
				c.HTML(200, "column_charge3.html", gin.H{})
			})
			Content.POST("/column", func(c *gin.Context) {
				db, err := dao.OpenDb()
				if err != nil {
					log.Println(err)
					c.JSON(200, gin.H{
						"msg": "数据库错误",
					})
					return
				}
				phoneoremail, err := c.Cookie("phoneoremail")

				if err != nil {
					log.Println(err)
					c.JSON(200, gin.H{
						"msg": "cookie错误",
					})
					return
				}

				rows, err := db.Query("select * from column where phoneoremail=?", phoneoremail)
				log.Println(rows)
				columns := make([]model.Column, 0)
				for rows.Next() {
					var column model.Column
					err := rows.Scan(&column.Column_title, &column.Column_intruduction, &column.Phoneoremail, &column.Id, &column.Cover)
					if err != nil {
						log.Println(err)
						return
					}
					columns = append(columns, column)
				}

				c.JSON(200, gin.H{
					"msg": columns,
				})
			})
			Content.GET("/boil", func(c *gin.Context) {
				c.HTML(200, "boil_charge1.html", gin.H{})
			})

			Data := Creator.Group("/data")
			{
				Data.GET("/content/article", func(c *gin.Context) {
					c.HTML(200, "article_data.html", gin.H{})
				})
				Data.POST("/content/article", func(c *gin.Context) {
					user_id, err := c.Cookie("user_id")
					if err != nil {
						log.Println(err)
						return
					}
					db, err := dao.OpenDb()
					if err != nil {
						log.Println(err)
						return
					}
					rows, err := db.Query("select view,comments,collection,`like` from article where author_id=? and postorboil= 0", user_id)
					if err != nil {
						log.Println(err)
						return
					}
					var all_article, article_view_number, like_number, comment_number, collection_number = 0, 0, 0, 0, 0
					var view, comments, collection, like int
					for rows.Next() {
						err := rows.Scan(&view, &comments, &collection, &like)
						if err != nil {
							log.Println(err)
							return
						}
						all_article, article_view_number, like_number, comment_number, collection_number = all_article+1, article_view_number+view, like_number+like, comment_number+comments, collection_number+collection

					}
					c.JSON(200, gin.H{
						"all_article":         all_article,
						"article_view_number": article_view_number,
						"like_number":         like_number,
						"comment_number":      comment_number,
						"collection_number":   collection_number,
					})
				})
				//专栏数据
				Data.GET("/content/column", func(c *gin.Context) {
					c.HTML(200, "column_data.html", gin.H{})
				})
				Data.POST("/content/column", func(c *gin.Context) {
					phoneoremail, err := c.Cookie("phoneoremail")
					if err != nil {
						log.Println(err)
						return
					}
					db, err := dao.OpenDb()
					if err != nil {
						log.Println(err)
						return
					}
					rows, err := db.Query("select id from `column` where phoneoremail=?", phoneoremail)
					if err != nil {
						log.Println(err)
						c.JSON(200, gin.H{
							"column_number": 0,
						})
						return
					}
					var column_number = 0
					for rows.Next() {
						column_number++
					}
					c.JSON(200, gin.H{
						"column_number": column_number,
					})
				})
				//沸点数据

				Data.GET("/content/pin", func(c *gin.Context) {
					c.HTML(200, "boil_data.html", gin.H{})
				})
				Data.POST("/content/pin", func(c *gin.Context) {
					user_id, err := c.Cookie("user_id")
					if err != nil {
						log.Println(err)
						return
					}
					db, err := dao.OpenDb()
					if err != nil {
						log.Println(err)
						return
					}
					rows, err := db.Query("select comments,`like` from article where author_id=? and postorboil= 1", user_id)
					if err != nil {
						log.Println(err)
						return
					}
					var pins_number, pins_like_number, pins_comment_number = 0, 0, 0
					var comment, like int
					for rows.Next() {
						pins_number++
						err = rows.Scan(&comment, &like)
						if err != nil {
							log.Println(err)
							return
						}
						pins_comment_number, pins_like_number = pins_comment_number+comment, pins_like_number+like
					}
					c.JSON(200, gin.H{
						"pins_number":         pins_number,
						"pins_like_number":    pins_like_number,
						"pins_comment_number": pins_comment_number,
					})
				})
				Data.GET("/follow/data", func(c *gin.Context) {
					c.HTML(200, "follower_data.html", gin.H{})
				})
				Data.POST("/follow/data", func(c *gin.Context) {
					data, _ := c.GetRawData()
					var m map[string]any
					_ = json.Unmarshal(data, &m)
					date := m["time"]

					user_id, err := c.Cookie("user_id")
					if err != nil {
						log.Println(err)
						return
					}
					db, err := dao.OpenDb()
					if err != nil {
						log.Println(err)
						return
					}
					var follower_number, new_follower_number, cancel_number, added_number = 0, 0, 0, 0
					rows, err := db.Query("select date,status from attention where followed=?", user_id)
					if err != nil {
						log.Println(err)
						follower_number = 0
					}
					for rows.Next() {
						follower_number++
					}
					rows, err = db.Query("select status from attention where followed=? and FROM_UNIXTIME(?,'%Y-%m-%d')=DATE_SUB(CURDATE(), INTERVAL 1 DAY)", user_id, date)
					if err != nil {
						log.Println(err)
					}
					var status string
					for rows.Next() {
						rows.Scan(&status)
						if status == "关注" {
							new_follower_number++
						} else {
							cancel_number++
						}
					}
					added_number = new_follower_number - cancel_number
					c.JSON(200, gin.H{
						"follower_number":     follower_number,
						"new_follower_number": new_follower_number,
						"cancel_number":       cancel_number,
						"added_number":        added_number,
					})
				})

				Data.GET("/follower/list", func(c *gin.Context) {
					c.HTML(200, "follower_list.html", gin.H{})
				})
			}

		}

	}

	//创建专栏
	Router.POST("/create_column", func(c *gin.Context) {
		data, _ := c.GetRawData()
		var m map[string]any
		_ = json.Unmarshal(data, &m)

		column_title := m["column_title"]
		column_intruduction := m["column_intruduction"]
		phoneoremail, err := c.Cookie("phoneoremail")
		if err != nil {
			log.Println(err)
			return
		}
		cover := m["cover"]

		time := m["time"]
		db, err := dao.OpenDb()
		if err != nil {
			log.Println(err)
			return
		}
		result, err := db.Exec("insert `column` (column_title, column_intruduction, phoneoremail, cover,column_number,time) VALUE (?,?,?,?,?,?)", column_title, column_intruduction, phoneoremail, cover, 0, time)
		log.Println(result)
		if err != nil {
			log.Println(err)
			return
		}

		row := db.QueryRow("select id from `column` where column_title=?", column_title)
		var id int
		err = row.Scan(&id)
		if err != nil {
			log.Println(err)
			return
		}
		c.JSON(200, gin.H{
			"msg": "创建成功",
			"id":  id,
		})
	})
	//评论
	Router.POST("/comment", func(c *gin.Context) {
		data, _ := c.GetRawData()
		var m map[string]any
		_ = json.Unmarshal(data, &m)

		projectid := m["projectid"]
		parentid := m["parentid"]
		content := m["content"]
		date := m["time"]
		var user_id int
		db, err := dao.OpenDb()
		if err != nil {
			log.Println(err)
			c.JSON(200, gin.H{
				"msg": "数据库错误",
			})
			return
		}
		phoneoremail, err := c.Cookie("phoneoremail")
		if err != nil {
			log.Println(err)
			c.JSON(200, gin.H{
				"msg": "cookie错误",
			})
			return
		}
		row := db.QueryRow("select id from user where phoneoremail=?", phoneoremail)

		row.Scan(&user_id)
		if projectid == nil {
			log.Println("项目不得为空")
			c.JSON(200, gin.H{
				"msg": "项目不得为空",
			})
			return
		}

		result, err := db.Exec("update article set comments =comments+1 where id=?", projectid)
		log.Println(result)
		if err != nil {
			log.Println(err)
			c.JSON(200, gin.H{
				"msg": "评论失败",
			})
			return
		}
		result, err = db.Exec("insert into comment (id,projectid,userid,content,parentid,time) value (?,?,?,?,?,?)", nil, projectid, user_id, content, parentid, date)
		log.Println(result)
		if err != nil {
			log.Println(err)
			c.JSON(200, gin.H{
				"msg": "评论失败",
			})
			return
		}
		var id int
		row = db.QueryRow("select id from comment where content=? and userid=?", content, user_id)
		err = row.Scan(&id)
		if err != nil {
			log.Println(err)
			return
		}
		c.JSON(200, gin.H{
			"msg": "评论成功",
			"id":  id,
		})
	})

	//获取评论
	Router.POST("/getcomment", func(c *gin.Context) {
		data, _ := c.GetRawData()
		var m map[string]any
		_ = json.Unmarshal(data, &m)
		var i = 0
		db, err := dao.OpenDb()
		if err != nil {
			log.Println(err)
			c.JSON(200, gin.H{
				"msg": "数据库错误",
			})
			return
		}
		var user_id int
		project_id := m["id"]
		phoneoremail, _ := c.Cookie("phoneoremail")
		row := db.QueryRow("select id from user where phoneoremail=?", phoneoremail)

		err = row.Scan(&user_id)

		if err != nil {

			log.Println(err)
			c.JSON(200, gin.H{
				"msg": "数据库错误",
			})
			return
		}
		rows, _ := db.Query("select * from comment where projectid=? and parentid=?", project_id, 0)
		comments := make([]model.Comment, 0)

		for rows.Next() {
			var comment model.Comment
			err = rows.Scan(&comment.ID, &comment.ProjectId, &comment.UserId, &comment.Content, &comment.ParentId, &comment.Time)
			if err != nil {
				log.Println("hhh")
				log.Println(err)
				return
			}
			comments = append(comments, comment)
			rows2, _ := db.Query("select * from comment where projectid=? and parentid=?", project_id, comment.ID)
			i++
			for rows2.Next() {
				var comment2 model.Comment
				err = rows2.Scan(&comment2.ID, &comment2.ProjectId, &comment2.UserId, &comment2.Content, &comment2.ParentId, &comment2.Time)
				comments = append(comments, comment2)
				i++
			}
		}
		c.JSON(200, gin.H{
			"msg":    comments,
			"length": i,
		})

	})
	//编辑器
	Router.GET("/write", func(c *gin.Context) {
		c.HTML(200, "write3.html", gin.H{})
	})
	Router.POST("/editor/drafts/new", func(c *gin.Context) {
		db, err := dao.OpenDb()
		if err != nil {
			log.Println(err)
			c.JSON(200, gin.H{
				"msg": "数据库错误",
			})
			return
		}
		data, _ := c.GetRawData()
		var m map[string]any
		_ = json.Unmarshal(data, &m)
		title := m["title"]
		content := m["content"]
		column := m["column"]
		category := m["category"]
		label := m["label"]
		cover := m["cover"]
		time := m["date"]
		phoneoremail, err := c.Cookie("phoneoremail")
		row := db.QueryRow("select name from userinfo where phoneoremail=?", phoneoremail)
		var author string
		row.Scan(&author)

		author_id, err := c.Cookie("user_id")

		if err != nil {
			log.Println(err)
			return
		}

		result, err := db.Exec("insert into article (article_title,article_content,date,category,label,`column`,`like`,id,author,author_id,view,comments,postorboil,cover,collection) value (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)", title, content, time, category, label, column, 0, nil, author, author_id, 0, 0, 0, cover, 0)
		log.Println(result)
		if err != nil {
			log.Println(err)
			return
		}

		c.JSON(200, gin.H{
			"msg": "创建成功",
		})
		log.Println("ok")
	})

	//获取用户信息
	Router.POST("/getuserinfo", func(c *gin.Context) {
		data, _ := c.GetRawData()
		var m map[string]any
		_ = json.Unmarshal(data, &m)
		user_id := m["user_id"]
		db, _ := dao.OpenDb()
		log.Println(user_id)
		row := db.QueryRow("select phoneoremail from user where id =?", user_id)

		var phoneoremail string
		err := row.Scan(&phoneoremail)

		if err != nil {
			log.Println(1)
			log.Println(err)
			return
		}
		row = db.QueryRow("select * from userinfo where phoneoremail= ?", phoneoremail)
		var userinfo model.Userinfo
		err = row.Scan(&userinfo.Name, &userinfo.Position, &userinfo.Company, &userinfo.Web, &userinfo.Introduce, &userinfo.Phoneoremail, &userinfo.Cover)
		if err != nil {
			log.Println(2)
			log.Println(err)
			return
		}

		c.JSON(200, gin.H{
			"msg": userinfo,
		})
	})

	//获取评论人信息
	Router.POST("/commentorinfo", func(c *gin.Context) {
		data, _ := c.GetRawData()
		var m map[string]any
		_ = json.Unmarshal(data, &m)
		id := m["parentid"]
		db, _ := dao.OpenDb()
		row := db.QueryRow("select userid from comment where parentid =?", id)
		var userid int
		err := row.Scan(&userid)
		if err != nil {
			log.Println("1")
			log.Println(err)
			return
		}
		row = db.QueryRow("select phoneoremail from user where id =?", userid)
		var phoneoremail string
		row.Scan(&phoneoremail)
		row = db.QueryRow("select * from userinfo where phoneoremail=?", phoneoremail)
		var userinfo model.Userinfo
		err = row.Scan(&userinfo.Name, &userinfo.Position, &userinfo.Company, &userinfo.Web, &userinfo.Introduce, &userinfo.Phoneoremail, &userinfo.Cover)
		if err != nil {
			log.Println("2")
			log.Println(err)
			return
		}

		c.JSON(200, gin.H{
			"msg": userinfo,
		})
	})
	//获取所有专栏
	Router.POST("/getcolumns", func(c *gin.Context) {
		data, _ := c.GetRawData()
		var m map[string]any
		_ = json.Unmarshal(data, &m)
		user_id := m["user_id"]
		db, err := dao.OpenDb()
		if err != nil {
			log.Println(err)
			return
		}

		rows, err := db.Query("select * from `column` where phoneoremail in (select phoneoremail from  user where user.id=?)", user_id)
		if err != nil {
			log.Println(err)
			return
		}
		columns := make([]model.Column, 0)
		for rows.Next() {
			var column model.Column
			err := rows.Scan(&column.Column_title, &column.Column_intruduction, &column.Phoneoremail, &column.Id, &column.Cover, &column.Column_number, &column.Time)
			if err != nil {
				log.Println(err)
				return
			}
			columns = append(columns, column)
		}
		c.JSON(200, gin.H{
			"msg": columns,
		})
	})
	//文章详细页
	Router.GET("/content", func(c *gin.Context) {
		c.HTML(200, "article1.html", gin.H{})
	})
	Router.POST("/content", func(c *gin.Context) {
		data, _ := c.GetRawData()
		var m map[string]any
		_ = json.Unmarshal(data, &m)
		var likeornot int
		id := m["id"]
		log.Println(id)
		db, _ := dao.OpenDb()
		row := db.QueryRow("select * from article where id= ?", id)
		var article model.Article
		err := row.Scan(&article.Article_title, &article.Article_content, &article.Date, &article.Category, &article.Label, &article.Column, &article.Like, &article.Id, &article.Author, &article.Author_id, &article.View, &article.Postorboil, &article.Cover, &article.Comment, &article.Collection)
		if err != nil {
			log.Println(err)
			return
		}
		phoneoremail, err := c.Cookie("phoneoremail")
		var status string
		row = db.QueryRow("select status from `like` where liker_name=? and liked_id= ?", phoneoremail, id)
		err = row.Scan(&status)
		if status == "点赞成功" {
			likeornot = 1
		} else {
			likeornot = 0
		}
		user_id, err := c.Cookie("user_id")
		if err != nil {
			log.Println(err)
			return
		}
		var collectornot int
		var date string
		db.Exec("insert into view (viewer_phoneoremail, article_id) VALUE (?,?)", phoneoremail, id)
		db.Exec("update article set view =? where id =?", article.View+1, id)
		row = db.QueryRow("select date from collections where collector_id=? and article_id=?", user_id, id)
		err = row.Scan(&date)
		if err != nil {
			collectornot = 0
		} else {
			collectornot = 1
		}
		log.Println(collectornot)
		c.JSON(200, gin.H{
			"msg":          article,
			"likeornot":    likeornot,
			"collectornot": collectornot,
		})
	})
	//点赞
	Router.POST("/like", func(c *gin.Context) {
		data, _ := c.GetRawData()
		var m map[string]any
		_ = json.Unmarshal(data, &m)
		id := m["id"]
		db, _ := dao.OpenDb()
		row := db.QueryRow("select postorboil,`like` from  article where id=?", id)
		var postorboil, like int
		err := row.Scan(&postorboil, &like)
		if err != nil {
			log.Println(err)
			return
		}
		phoneoremail, _ := c.Cookie("phoneoremail")
		user_id, _ := c.Cookie("user_id")
		log.Println(user_id)
		row = db.QueryRow("select status from `like` where liked_id=? and liker_id=?", id, user_id)
		var status string
		err = row.Scan(&status)

		if err != nil {
			db.Exec("insert into `like` (type, liker_id, liker_name, liked_id, status) VALUE (?,?,?,?,?)", 0, user_id, phoneoremail, id, "点赞")
			db.Exec("update article set `like` =? where id=?", like+1, id)
			c.JSON(200, gin.H{
				"msg": "点赞成功",
			})
			return
		}
		if status == "取消点赞" {
			db.Exec("update `like` set status =? where liker_id=? and liked_id=?", "点赞", user_id, id)
			db.Exec("update article set `like` =? where id=?", like+1, id)
			c.JSON(200, gin.H{
				"msg": "点赞成功",
			})
			return
		}
		db.Exec("update `like` set status =? where liker_id=? and liked_id=?", "取消点赞", user_id, id)
		db.Exec("update article set `like` =? where id=?", like-1, id)
		c.JSON(200, gin.H{
			"msg": "取消点赞",
		})
	})
	//首页
	Router.GET("/home", func(c *gin.Context) {
		c.HTML(200, "home2.html", gin.H{})
	})

	Router.POST("/home", func(c *gin.Context) {

		data, _ := c.GetRawData()
		var m map[string]any
		_ = json.Unmarshal(data, &m)
		date := m["date"]
		status := c.DefaultQuery("status", "")

		db, err := dao.OpenDb()
		if err != nil {
			log.Println(err)
			return
		}
		if status == "" {

			rows, err := db.Query("select * from article where postorboil=?", 0)
			if err != nil {
				log.Println(err)
				return
			}

			articles := make([]model.Article, 0)
			//  迭代查询获取数据  必须调用
			for rows.Next() {
				// row.scan 必须按照先后顺序 &获取数据
				var article model.Article
				err := rows.Scan(&article.Article_title, &article.Article_content, &article.Date, &article.Category, &article.Label, &article.Column, &article.Like, &article.Id, &article.Author, &article.Author_id, &article.View, &article.Postorboil, &article.Cover, &article.Comment, &article.Collection)
				if err != nil {
					log.Println(err)
					return
				}

				articles = append(articles, article)
			}
			var signornot = 1
			row := db.QueryRow("select date from sign where date(?)=DATE_SUB(CURDATE(), INTERVAL 0 DAY)", date)
			err = row.Scan(&date)
			if err != nil {
				signornot = 0
			}
			c.JSON(200, gin.H{
				"msg":       articles,
				"signornot": signornot,
			})
		} else if status == "new" {
			rows, err := db.Query("select * from article where postorboil= 0 order by date desc  ")
			if err != nil {
				log.Println(err)
				return
			}

			articles := make([]model.Article, 0)
			//  迭代查询获取数据  必须调用
			for rows.Next() {
				// row.scan 必须按照先后顺序 &获取数据
				var article model.Article
				err := rows.Scan(&article.Article_title, &article.Article_content, &article.Date, &article.Category, &article.Label, &article.Column, &article.Like, &article.Id, &article.Author, &article.Author_id, &article.View, &article.Postorboil, &article.Cover, &article.Comment, &article.Collection)
				if err != nil {
					log.Println(err)
					return
				}

				articles = append(articles, article)
			}
			var signornot = 1
			row := db.QueryRow("select date from sign where date(?)=DATE_SUB(CURDATE(), INTERVAL 0 DAY)", date)
			err = row.Scan(&date)
			if err != nil {
				signornot = 0
			}
			log.Println(articles)
			c.JSON(200, gin.H{
				"msg":       articles,
				"signornot": signornot,
			})
		} else if status == "hot" {
			rows, err := db.Query("select * from article where postorboil= 0  order by `like` desc ,view desc ")
			if err != nil {
				log.Println(err)
				return
			}

			articles := make([]model.Article, 0)
			//  迭代查询获取数据  必须调用
			for rows.Next() {
				// row.scan 必须按照先后顺序 &获取数据
				var article model.Article
				err := rows.Scan(&article.Article_title, &article.Article_content, &article.Date, &article.Category, &article.Label, &article.Column, &article.Like, &article.Id, &article.Author, &article.Author_id, &article.View, &article.Postorboil, &article.Cover, &article.Comment, &article.Collection)
				if err != nil {
					log.Println(err)
					return
				}

				articles = append(articles, article)
			}
			var signornot = 1
			row := db.QueryRow("select date from sign where date(?)=DATE_SUB(CURDATE(), INTERVAL 0 DAY)", date)
			err = row.Scan(&date)
			if err != nil {
				signornot = 0
			}
			c.JSON(200, gin.H{
				"msg":       articles,
				"signornot": signornot,
			})
		}

	})
	Router.GET("/following", func(c *gin.Context) {
		c.HTML(200, "favorite2.html", gin.H{})
	})
	Router.POST("/following", func(c *gin.Context) {
		status := c.DefaultQuery("status", "")
		user_id, err := c.Cookie("user_id")
		if err != nil {
			log.Println(err)
			return
		}
		db, err := dao.OpenDb()
		if err != nil {
			log.Println(err)
			return
		}
		rows, err := db.Query("select followed from attention where follow=?", user_id)
		if err != nil {
			log.Println(err)
			return
		}
		articles := make([]model.Article, 0)

		if status == "" {
			for rows.Next() {
				var author_id int
				err = rows.Scan(&author_id)
				if err != nil {
					log.Println()
				}
				rows2, err := db.Query("select * from article where author_id =? and postorboil= 0 ", author_id)
				if err != nil {
					log.Println(err)
					return
				}
				for rows2.Next() {
					var article model.Article
					err = rows2.Scan(&article.Article_title, &article.Article_content, &article.Date, &article.Category, &article.Label, &article.Column, &article.Like, &article.Id, &article.Author, &article.Author_id, &article.View, &article.Postorboil, &article.Cover, &article.Comment, article.Collection)
					if err != nil {
						log.Println(err)
						return
					}
					articles = append(articles, article)
				}
			}
		} else if status == "new" {

			rows2, err := db.Query("select * from article where author_id in (select followed from attention where follow=?) and  postorboil= 0 order by date desc ", user_id)

			for rows2.Next() {
				var article model.Article
				err = rows2.Scan(&article.Article_title, &article.Article_content, &article.Date, &article.Category, &article.Label, &article.Column, &article.Like, &article.Id, &article.Author, &article.Author_id, &article.View, &article.Postorboil, &article.Cover, &article.Comment, article.Collection)
				if err != nil {
					log.Println(err)
					return
				}
				articles = append(articles, article)
			}
		} else if status == "hot" {
			rows2, err := db.Query("select * from article where author_id in (select followed from attention where follow=?)  and postorboil= 0 order by date desc ", user_id)
			for rows2.Next() {
				var article model.Article
				err = rows2.Scan(&article.Article_title, &article.Article_content, &article.Date, &article.Category, &article.Label, &article.Column, &article.Like, &article.Id, &article.Author, &article.Author_id, &article.View, &article.Postorboil, &article.Cover, &article.Comment, article.Collection)
				if err != nil {
					log.Println(err)
					return
				}
				articles = append(articles, article)
			}
		}

		c.JSON(200, gin.H{
			"msg": articles,
		})
	})

	//收藏
	Router.POST("/collection", func(c *gin.Context) {
		data, _ := c.GetRawData()
		var m map[string]any
		_ = json.Unmarshal(data, &m)
		article_id := m["article_id"]
		date := m["date"]
		user_id, err := c.Cookie("user_id")
		if err != nil {
			log.Println(err)
			return
		}

		db, _ := dao.OpenDb()
		var id int
		row := db.QueryRow("select id from collections where article_id=? and  collector_id=?", article_id, user_id)
		err = row.Scan(&id)
		if err != nil {
			result, err := db.Exec("insert into collections (article_id, collector_id, date) VALUE (?,?,?)", article_id, user_id, date)
			log.Println(result)

			if err != nil {
				log.Println(err)
				return
			}
			row := db.QueryRow("select collection from  article where  id =? and postorboil= 0 ", article_id)
			var collection int
			row.Scan(&collection)
			db.Exec("update article set collection=? where id=?", collection+1, article_id)
			c.JSON(200, gin.H{
				"msg": "收藏成功",
			})
			return
		}
		result, err := db.Exec("delete from collections where collector_id=? and  article_id=?", user_id, article_id)
		log.Println(result)
		if err != nil {
			log.Println(err)
			return
		}
		c.JSON(200, gin.H{
			"msg": "取消收藏",
		})
	})
	//首页后端
	Router.GET("/backend", func(c *gin.Context) {
		c.HTML(200, "backend2.html", gin.H{})
	})
	Router.POST("/backend", func(c *gin.Context) {
		status := c.DefaultQuery("status", "")
		db, err := dao.OpenDb()
		if err != nil {
			log.Println(err)
			return
		}
		articles := make([]model.Article, 0)
		if status == "" {

			rows, err := db.Query("select * from article where category =? and postorboil= 0 ", "后端")
			if err != nil {
				log.Println(err)
				return
			}

			for rows.Next() {

				var article model.Article
				err := rows.Scan(&article.Article_title, &article.Article_content, &article.Date, &article.Category, &article.Label, &article.Column, &article.Like, &article.Id, &article.Author, &article.Author_id, &article.View, &article.Postorboil, &article.Cover, &article.Comment, &article.Collection)
				if err != nil {
					log.Println(err)
					return
				}

				articles = append(articles, article)
			}
		} else if status == "new" {
			rows, err := db.Query("select * from article where category =? and postorboil= 0 order by date desc ", "后端")
			if err != nil {
				log.Println(err)
				return
			}

			for rows.Next() {

				var article model.Article
				err := rows.Scan(&article.Article_title, &article.Article_content, &article.Date, &article.Category, &article.Label, &article.Column, &article.Like, &article.Id, &article.Author, &article.Author_id, &article.View, &article.Postorboil, &article.Cover, &article.Comment, &article.Collection)
				if err != nil {
					log.Println(err)
					return
				}

				articles = append(articles, article)
			}
		} else if status == "hot" {
			rows, err := db.Query("select * from article where category =? and postorboil= 0 order by `like` desc,view desc ", "后端")
			if err != nil {
				log.Println(err)
				return
			}

			for rows.Next() {

				var article model.Article
				err := rows.Scan(&article.Article_title, &article.Article_content, &article.Date, &article.Category, &article.Label, &article.Column, &article.Like, &article.Id, &article.Author, &article.Author_id, &article.View, &article.Postorboil, &article.Cover, &article.Comment, &article.Collection)
				if err != nil {
					log.Println(err)
					return
				}

				articles = append(articles, article)
			}
		}

		c.JSON(200, gin.H{
			"msg": articles,
		})

	})

	//首页前端
	Router.GET("/frontend", func(c *gin.Context) {
		c.HTML(200, "headend2.html", gin.H{})
	})
	Router.POST("/frontend", func(c *gin.Context) {

		status := c.DefaultQuery("status", "")
		db, err := dao.OpenDb()
		if err != nil {
			log.Println(err)
			return
		}
		articles := make([]model.Article, 0)
		if status == "" {

			rows, err := db.Query("select * from article where category =? and postorboil= 0 ", "前端")
			if err != nil {
				log.Println(err)
				return
			}

			for rows.Next() {

				var article model.Article
				err := rows.Scan(&article.Article_title, &article.Article_content, &article.Date, &article.Category, &article.Label, &article.Column, &article.Like, &article.Id, &article.Author, &article.Author_id, &article.View, &article.Postorboil, &article.Cover, &article.Comment, &article.Collection)
				if err != nil {
					log.Println(err)
					return
				}

				articles = append(articles, article)
			}
		} else if status == "new" {
			rows, err := db.Query("select * from article where category =? and postorboil= 0 order by date desc ", "前端")
			if err != nil {
				log.Println(err)
				return
			}

			for rows.Next() {

				var article model.Article
				err := rows.Scan(&article.Article_title, &article.Article_content, &article.Date, &article.Category, &article.Label, &article.Column, &article.Like, &article.Id, &article.Author, &article.Author_id, &article.View, &article.Postorboil, &article.Cover, &article.Comment, &article.Collection)
				if err != nil {
					log.Println(err)
					return
				}

				articles = append(articles, article)
			}
		} else if status == "hot" {
			rows, err := db.Query("select * from article where category =? and postorboil= 0 order by `like` desc,view desc ", "前端")
			if err != nil {
				log.Println(err)
				return
			}

			for rows.Next() {

				var article model.Article
				err := rows.Scan(&article.Article_title, &article.Article_content, &article.Date, &article.Category, &article.Label, &article.Column, &article.Like, &article.Id, &article.Author, &article.Author_id, &article.View, &article.Postorboil, &article.Cover, &article.Comment, &article.Collection)
				if err != nil {
					log.Println(err)
					return
				}

				articles = append(articles, article)
			}
		}

		c.JSON(200, gin.H{
			"msg": articles,
		})

	})
	//首页安卓
	Router.GET("/android", func(c *gin.Context) {
		c.HTML(200, "Android2.html", gin.H{})
	})
	Router.POST("/android", func(c *gin.Context) {

		status := c.DefaultQuery("status", "")
		db, err := dao.OpenDb()
		if err != nil {
			log.Println(err)
			return
		}
		articles := make([]model.Article, 0)
		if status == "" {

			rows, err := db.Query("select * from article where category =? and postorboil= ? ", "Android", 0)
			if err != nil {
				log.Println(err)
				return
			}

			for rows.Next() {

				var article model.Article
				err := rows.Scan(&article.Article_title, &article.Article_content, &article.Date, &article.Category, &article.Label, &article.Column, &article.Like, &article.Id, &article.Author, &article.Author_id, &article.View, &article.Postorboil, &article.Cover, &article.Comment, &article.Collection)
				if err != nil {
					log.Println(err)
					return
				}

				articles = append(articles, article)
			}
		} else if status == "new" {
			rows, err := db.Query("select * from article where category =?and postorboil= ?  order by date desc ", "Android", 0)
			if err != nil {
				log.Println(err)
				return
			}

			for rows.Next() {

				var article model.Article
				err := rows.Scan(&article.Article_title, &article.Article_content, &article.Date, &article.Category, &article.Label, &article.Column, &article.Like, &article.Id, &article.Author, &article.Author_id, &article.View, &article.Postorboil, &article.Cover, &article.Comment, &article.Collection)
				if err != nil {
					log.Println(err)
					return
				}

				articles = append(articles, article)
			}
		} else if status == "hot" {
			rows, err := db.Query("select * from article where category =?and postorboil= ?  order by `like` desc,view desc ", "Android", 0)
			if err != nil {
				log.Println(err)
				return
			}

			for rows.Next() {

				var article model.Article
				err := rows.Scan(&article.Article_title, &article.Article_content, &article.Date, &article.Category, &article.Label, &article.Column, &article.Like, &article.Id, &article.Author, &article.Author_id, &article.View, &article.Postorboil, &article.Cover, &article.Comment, &article.Collection)
				if err != nil {
					log.Println(err)
					return
				}

				articles = append(articles, article)
			}
		}

		c.JSON(200, gin.H{
			"msg": articles,
		})

	})
	//首页ios
	Router.GET("/ios", func(c *gin.Context) {
		c.HTML(200, "iOS2.html", gin.H{})
	})
	Router.POST("/ios", func(c *gin.Context) {

		status := c.DefaultQuery("status", "")
		db, err := dao.OpenDb()
		if err != nil {
			log.Println(err)
			return
		}
		articles := make([]model.Article, 0)
		if status == "" {

			rows, err := db.Query("select * from article where category =? and postorboil= ? ", "iOS", 0)
			if err != nil {
				log.Println(err)
				return
			}

			for rows.Next() {

				var article model.Article
				err := rows.Scan(&article.Article_title, &article.Article_content, &article.Date, &article.Category, &article.Label, &article.Column, &article.Like, &article.Id, &article.Author, &article.Author_id, &article.View, &article.Postorboil, &article.Cover, &article.Comment, &article.Collection)
				if err != nil {
					log.Println(err)
					return
				}

				articles = append(articles, article)
			}
		} else if status == "new" {
			rows, err := db.Query("select * from article where category =?  and postorboil= ?  order by date desc ", "iOS", 0)
			if err != nil {
				log.Println(err)
				return
			}

			for rows.Next() {

				var article model.Article
				err := rows.Scan(&article.Article_title, &article.Article_content, &article.Date, &article.Category, &article.Label, &article.Column, &article.Like, &article.Id, &article.Author, &article.Author_id, &article.View, &article.Postorboil, &article.Cover, &article.Comment, &article.Collection)
				if err != nil {
					log.Println(err)
					return
				}

				articles = append(articles, article)
			}
		} else if status == "hot" {
			rows, err := db.Query("select * from article where category =? and postorboil= ? order by `like` desc,view desc ", "iOS", 0)
			if err != nil {
				log.Println(err)
				return
			}

			for rows.Next() {

				var article model.Article
				err := rows.Scan(&article.Article_title, &article.Article_content, &article.Date, &article.Category, &article.Label, &article.Column, &article.Like, &article.Id, &article.Author, &article.Author_id, &article.View, &article.Postorboil, &article.Cover, &article.Comment, &article.Collection)
				if err != nil {
					log.Println(err)
					return
				}

				articles = append(articles, article)
			}
		}
		c.JSON(200, gin.H{
			"msg": articles,
		})
	})

	//首页ai
	Router.GET("/ai", func(c *gin.Context) {
		c.HTML(200, "AI2.html", gin.H{})
	})
	Router.POST("/ai", func(c *gin.Context) {

		status := c.DefaultQuery("status", "")
		db, err := dao.OpenDb()
		if err != nil {
			log.Println(err)
			return
		}
		articles := make([]model.Article, 0)
		if status == "" {

			rows, err := db.Query("select * from article where category =? and postorboil= ? ", "人工智能", 0)
			if err != nil {
				log.Println(err)
				return
			}

			for rows.Next() {

				var article model.Article
				err := rows.Scan(&article.Article_title, &article.Article_content, &article.Date, &article.Category, &article.Label, &article.Column, &article.Like, &article.Id, &article.Author, &article.Author_id, &article.View, &article.Postorboil, &article.Cover, &article.Comment, &article.Collection)
				if err != nil {
					log.Println(err)
					return
				}

				articles = append(articles, article)
			}
		} else if status == "new" {
			rows, err := db.Query("select * from article where category =? and postorboil= ?  order by date desc ", "人工智能", 0)
			if err != nil {
				log.Println(err)
				return
			}

			for rows.Next() {

				var article model.Article
				err := rows.Scan(&article.Article_title, &article.Article_content, &article.Date, &article.Category, &article.Label, &article.Column, &article.Like, &article.Id, &article.Author, &article.Author_id, &article.View, &article.Postorboil, &article.Cover, &article.Comment, &article.Collection)
				if err != nil {
					log.Println(err)
					return
				}

				articles = append(articles, article)
			}
		} else if status == "hot" {
			rows, err := db.Query("select * from article where category =? and postorboil= ? order by `like` desc,view desc ", "人工智能", 0)
			if err != nil {
				log.Println(err)
				return
			}

			for rows.Next() {

				var article model.Article
				err := rows.Scan(&article.Article_title, &article.Article_content, &article.Date, &article.Category, &article.Label, &article.Column, &article.Like, &article.Id, &article.Author, &article.Author_id, &article.View, &article.Postorboil, &article.Cover, &article.Comment, &article.Collection)
				if err != nil {
					log.Println(err)
					return
				}

				articles = append(articles, article)
			}
		}
		c.JSON(200, gin.H{
			"msg": articles,
		})

	})
	//首页开发工具
	Router.GET("/freebie", func(c *gin.Context) {
		c.HTML(200, "tool2.html", gin.H{})
	})
	Router.POST("/freebie", func(c *gin.Context) {

		status := c.DefaultQuery("status", "")
		db, err := dao.OpenDb()
		if err != nil {
			log.Println(err)
			return
		}
		articles := make([]model.Article, 0)
		if status == "" {

			rows, err := db.Query("select * from article where category =? and postorboil= ? ", "开发工具", 0)
			if err != nil {
				log.Println(err)
				return
			}

			for rows.Next() {

				var article model.Article
				err := rows.Scan(&article.Article_title, &article.Article_content, &article.Date, &article.Category, &article.Label, &article.Column, &article.Like, &article.Id, &article.Author, &article.Author_id, &article.View, &article.Postorboil, &article.Cover, &article.Comment, &article.Collection)
				if err != nil {
					log.Println(err)
					return
				}

				articles = append(articles, article)
			}
		} else if status == "new" {
			rows, err := db.Query("select * from article where category =? and postorboil= ?  order by date desc ", "开发工具", 0)
			if err != nil {
				log.Println(err)
				return
			}

			for rows.Next() {

				var article model.Article
				err := rows.Scan(&article.Article_title, &article.Article_content, &article.Date, &article.Category, &article.Label, &article.Column, &article.Like, &article.Id, &article.Author, &article.Author_id, &article.View, &article.Postorboil, &article.Cover, &article.Comment, &article.Collection)
				if err != nil {
					log.Println(err)
					return
				}

				articles = append(articles, article)
			}
		} else if status == "hot" {
			rows, err := db.Query("select * from article where category =? and postorboil= ?  order by `like` desc,view desc ", "开发工具", 0)
			if err != nil {
				log.Println(err)
				return
			}

			for rows.Next() {

				var article model.Article
				err := rows.Scan(&article.Article_title, &article.Article_content, &article.Date, &article.Category, &article.Label, &article.Column, &article.Like, &article.Id, &article.Author, &article.Author_id, &article.View, &article.Postorboil, &article.Cover, &article.Comment, &article.Collection)
				if err != nil {
					log.Println(err)
					return
				}

				articles = append(articles, article)
			}
		}
		c.JSON(200, gin.H{
			"msg": articles,
		})
	})
	//代码人生
	Router.GET("/career", func(c *gin.Context) {
		c.HTML(200, "code2.html", gin.H{})
	})
	Router.POST("/career", func(c *gin.Context) {

		status := c.DefaultQuery("status", "")
		db, err := dao.OpenDb()
		if err != nil {
			log.Println(err)
			return
		}
		articles := make([]model.Article, 0)
		if status == "" {

			rows, err := db.Query("select * from article where category =? and postorboil= ? ", "代码人生", 0)
			if err != nil {
				log.Println(err)
				return
			}

			for rows.Next() {

				var article model.Article
				err := rows.Scan(&article.Article_title, &article.Article_content, &article.Date, &article.Category, &article.Label, &article.Column, &article.Like, &article.Id, &article.Author, &article.Author_id, &article.View, &article.Postorboil, &article.Cover, &article.Comment, &article.Collection)
				if err != nil {
					log.Println(err)
					return
				}

				articles = append(articles, article)
			}
		} else if status == "new" {
			rows, err := db.Query("select * from article where category = ? and postorboil=?   order by date desc ", "代码人生", 0)
			if err != nil {
				log.Println(err)
				return
			}

			for rows.Next() {

				var article model.Article
				err := rows.Scan(&article.Article_title, &article.Article_content, &article.Date, &article.Category, &article.Label, &article.Column, &article.Like, &article.Id, &article.Author, &article.Author_id, &article.View, &article.Postorboil, &article.Cover, &article.Comment, &article.Collection)
				if err != nil {
					log.Println(err)
					return
				}

				articles = append(articles, article)
			}
		} else if status == "hot" {
			rows, err := db.Query("select * from article where category =? and postorboil= ? order by `like` desc, view desc ", "代码人生", 0)
			if err != nil {
				log.Println(err)
				return
			}

			for rows.Next() {

				var article model.Article
				err := rows.Scan(&article.Article_title, &article.Article_content, &article.Date, &article.Category, &article.Label, &article.Column, &article.Like, &article.Id, &article.Author, &article.Author_id, &article.View, &article.Postorboil, &article.Cover, &article.Comment, &article.Collection)
				if err != nil {
					log.Println(err)
					return
				}

				articles = append(articles, article)
			}
		}
		c.JSON(200, gin.H{
			"msg": articles,
		})
	})
	Router.GET("/read", func(c *gin.Context) {
		c.HTML(200, "read2.html", gin.H{})
	})
	Router.POST("/read", func(c *gin.Context) {
		status := c.DefaultQuery("status", "")
		db, err := dao.OpenDb()
		if err != nil {
			log.Println(err)
			return
		}
		articles := make([]model.Article, 0)
		if status == "" {

			rows, err := db.Query("select * from article where category =? and postorboil= ? ", "阅读", 0)
			if err != nil {
				log.Println(err)
				return
			}

			for rows.Next() {

				var article model.Article
				err := rows.Scan(&article.Article_title, &article.Article_content, &article.Date, &article.Category, &article.Label, &article.Column, &article.Like, &article.Id, &article.Author, &article.Author_id, &article.View, &article.Postorboil, &article.Cover, &article.Comment, &article.Collection)
				if err != nil {
					log.Println(err)
					return
				}

				articles = append(articles, article)
			}
		} else if status == "new" {
			rows, err := db.Query("select * from article where category =? and postorboil= ?  order by date desc ", "阅读", 0)
			if err != nil {
				log.Println(err)
				return
			}

			for rows.Next() {

				var article model.Article
				err := rows.Scan(&article.Article_title, &article.Article_content, &article.Date, &article.Category, &article.Label, &article.Column, &article.Like, &article.Id, &article.Author, &article.Author_id, &article.View, &article.Postorboil, &article.Cover, &article.Comment, &article.Collection)
				if err != nil {
					log.Println(err)
					return
				}

				articles = append(articles, article)
			}
		} else if status == "hot" {
			rows, err := db.Query("select * from article where category =? and postorboil= ?  order by `like` desc,view desc ", "阅读", 0)
			if err != nil {
				log.Println(err)
				return
			}

			for rows.Next() {

				var article model.Article
				err := rows.Scan(&article.Article_title, &article.Article_content, &article.Date, &article.Category, &article.Label, &article.Column, &article.Like, &article.Id, &article.Author, &article.Author_id, &article.View, &article.Postorboil, &article.Cover, &article.Comment, &article.Collection)
				if err != nil {
					log.Println(err)
					return
				}

				articles = append(articles, article)
			}
		}
		c.JSON(200, gin.H{
			"msg": articles,
		})
	})

	Router.GET("/signin", func(c *gin.Context) {
		c.HTML(200, "qiandao.html", gin.H{})
	})
	//签到
	Router.POST("/signin", func(c *gin.Context) {
		data, _ := c.GetRawData()
		var m map[string]any
		_ = json.Unmarshal(data, &m)
		date := m["date"]
		db, err := dao.OpenDb()
		if err != nil {
			log.Println(err)
			return
		}
		user_id, err := c.Cookie("user_id")
		if err != nil {
			log.Println(err)
			return
		}

		row := db.QueryRow("select date from sign where date(?)=DATE_SUB(CURDATE(), INTERVAL 0 DAY)", date)
		err = row.Scan(&date)
		if err != nil {
			db.Exec("insert into sign (user_id, date) VALUE (?,?)", user_id, date)
		}

		rows, err := db.Query("select date from sign where user_id=?", user_id)
		var days = 0
		for rows.Next() {
			days++
		}

		c.JSON(200, gin.H{
			"days": days,
		})

	})
	//搜索
	Router.GET("/search", func(c *gin.Context) {
		c.HTML(200, "search1.html", gin.H{})
	})
	Router.POST("/search", func(c *gin.Context) {
		data, _ := c.GetRawData()
		var m map[string]any
		_ = json.Unmarshal(data, &m)
		search := m["search"]
		log.Println(search)
		db, err := dao.OpenDb()
		if err != nil {
			log.Println(err)
			return
		}
		articles := make([]model.Article, 0)
		rows, err := db.Query("select * from article")
		for rows.Next() {
			var article model.Article
			err := rows.Scan(&article.Article_title, &article.Article_content, &article.Date, &article.Category, &article.Label, &article.Column, &article.Like, &article.Id, &article.Author, &article.Author_id, &article.View, &article.Postorboil, &article.Cover, &article.Comment, &article.Collection)
			if err != nil {
				log.Println(err)
				return
			}
			if strings.Contains(article.Article_title, fmt.Sprint(search)) {
				articles = append(articles, article)
			}
		}
		log.Println(articles)
		c.JSON(200, gin.H{
			"msg": articles,
		})
	})
	//课程

	Router.GET("/course", func(c *gin.Context) {
		c.HTML(200, "class-home.html", gin.H{})
	})

	Router.GET("/course/ai", func(c *gin.Context) {
		c.HTML(200, "class-home-AI.html", gin.H{})
	})
	Router.GET("/course/android", func(c *gin.Context) {
		c.HTML(200, "class-home-Android.html", gin.H{})
	})
	Router.GET("/course/read", func(c *gin.Context) {
		c.HTML(200, "class-home-read.html", gin.H{})
	})
	Router.GET("/course/backend", func(c *gin.Context) {
		c.HTML(200, "class-home-backend.html", gin.H{})
	})
	Router.GET("/course/code", func(c *gin.Context) {
		c.HTML(200, "class-home-code.html", gin.H{})
	})
	Router.GET("/course/frontend", func(c *gin.Context) {
		c.HTML(200, "class-home-frontend.html", gin.H{})
	})
	Router.GET("/course/ios", func(c *gin.Context) {
		c.HTML(200, "class-home-iOS.html", gin.H{})
	})
	Router.GET("/course/tool", func(c *gin.Context) {
		c.HTML(200, "class-home-tool.html", gin.H{})
	})
	Router.POST("/course", func(c *gin.Context) {
		data, _ := c.GetRawData()
		var m map[string]any
		_ = json.Unmarshal(data, &m)
		catalogue := m["catalogue"]
		status := m["status"]
		db, err := dao.OpenDb()

		if err != nil {
			log.Println(err)
			return
		}
		courses := make([]model.Course, 0)
		if catalogue == "" {
			if status == "" {

				rows, err := db.Query("select * from course ")

				for rows.Next() {
					var course model.Course
					err = rows.Scan(&course.Course_title, &course.Introduction, &course.Course_id, &course.Abstract, &course.Cover, &course.Catalogue, &course.Hot)
					if err != nil {
						log.Println(err)
						return
					}
					courses = append(courses, course)
				}
				log.Println("1")
			} else {
				rows, err := db.Query("select * from course order by hot")

				for rows.Next() {
					var course model.Course
					err = rows.Scan(&course.Course_title, &course.Introduction, &course.Course_id, &course.Abstract, &course.Cover, &course.Catalogue, &course.Hot)
					if err != nil {
						log.Println(err)
						return
					}
					courses = append(courses, course)
				}
				log.Println("2")
			}
		} else {

			if status == "hot" {
				rows, err := db.Query("select * from course where catalogue =? order by hot desc ", catalogue)

				for rows.Next() {
					var course model.Course
					err = rows.Scan(&course.Course_title, &course.Introduction, &course.Course_id, &course.Abstract, &course.Cover, &course.Catalogue, &course.Hot)
					if err != nil {
						log.Println(err)
						return
					}
					courses = append(courses, course)
				}
				log.Println("3")
			} else if status == "" {
				rows, err := db.Query("select * from course where catalogue=?", catalogue)

				for rows.Next() {
					var course model.Course
					err = rows.Scan(&course.Course_title, &course.Introduction, &course.Course_id, &course.Abstract, &course.Cover, &course.Catalogue, &course.Hot)
					if err != nil {
						log.Println(err)
						return
					}
					courses = append(courses, course)
				}
				log.Println("4")
			}
		}

		c.JSON(200, gin.H{
			"msg": courses,
		})
	})

	//课程详细也
	Router.GET("/book", func(c *gin.Context) {
		c.HTML(200, "class-detail.html", gin.H{})
	})
	Router.POST("/book", func(c *gin.Context) {
		data, _ := c.GetRawData()
		var m map[string]any
		_ = json.Unmarshal(data, &m)
		course_id := m["course_id"]
		db, err := dao.OpenDb()
		if err != nil {
			log.Println(err)
			return
		}
		row := db.QueryRow("select * from course where course_id=?", course_id)
		var course model.Course
		err = row.Scan(&course.Course_title, &course.Introduction, &course.Course_id, &course.Abstract, &course.Cover, &course.Catalogue, &course.Hot)
		if err != nil {
			log.Println(err)
			return
		}
		var purchaseornot = 1
		user_id, err := c.Cookie("user_id")
		if err != nil {
			log.Println(err)
			return
		}
		row = db.QueryRow("select user_id from course_purchase where user_id=? and course_id=?", user_id, course_id)
		err = row.Scan(&user_id)
		if err != nil {
			purchaseornot = 0
		}

		c.JSON(200, gin.H{
			"msg":           course,
			"purchaseornot": purchaseornot,
		})
	})
	Router.POST("/purchase", func(c *gin.Context) {
		data, _ := c.GetRawData()
		var m map[string]any
		_ = json.Unmarshal(data, &m)
		course_id := m["course_id"]
		user_id, err := c.Cookie("user_id")

		db, err := dao.OpenDb()
		if err != nil {
			log.Println(err)
			return
		}

		result, err := db.Exec("insert into course_purchase (user_id, course_id) VALUE (?,?)", user_id, course_id)
		log.Println(result)
		if err != nil {
			log.Println(err)
			return
		}
		result, err = db.Exec("update course set hot = hot+1 where course_id=?", course_id)
		log.Println(result)
		if err != nil {
			log.Println(err)
			return
		}

		c.JSON(200, gin.H{
			"msg": "购买成功",
		})
	})
	Router.GET("/catelogue", func(c *gin.Context) {
		c.HTML(200, "class-catelogue.html", gin.H{})
	})
	Router.POST("/catelogue", func(c *gin.Context) {
		data, _ := c.GetRawData()
		var m map[string]any
		_ = json.Unmarshal(data, &m)
		course_id := m["course_id"]

		db, err := dao.OpenDb()
		if err != nil {
			log.Println(err)
			return
		}
		rows, err := db.Query("select * from catalogue where course_id=? order by `order`", course_id)
		if err != nil {
			log.Println(err)
			return
		}
		catalogues := make([]model.Catalogue, 0)
		for rows.Next() {
			var catalogue model.Catalogue
			err := rows.Scan(&catalogue.Catalogue_id, &catalogue.Course_id, &catalogue.Catalogue_title, &catalogue.Trialornot, &catalogue.Order, &catalogue.Content)
			if err != nil {
				log.Println(err)
				return
			}
			catalogues = append(catalogues, catalogue)

		}
		var purchaseornot = 1
		user_id, err := c.Cookie("user_id")
		if err != nil {
			log.Println(err)
			return
		}
		row := db.QueryRow("select user_id from course_purchase where user_id=? and course_id=?", user_id, course_id)
		err = row.Scan(&user_id)
		if err != nil {
			purchaseornot = 0
		}
		row = db.QueryRow("select * from course where course_id=?", course_id)
		var course model.Course
		err = row.Scan(&course.Course_title, &course.Introduction, &course.Course_id, &course.Abstract, &course.Cover, &course.Catalogue, &course.Hot)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(purchaseornot)

		c.JSON(200, gin.H{
			"course":        course,
			"catalogues":    catalogues,
			"purchaseornot": purchaseornot,
			"msg":           catalogues,
		})
	})

	Router.GET("/book/section", func(c *gin.Context) {
		c.HTML(200, "class-content.html", gin.H{})
	})
	//章节详细页
	Router.POST("/book/section", func(c *gin.Context) {
		data, _ := c.GetRawData()
		var m map[string]any
		_ = json.Unmarshal(data, &m)
		catalogue_id := m["catalogue_id"]
		db, err := dao.OpenDb()
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(catalogue_id)
		row := db.QueryRow("select * from catalogue where catalogue_id=?", catalogue_id)
		var catalogue model.Catalogue
		err = row.Scan(&catalogue.Catalogue_id, &catalogue.Course_id, &catalogue.Catalogue_title, &catalogue.Trialornot, &catalogue.Order, &catalogue.Content)
		if err != nil {
			log.Println(err)
			return
		}
		user_id, err := c.Cookie("user_id")
		if err != nil {
			log.Println(err)
			return
		}
		var purchaseornot = 1
		row = db.QueryRow("select user_id from course_purchase where course_id in (select course_id from catalogue where catalogue_id=?) and user_id=?", catalogue_id, user_id)
		err = row.Scan(&user_id)
		if err != nil {
			purchaseornot = 0
		}
		log.Println(purchaseornot)
		c.JSON(200, gin.H{
			"msg":           catalogue,
			"purchaseornot": purchaseornot,
		})
	})
	//关注
	Router.POST("/notice", func(c *gin.Context) {
		data, _ := c.GetRawData()
		var m map[string]any
		_ = json.Unmarshal(data, &m)
		followed_id := m["followed_id"]
		date := m["date"]
		db, err := dao.OpenDb()

		user_id, err := c.Cookie("user_id")
		if err != nil {
			log.Println(err)
			return
		}
		if err != nil {
			log.Println(err)
			return
		}

		row := db.QueryRow("select status from attention where follow=? and followed=?", user_id, followed_id)
		var status string
		err = row.Scan(&status)

		if err != nil {
			db.Exec("insert into attention (followed, follow, date, status) VALUE (?, ?, ?, ?)", followed_id, user_id, date, "关注")
			c.JSON(200, gin.H{
				"msg": "关注成功",
			})
			return
		}
		if status == "取消关注" {
			db.Exec("update attention set status =? where followed=? and follow=?", "关注", followed_id, user_id)

			c.JSON(200, gin.H{
				"msg": "关注成功",
			})
			return
		}
		db.Exec("update attention set status =? where followed=? and follow=?", "取消关注", followed_id, user_id)

		c.JSON(200, gin.H{
			"msg": "取消关注",
		})
	})
	//创建沸点
	Router.POST("/pins", func(c *gin.Context) {
		data, _ := c.GetRawData()
		var m map[string]any
		_ = json.Unmarshal(data, &m)
		content := m["content"]
		date := m["date"]
		phoneoremail, err := c.Cookie("phoneoremail")
		user_id, err := c.Cookie("user_id")
		db, err := dao.OpenDb()
		if err != nil {
			log.Println(err)
			return
		}
		result, err := db.Exec("insert into article (article_title,label,`column`,article_content,author_id,author,view,postorboil,comments,`like`,date,category,cover,collection) value (?,?,?,?,?,?,?,?,?,?,?,?,?,?)", "", "", "", content, user_id, phoneoremail, 0, 1, 0, 0, date, "", "", 0)
		log.Println(result)
		if err != nil {
			log.Println(err)
			return
		}
		row := db.QueryRow("select `like`,view,comments from article where article_content=? and author_id=?", content, user_id)

		var like, view, comment int
		err = row.Scan(&like, &view, &comment)
		if err != nil {
			log.Println(err)
			return
		}
		c.JSON(200, gin.H{
			"msg":     "创建沸点成功",
			"content": content,
			"date":    date,
			"like":    like,
			"view":    view,
			"comment": comment,
		})

	})

	//专栏详情页
	Router.GET("/column", func(c *gin.Context) {
		c.HTML(200, "column.html", gin.H{})
	})
	Router.POST("/column", func(c *gin.Context) {
		data, _ := c.GetRawData()
		var m map[string]any
		_ = json.Unmarshal(data, &m)
		column_id := m["column_id"]
		db, err := dao.OpenDb()
		if err != nil {
			log.Println(err)
			return
		}
		row := db.QueryRow("select * from `column` where id=?", column_id)
		var column model.Column
		err = row.Scan(&column.Column_title, &column.Column_intruduction, &column.Phoneoremail, &column.Id, &column.Cover, &column.Column_number, &column.Time)
		if err != nil {
			log.Println(err)
			return
		}
		rows, err := db.Query("select * from article where `column` in (select column_title from `column` where `column`.id=?) and postorboil=0", column_id)
		if err != nil {
			log.Println(err)
			return
		}

		articles := make([]model.Article, 0)

		for rows.Next() {

			var article model.Article
			err := rows.Scan(&article.Article_title, &article.Article_content, &article.Date, &article.Category, &article.Label, &article.Column, &article.Like, &article.Id, &article.Author, &article.Author_id, &article.View, &article.Postorboil, &article.Cover, &article.Comment, &article.Collection)
			if err != nil {
				log.Println(err)
				return
			}

			articles = append(articles, article)
		}

		c.JSON(200, gin.H{
			"column":   column,
			"articles": articles,
		})

	})

	Router.POST("/deletearticle", func(c *gin.Context) {
		data, _ := c.GetRawData()
		var m map[string]any
		_ = json.Unmarshal(data, &m)
		article_id := m["article_id"]
		db, err := dao.OpenDb()
		if err != nil {
			log.Println(err)
			return
		}
		result, err := db.Exec("delete from article where id=?", article_id)
		log.Println(result)
		if err != nil {
			log.Println(err)
			return
		}
		result, err = db.Exec("delete from collections where article_id=?", article_id)
		log.Println(result)
		if err != nil {
			log.Println(err)
			return
		}
		result, err = db.Exec("delete from comment where projectid=? ", article_id)
		log.Println(result)
		if err != nil {
			log.Println(err)
			return
		}
		result, err = db.Exec("delete from `like` where liked_id=? ", article_id)
		log.Println(result)
		if err != nil {
			log.Println(err)
			return
		}
		result, err = db.Exec("delete from view where article_id=? ", article_id)
		log.Println(result)
		if err != nil {
			log.Println(err)
			return
		}
		c.JSON(200, gin.H{
			"msg": "删除成功",
		})
	})
	Router.POST("/deletecolumn", func(c *gin.Context) {
		data, _ := c.GetRawData()
		var m map[string]any
		_ = json.Unmarshal(data, &m)
		column_id := m["column_id"]
		db, err := dao.OpenDb()
		if err != nil {
			log.Println(err)
			return
		}
		result, err := db.Exec("delete from `column` where id=?", column_id)
		log.Println(result)
		if err != nil {
			log.Println(err)
			return
		}
		result, err = db.Exec("update article set `column`='' where `column` in (select column_title from `column` where `column`.id=?)", column_id)
		log.Println(result)
		if err != nil {
			log.Println(err)
			return
		}
		c.JSON(200, gin.H{
			"msg": "删除专栏",
		})
	})
	Router.Run(":8080")
}
