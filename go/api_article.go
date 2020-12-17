package swagger

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/chuhongwei/BlogServer/source"
)

// GET /articles?page=3
// 包含一个 query parameter: page
func GetArticles(w http.ResponseWriter, r *http.Request) {
	// query parameter
	r.ParseForm()
	pageStr := r.FormValue("page")
	page := int64(1)
	if pageStr != "" {
		page, _ = strconv.ParseInt(pageStr, 10, 64)
	}
	articles := db.GetArticles(-1, page)
	Response(ResponseMessage{articles,nil,}, w, http.StatusOK)
}

// GET /article/1
// 包含一个 path parameter: id
func GetArticleByID(w http.ResponseWriter, r *http.Request) {
	// path parameter
    // 分解 URL
	str := strings.Split(r.URL.Path, "/")
	articleID, err := strconv.ParseInt((str[len(str)-1]), 10, 64)
	fmt.Println(articleID)
	// Article ID 字符串转换为数字失败
	if err != nil {
		Response(ResponseMessage{nil, "Invalid Article ID !"}, w, http.StatusBadRequest)
		return
	}

	articles := db.GetArticles(articleID, 0)
	if len(articles) != 0 {
		Response(ResponseMessage{articles[0],nil,}, w, http.StatusOK)
		return
	} else {
		// Article ID 对应的文章不存在
		Response(ResponseMessage{nil, "Article Not Exists !",}, w, http.StatusNotFound)
		return 
	}
}

