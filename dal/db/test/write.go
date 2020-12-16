package main

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/BlogByFourMan/Server/dal/db"

	"github.com/BlogByFourMan/Server/dal/model"
)

func main() {
	articles := make([]model.Article, 7)
	users := make([]model.User, 7)
	titles := []string{"2021年GO语言（Golang就业班全系列", "Beego_ubuntu下golang及beego环境的全局配置", "Go - 学习/实践", 
					   "阿里巴巴十年Java架构师分享，会了这个知识点的人都去BAT了", "领域驱动实践总结(基本理论总结与分析+架构分析与代码设计V+具体应用设计分析)",
					   "B/S架构及其运行原理", "Go&&阿里云服务器（Ubuntu）-- Golang项目（beego）服务器部署" }
	authors := []string{"小龙in武汉", "Grayan", "Hi ning先森", "JAVA高级架构v", "张彦峰ZYF", "蜘蛛侠不会飞", "李歘歘"}
	times := []string{"2020-09-17 14:04:21", "2019-03-13 22:28:32", "2019-11-04 11:43:11", 
	                  "2019-05-14 17:38:32", "2020-03-26 18:34:46", "2018-03-23 22:31:19", "2019-07-19 14:23:37"}

	for i := 0; i < 7; i++ {
		f, err := os.Open(strconv.FormatInt(int64(i), 10) + ".html")
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		content, _ := ioutil.ReadAll(f)
		
		// 得到文章和用户
		article := model.Article{int64(i + 1), titles[i], authors[i], nil, times[i], string(content), nil}
		articles = append(articles, article)
		user := model.User{authors[i], "1234567"}
		users = append(users, user)
	}
	
	// 保存进数据库
	db.PutArticles(articles)
	db.PutUsers(users)
}