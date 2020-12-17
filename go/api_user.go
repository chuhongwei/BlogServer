package swagger

import (
	"encoding/json"
	"net/http"
	"fmt"

	"github.com/chuhongwei/BlogServer/source"
)

// POST /user/register
// 包含 body parameter: userName 和 password
func PostUserRegister(w http.ResponseWriter, r *http.Request) {

	db.Init()
	var user db.User

	// 读取参数，转化为 User 结构体
	// body parameter
	err := json.NewDecoder(r.Body).Decode(&user)
	// 参数输入错误
	if err != nil {
		Response(ResponseMessage{nil,"Parameter Error !",}, w, http.StatusBadRequest)
		return
	}
	
	// userName 已存在
	name := db.GetUser(user.Username).Username
	if name != "" {
		Response(ResponseMessage{nil,"UserName has existed !",}, w, http.StatusBadRequest)
		return
	}

	// 将 user 写入数据库
	err = db.PutUsers([]db.User{user})
	if err != nil {
		Response(ResponseMessage{nil,err.Error(),}, w, http.StatusBadRequest)
		return
	}
	
	// 产生用户认证
	tokenStr, err := SignToken(user.Username, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Error when sign the token !")
		Response(ResponseMessage{"Sign token error !",nil,}, w, http.StatusOK)
	}

	Response(ResponseMessage{tokenStr,nil,}, w, http.StatusOK)
}

// POST /user/login
// 包含 body parameter: userName 和 password, 验证成功返回一个 token 作为该登陆用户的验证信息
func PostUserLogin(w http.ResponseWriter, r *http.Request) {

	db.Init()
	var user db.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		Response(ResponseMessage{nil,"Parameter Error !",}, w, http.StatusBadRequest)
		return
	}

	// 从数据库中获得用户名对应的 User 信息
	name := db.GetUser(user.Username).Username
	password := db.GetUser(user.Username).Password
	// 验证用户名与密码是否对应
	if name != user.Username || password != user.Password {
		Response(ResponseMessage{nil,"UserName or Password Error !",}, w, http.StatusBadRequest)
		return
	}

	// 产生用户认证
	tokenStr, err := SignToken(user.Username, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Error when sign the token !")
		Response(ResponseMessage{"Sign token error !",nil,}, w, http.StatusOK)
	}

	Response(ResponseMessage{tokenStr,nil,}, w, http.StatusOK)
}

