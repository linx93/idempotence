package idempotence

import (
	"net/http"
)

var defaultGetTokenUrl string = "/getToken"

func GetTokenHandler(w http.ResponseWriter, r *http.Request) {

}

//func Token(getTokenUrl string) {
//	if getTokenUrl == "" {
//		getTokenUrl = defaultGetTokenUrl
//	}
//	http.HandleFunc(getTokenUrl, GetTokenHandler)
//}
