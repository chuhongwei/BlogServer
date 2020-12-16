package swagger

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/chuhongwei/BlogServer/dal/db"
)

// 包含一个 query parameter: page
func GetArticles(w http.ResponseWriter, r *http.Request) {
	// query parameter
	r.ParseForm()
	pagestr := r.FormValue("page")
	page := int64(1)
	if pagestr != "" {
		page, _ = strconv.ParseInt(pagestr, 10, 64)
	}
	articles := db.GetArticles(-1, page)
	Response(MyResponse{
		articles,
		nil,
	}, w, http.StatusOK)
}

// 包含一个 path parameter: id
func GetArticleID(w http.ResponseWriter, r *http.Request) {
	// path parameter
    // 分解 URL
	s := strings.Split(r.URL.Path, "/")
	articleID, err := strconv.ParseInt((s[len(s)-1]), 10, 64)
	fmt.Println(articleID)
	if err != nil {
		Response(MyResponse{nil, "Not Found"}, w, http.StatusOK)
		return
	}

	articles := db.GetArticles(articleID, 0)
	if len(articles) != 0 {
		Response(MyResponse{
			articles[0],
			nil,
		}, w, http.StatusOK)
	}
}


func GetArticleIDComments(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
    // 分解 URL
	s := strings.Split(r.URL.Path, "/")
	articleID, err := strconv.ParseInt((s[len(s)-2]), 10, 64)
	fmt.Println(articleID)
	if err != nil {
		Response(MyResponse{nil, err.Error()}, w, http.StatusBadRequest)
		return
	}

	articles := db.GetArticles(articleID, 0)
	if len(articles) != 0 && len(articles[0].Comments) != 0 {
		Response(MyResponse{
			articles[0].Comments,
			nil,
		}, w, http.StatusOK)
	} else if len(articles) == 0 {
		Response(MyResponse{nil, "Article Not Found"}, w, http.StatusOK)
	} else {
		Response(MyResponse{nil, nil}, w, http.StatusOK)
	}
}




