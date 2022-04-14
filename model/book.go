package model

import "time"

// Book 书籍表
//type Book struct {
//	ID       int
//	BookName string
//	Author   string
//	Price    float64
//	Sales    int
//	Stock    int
//	ImgPath  string
//}

type Book struct {
	ISBN        int64
	BookTitle   string
	BookAuthor  string
	Publication int
	Publisher   string
	ImageURLL   string
	IsHidden    bool
	Categories  []string
	Timestamp   time.Time
	Labels      []string
	Comment     string
}
