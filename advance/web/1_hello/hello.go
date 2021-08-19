/*
建立簡單頁面

這邊利用 Golang 提供的 net/http 套件，
這邊透過此套件簡單建立一個可以執行的網站，

在瀏覽器輸入以下，嘗試訪問建立的頁面
http://localhost:9090

在瀏覽器輸入以下，可以在編輯器上的輸出檢視網站資訊
http://localhost:9090/?url_long=111&url_long=222

另外此套件也可以很簡單的對路由、靜態檔案、模板、Cookie 等資料進行設定和操作。
*/
package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析參數，預設是不會解析的
	fmt.Println(r.Form) //這些資訊是輸出到伺服器端的列印資訊
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") //這個寫入到 w 的是輸出到客戶端的
}

func main() {
	http.HandleFunc("/", sayhelloName)       //設定存取的路由
	err := http.ListenAndServe(":9090", nil) //設定監聽的埠
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
