package main

import (
	"github.com/zzy2005137/booklist/controller"
	"net/http"
)




func main() {

	controller := controller.Controller{}
	controller.InitBooks()

	//创建服务器实例，指定静态文件所在路径
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/",http.StripPrefix("/static/", fs))

	http.HandleFunc("/", controller.Welcome)
	http.HandleFunc("/add.html", controller.AddBookView)
	http.HandleFunc("/add", controller.AddBook)

	http.ListenAndServe(":8080",nil)


}

