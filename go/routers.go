package swagger

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name           string
	Method        string
	Pattern         string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("go/index.html")
	if err != nil {
		log.Fatal(err)
	}
	index, _ := ioutil.ReadAll(f)
	fmt.Fprintf(w, string(index))
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},

	// Article API
	Route{
		"GetArticles",
		strings.ToUpper("Get"),
		"/articles",
		GetArticles,
	},

	Route{
		"GetArticleByID",
		strings.ToUpper("Get"),
		"/article/{id}",
		GetArticleByID, 
	},
	

	// User API
	Route{
		"PostUserRegister",
		strings.ToUpper("Post"),
		"/user/register",
		PostUserRegister,
	},

	Route{
		"PostUserLogin",
		strings.ToUpper("Post"),
		"/user/login",
		PostUserLogin,
	},

	
	// Comment API
	Route{
		"GetArticleCommentsByID",
		strings.ToUpper("Get"),
		"/article/{id}/comments",
		GetArticleCommentsByID,
	},

	Route{
		"PostArticleCommentsByID",
		strings.ToUpper("Post"),
		"/article/{id}/comment",
		PostArticleCommentsByID,
	},


	// Option API
	Route{
		"Options",
		strings.ToUpper("options"),
		"/article/{id}/comment",
		Options,
	},
}

