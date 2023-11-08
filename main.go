package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 允许所有域名的跨域请求
		w.Header().Set("Access-Control-Allow-Origin", "*")

		if r.Method == http.MethodPost {
			idCard := r.FormValue("idCard")
			fullName := r.FormValue("fullName")

			// 这里调用实名认证的函数，例如 verifyIDCard 函数
			result := verifyIDCard(idCard, fullName)

			if result {
				fmt.Fprintf(w, "实名认证成功")
			} else {
				fmt.Fprintf(w, "实名认证失败")
			}
		}
	})
	fmt.Println("真名认证系统已启动")
	http.ListenAndServe(":80", nil)
}

func verifyIDCard(idCard string, fullName string) bool {
	// 实现认证逻辑，返回 true 表示认证通过，否则返回 false
	// 可以使用第三方包来发送HTTP请求至实名认证服务提供商
	// 请根据实际需求来实现这个函数
	// ...
	if idCard == "6666" && fullName == "guoyao" {
		return true
	}

	return false
}
