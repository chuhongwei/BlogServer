package swagger

import (
	"encoding/json"
	"net/http"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/chuhongwei/BlogServer/dal/db"
	"github.com/chuhongwei/BlogServer/dal/model"
	"github.com/dgrijalva/jwt-go"
)

// 包含 body parameter: userName 和 password
func PostUserRegister(w http.ResponseWriter, r *http.Request) {

	db.Init()

	var user model.User

	// body parameter
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		Response(MyResponse{
			nil,
			"parameter error",
		}, w, http.StatusBadRequest)
		return
	}

	check := db.GetUser(user.Username)

	if check.Username != "" {
		Response(MyResponse{
			nil,
			"username existed",
		}, w, http.StatusBadRequest)
		return
	}

	err = db.PutUsers([]model.User{user})

	if err != nil {
		Response(MyResponse{
			nil,
			err.Error(),
		}, w, http.StatusBadRequest)
		return
	}
	tokenString, err := GenerateToken(user.Username, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Error while generate the token")
		tokenString = "generate token error"
	}
	Response(MyResponse{
		tokenString,
		nil,
	}, w, http.StatusOK)
}

// 包含 body parameter: userName 和 password, 验证成功返回一个 token 作为该登陆用户的验证信息
func PostUserLogin(w http.ResponseWriter, r *http.Request) {

	db.Init()

	var user model.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		Response(MyResponse{
			nil,
			"parameter error",
		}, w, http.StatusBadRequest)
		return
	}

	check := db.GetUser(user.Username)

	if check.Username != user.Username || check.Password != user.Password {
		Response(MyResponse{
			nil,
			"username or password error",
		}, w, http.StatusBadRequest)
		return
	}

	tokenString, err := GenerateToken(user.Username, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Error while generate the token")
		tokenString = "generate token error"
	}

	Response(MyResponse{
		tokenString,
		nil,
	}, w, http.StatusOK)
}

// 包含 path parameter id, body parameter: 用户验证 token 和 评论内容
func PostArticleIDComment(w http.ResponseWriter, r *http.Request) {
	db.Init()
	token, isValid := CheckToken(w, r)
	if isValid == false {
		Response(MyResponse{
			nil,
			"authentication fail",
		}, w, http.StatusBadRequest)
		return
	}

	var comment model.Comment

	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		Response(MyResponse{
			nil,
			err.Error(),
		}, w, http.StatusBadRequest)
		return
	}

	if v, ok := token.Claims.(jwt.MapClaims); ok {
		name, _ := v["name"].(string)
		comment.User = name
	}

	articleId := strings.Split(r.URL.Path, "/")[2]
	comment.ArticleId, err = strconv.ParseInt(articleId, 10, 64)
	if err != nil {
		Response(MyResponse{
			nil,
			err.Error(),
		}, w, http.StatusBadRequest)
		return
	}

	comment.Date = fmt.Sprintf("%d-%d-%d", time.Now().Year(), time.Now().Month(), time.Now().Day())

	articles := db.GetArticles(comment.ArticleId, 0)

	if len(articles) == 0 {
		Response(MyResponse{
			nil,
			"articles not found",
		}, w, http.StatusBadRequest)
		return
	}

	for i := 0; i < len(articles); i++ {
		articles[i].Comments = append(articles[i].Comments, comment)
	}

	err = db.PutArticles(articles)

	if err != nil {
		Response(MyResponse{
			nil,
			err.Error(),
		}, w, http.StatusBadRequest)
		return
	}

	Response(MyResponse{
		comment,
		nil,
	}, w, http.StatusOK)
}