package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)


type Controller struct {

}

type BookInfo struct {
	BookName string
	FinishedTime string
	Comments string
}



func (c Controller) ShowView(w http.ResponseWriter, r *http.Request, tmpl string, data interface{})  {
	//创建路径  输出：web/tpl/templatename
	page := filepath.Join("template", tmpl)

	// 创建模板实例
	resultTemplate, err := template.ParseFiles(page)
	if err != nil {
		fmt.Println("创建模板实例错误: ", err)
		return
	}

	// 融合数据
	err = resultTemplate.Execute(w, data)
	if err != nil {
		fmt.Println("融合模板数据时发生错误", err)
		return
	}


}

func (c Controller) Welcome(w http.ResponseWriter, r *http.Request){

	//读取数据库
	//创建结构体
	//融合数据，调用showView

	allBooks := []BookInfo{
		{"The Great Gatzby", "2005.6.1", "hello"},
		{"Little Prince", "2007.3.2", "romantic"},
	}

	c.ShowView(w, r, "hello.html", struct {
		Books []BookInfo
	}{
		allBooks,
	})

}

func (c Controller) AddBook(w http.ResponseWriter, r *http.Request){
	//1.接收数据
	//2.创建结构体数据，追加到全局切片中
	//3.调用ShowView，传入结构体
}