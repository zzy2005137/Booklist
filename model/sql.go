package model

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"

)

type BookInfo struct {
	BookName string
	FinishedTime string
	Comments string
}

type BooklistModel struct {
	DB *sql.DB
}


func (m *BooklistModel) Init()  {

	db, err := sql.Open("mysql",
		"root:1230@tcp(localhost:3306)/booklist")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		// do something here
		fmt.Println("booklist数据库连接失败")
	} else {
		fmt.Println("booklist数据库连接成功")
	}

	m.DB = db
}

func (m *BooklistModel) QueryBooks() []BookInfo{

	//1、定义接收变量 和 存储多行数据的结构体切片
	var (
		BookName string
		FinishedTime string
		Comments string
	)
	books := make([]BookInfo,0)

	//2. m.DB.Query()
	rows, err := m.DB.Query("select BookName, FinishedTime, Comments from allbooks ")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()  //important!

	//3. 将查询结果赋值给变量，注意check error
	for rows.Next(){
		err = rows.Scan(&BookName, &FinishedTime, &Comments)
		if err != nil {
			log.Fatal(err)
		}
		//赋值
		books = append(books, BookInfo{BookName, FinishedTime, Comments})

	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return books

}

func (m *BooklistModel) AddBook(BookName, FinishedTime, Comments string) {


	stmt, err := m.DB.Prepare("INSERT INTO allbooks(BookName, FinishedTime, Comments) VALUES(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(BookName, FinishedTime, Comments)
	if err != nil {
		log.Fatal(err)
	}


}
