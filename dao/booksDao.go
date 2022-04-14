package dao

import (
	"fmt"
	"github.com/disintegration/imageorient"
	"goLibrary/model"
	"goLibrary/util"
	"math/rand"
	"net/http"
	"time"

	_ "image/jpeg"
	_ "image/png"
)

func GetAllBook(pageSize int) (model.Page, error) {
	// books 每页查询出来的数据
	var books []model.Book
	sqlStr2 := "select ISBN,BookTitle,BookAuthor,Publication,Publisher,ImageURLL from books limit ?,?"

	//for len(books) >= 12 {
	rand.Seed(time.Now().UnixNano())
	rows, err := util.Db.Query(sqlStr2, rand.Intn(99999), int64(pageSize))
	if err != nil {
		return model.Page{}, err
	}
	//cnt := 0
	for rows.Next() {
		//cnt -= 1
		book := model.Book{}
		rows.Scan(&book.ISBN, &book.BookTitle, &book.BookAuthor, &book.Publication, &book.Publisher, &book.ImageURLL)
		books = append(books, book)
	}
	fmt.Println(books)
	for i, b := range books {
		if b.ISBN <= 1 || len(b.BookTitle) <= 1 {
			//cnt += 1
			books = append(books[:i], books[i+1:]...)
		}

		resp, _ := http.Get(b.ImageURLL)
		//defer resp.Body.Close()
		//if (resp.Body)
		_, str, err := imageorient.Decode(resp.Body)
		if err != nil {
			books = append(books[:i], books[i+1:]...)
		}
		if len(str) <= 3 {
			fmt.Println("<><<><><<>><<>><")
		} else {
			fmt.Println(str[0:3])
		}
		//if err != nil {
		//	img.
		//	fmt.Println(img)
		//	panic("err")
		//}
	}
	//}

	return model.Page{
		Books: books,
	}, nil
}
