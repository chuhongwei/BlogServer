package swagger

import (
	"log"
	"net/http"
	"encoding/json"
)

// 存储 Response 信息
type ResponseMessage struct {
	OkMessage    interface{} `json:"ok,omitempty"`
	ErrorMessage    interface{} `json:"error,omitempty"`
}

// 发送回复信息
func Response(response interface{}, w http.ResponseWriter, code int) {
	jsonData, err := json.Marshal(&response)

	if err != nil {
		log.Fatal(err.Error())
	}

	w.Header().Set("Access-Control-Allow-Methods","PUT,POST,GET,DELETE,OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Origin, Access-Control-Allow-Credentials, Access-Control-Allow-Methods, Access-Control-Allow-Headers, Authorization, X-Requested-With")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(jsonData)
	w.WriteHeader(code)
}

// 处理 OPTIONS
func Options(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Methods","PUT,POST,GET,DELETE,OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Origin, Access-Control-Allow-Credentials, Access-Control-Allow-Methods, Access-Control-Allow-Headers, Authorization, X-Requested-With")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}