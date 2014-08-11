package main

import (
	"NetAppealIntelligence/controllers"
	"fmt"
	"log"
	"net/http"
	//"strings"
)

func main() {
	http.HandleFunc("/", controllers.SafeHandler(controllers.DisplayIndex)) //设置访问的路由
	http.HandleFunc("/listall", controllers.SafeHandler(controllers.ListHandler))
	//http.HandleFunc("/upload", controllers.SafeHandler(controllers.UploadHandler))
	http.HandleFunc("/search", controllers.SafeHandler(controllers.SearchHandler))
	http.HandleFunc("/searchlist", controllers.SafeHandler(controllers.SearchListHandler))
	http.Handle("/uploads/", http.FileServer(http.Dir(".")))
	fmt.Println(http.Dir("."))
	http.HandleFunc("/view", controllers.SafeHandler(controllers.ViewHandler))

	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
