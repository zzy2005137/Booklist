package controller

import (
	"fmt"
	"github.com/zzy2005137/booklist/model"
	"html/template"
	"net/http"
	"path/filepath"
)

type Controller struct {
	M		model.BooklistModel
	Init     bool
}

func (c *Controller) ShowView(w http.ResponseWriter, r *http.Request, tmpl string, data interface{}) {
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

func (c *Controller) Welcome(w http.ResponseWriter, r *http.Request,) {

	//读取数据库
	books := c.M.QueryBooks()
	//融合数据，调用showView

	c.ShowView(w, r, "hello.html", struct {
		Books []model.BookInfo
	}{
		books,
	})

}

func (c *Controller) AddBookView(w http.ResponseWriter, r *http.Request){
	c.ShowView(w, r, "add.html", nil)
}

func (c *Controller) AddBook(w http.ResponseWriter, r *http.Request) {
	//1.接收数据

	BookName := r.FormValue("BookName")
	FinishedTime := r.FormValue("FinishedTime")
	Comments := r.FormValue("Comments")
	flag := false


	//2.追加数据
	if BookName != "" {
		flag = true
		c.M.AddBook(BookName, FinishedTime, Comments)

		//3.传入结构体
		c.ShowView(w, r, "add.html", struct {
			Flag  bool
		}{
			flag,
		})
	}else {
		c.ShowView(w, r, "add.html", nil)
	}


}
