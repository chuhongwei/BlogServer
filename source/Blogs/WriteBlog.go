package main

import (
	"io/ioutil"
	"log"
	"os"
	
	//"user/BlogServer/source"
	"github.com/chuhongwei/BlogServer/source"
)

func writeOneBlog(id int64,title string,author string,tags []db.Tag,date string,content string,comments []db.Comment){

	articles := db.Article{id,title,author,tags,date,content,comments}
	users := db.User{author,author}
	db.PutArticle(articles)
	db.PutUser(users)
}
func main() {
	titles := []string{"2021年GO语言（Golang就业班全系列", "Beego_ubuntu下golang及beego环境的全局配置", "Go学习实践", 
					   "阿里巴巴十年Java架构师分享", "领域驱动实践总结","BS架构及其运行原理", "Go&&阿里云服务器" }
	authors := []string{"小龙in武汉", "Grayan", "Hi ning先森", "JAVA高级架构v", "张彦峰ZYF", "蜘蛛侠不会飞", "李歘歘"}
	times := []string{"2020-09-17 14:04:21", "2019-03-13 22:28:32", "2019-11-04 11:43:11", 
	                  "2019-05-14 17:38:32", "2020-03-26 18:34:46", "2018-03-23 22:31:19", "2019-07-19 14:23:37"}
	
	db.Init()

	for i := 0; i < 7; i++ {
		f, err := os.Open(titles[i] + ".html")
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		content, _ := ioutil.ReadAll(f)
		writeOneBlog(int64(i + 1), titles[i], authors[i], nil, times[i], string(content), nil)
	}
}