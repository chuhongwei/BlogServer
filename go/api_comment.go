package swagger

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
	"encoding/json"

	"github.com/chuhongwei/BlogServer/dal/db"
	"github.com/chuhongwei/BlogServer/dal/model"
	"github.com/dgrijalva/jwt-go"
)

// GET /article/1/comments
func GetArticleCommentsByID(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
    // 分解 URL
	str := strings.Split(r.URL.Path, "/")
	articleID, err := strconv.ParseInt((str[len(str)-2]), 10, 64)
	fmt.Println(articleID)
	if err != nil {
		Response(ResponseMessage{nil, err.Error()}, w, http.StatusBadRequest)
		return
	}

	articles := db.GetArticles(articleID, 0)
	if len(articles) != 0 && len(articles[0].Comments) != 0 {
		Response(ResponseMessage{articles[0].Comments,nil,}, w, http.StatusOK)
	} else if len(articles) == 0 {
		Response(ResponseMessage{nil, "Article Not Found"}, w, http.StatusOK)
	} else {
		Response(ResponseMessage{nil, nil}, w, http.StatusOK)
	}
}


// POST /article/1/comments
// 包含 path parameter id, body parameter: 用户验证 token 和 评论内容
func PostArticleCommentsByID(w http.ResponseWriter, r *http.Request) {
	db.Init()
	// 验证用户 token
	token, isValid := CheckToken(w, r)
	if isValid == false {
		Response(ResponseMessage{nil,"Authentication Fail !"}, w, http.StatusBadRequest)
		return
	}

	var comment model.Comment
	// 读 body parameter，decode 为 Comment 结构
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		Response(ResponseMessage{nil,err.Error()}, w, http.StatusBadRequest)
		return
	}

	// 根据 token 更新 User
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		name, _ := claims["name"].(string)
		comment.User = name
	}

	// 获取文章 ID
	articleId := strings.Split(r.URL.Path, "/")[2]
	comment.ArticleId, err = strconv.ParseInt(articleId, 10, 64)
	if err != nil {
		Response(ResponseMessage{nil,err.Error()}, w, http.StatusBadRequest)
		return
	}

	// 生成评论时间
	comment.Date = fmt.Sprintf("%d-%d-%d", time.Now().Year(), time.Now().Month(), time.Now().Day())

	articles := db.GetArticles(comment.ArticleId, 0)
	// 评论文章不存在
	if len(articles) == 0 {
		Response(ResponseMessage{nil,"Article Not Exists !"}, w, http.StatusBadRequest)
		return
	}

	// 将评论加入数据库
	for i := 0; i < len(articles); i++ {
		articles[i].Comments = append(articles[i].Comments, comment)
	}
    
	err = db.PutArticles(articles)
	if err != nil {
		Response(ResponseMessage{nil,err.Error()}, w, http.StatusBadRequest)
		return
	}

	// 返回评论数据
	Response(ResponseMessage{comment,nil}, w, http.StatusOK)
}

