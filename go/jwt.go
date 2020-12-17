package swagger

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

const jwtSecret = "serviceHW9"

// 生成 token
func SignToken(userName, password string) (string, error) {

	nowTime := time.Now()
	expireTime := nowTime.Add(time.Hour * time.Duration(1))

	claims := make(jwt.MapClaims)
	claims["name"] = userName
	claims["pwd"] = password
	claims["iat"] = nowTime.Unix()
	claims["exp"] = expireTime.Unix()
	
	// crypto.Hash 方案
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	//  该方法内部生成签名字符串，再用于获取完整、已签名的 token
	return token.SignedString([]byte(jwtSecret))
}

// 验证 token
func CheckToken(w http.ResponseWriter, r *http.Request) (*jwt.Token, bool) {
	// 用户登录请求取出 token
	token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecret), nil
		})
	
	if (err == nil && token.Valid) {
		return token, true
	}
	
	w.WriteHeader(http.StatusUnauthorized)

	if err != nil {
		fmt.Fprint(w, "Authorized access to this resource iss invalid !")
	}

	if !token.Valid {
		fmt.Fprint(w, "Token is invalid !")
	}

	return token, false
}