package main

import (
	"fmt"
	"goLibrary/model"
	"goLibrary/util"
	"strconv"
	"time"
)

func main() {
	//{
	//	sqlStr := "SELECT book_id, best_book_id, average_rating FROM books;"
	//	rows, _ := util.Db.Query(sqlStr)
	//	for rows.Next() {
	//		pi := model.PreItem{}
	//		rows.Scan(&pi.ItemId, &pi.BestId, &pi.Rate)
	//		pi.Rate = pi.Rate[:1]
	//		sqlIns := "INSERT INTO items(item_id, categories, time_stamp, labels, comment) VALUES(?,?,?,?,?)"
	//		sqlCom := "SELECT tag_name FROM booktags WHERE tag_id=? LIMIT 1"
	//		rowc, err := util.Db.Query(sqlCom, pi.BestId)
	//		if err != nil {
	//			panic(err)
	//		}
	//		var str string
	//		for rowc.Next() {
	//			rowc.Scan(&str)
	//		}
	//		lb, _ := json.Marshal(str)
	//		r, _ := json.Marshal(pi.Rate)
	//		_, err = util.Db.Exec(sqlIns, pi.ItemId, string(r), time.Now(), string(lb), str)
	//		if err != nil {
	//			panic(err)
	//		}
	//	}
	//}

	{
		sqlStr := "SELECT book_id, user_id, rating FROM ratings"
		rows, _ := util.Db.Query(sqlStr)
		for rows.Next() {
			fb := model.Feedback{}
			rows.Scan(&fb.ItemId, &fb.UserId, &fb.FeedbackType)
			rate, _ := strconv.Atoi(fb.FeedbackType)
			switch rate {
			case 1:
			case 2:
				fb.FeedbackType = "dislike"
			case 3:
				fb.FeedbackType = "read"
			case 4:
				fb.FeedbackType = "like"
			case 5:
				fb.FeedbackType = "star"
			}

			fmt.Println(fb)
			sqlIns := "INSERT INTO feedback(feedback_type, user_id, item_id, time_stamp, comment) VALUES(?, ?, ?, ?, ?)"
			_, err := util.Db.Exec(sqlIns, fb.FeedbackType, fb.UserId, fb.ItemId, time.Now(), "")
			if err != nil {
				panic(err)
			}
		}
	}

	//router := gin.Default()
	//router.LoadHTMLGlob("views/**/*")
	//router.LoadHTMLFiles("templates/index.html")

	//router.GET("/", func(c *gin.Context) {
	//	var page model.Page
	//	var books []model.Book
	//	sqlStr2 := "select ISBN,BookTitle,BookAuthor,Publication,Publisher,ImageURLL from books limit ?,?"
	//	//rand.Seed(time.Now().UnixNano())
	//	idx := 0
	//
	//	for len(books) < 12 {
	//		rows, _ := util.Db.Query(sqlStr2, idx, int64(12))
	//		idx += 1
	//		for rows.Next() && len(books) < 12 {
	//			book := model.Book{}
	//			rows.Scan(&book.ISBN, &book.BookTitle, &book.BookAuthor, &book.Publication, &book.Publisher, &book.ImageURLL)
	//			fmt.Println(book.ISBN, len(book.BookTitle))
	//			if book.ISBN > 0 && len(book.BookTitle) > 0 {
	//				resp, err := http.Get(book.ImageURLL)
	//				if err == nil {
	//					pix, _ := ioutil.ReadAll(resp.Body)
	//					if len(pix) > 100 && cap(pix) > 100 {
	//						fmt.Println(len(pix))
	//						books = append(books, book)
	//					} else {
	//						fmt.Println(pix)
	//					}
	//				}
	//				resp.Body.Close()
	//			}
	//		}
	//	}
	//
	//	page = model.Page{
	//		Books: books,
	//	}
	//	c.HTML(http.StatusOK, "index.html", page)
	//})

	//router.GET("/clear", func(c *gin.Context) {
	//	for {
	//		sqlStr := "select ISBN,BookTitle,BookAuthor,Publication,Publisher,ImageURLL from books"
	//		rows, _ := util.Db.Query(sqlStr)
	//		for rows.Next() {
	//			flag := true
	//			book := model.Book{}
	//			rows.Scan(&book.ISBN, &book.BookTitle, &book.BookAuthor, &book.Publication, &book.Publisher, &book.ImageURLL)
	//			if len(book.BookTitle) == 0 {
	//				//DELETE FROM table_name WHERE some_column=some_value;
	//				dStr := "DELETE FROM books WHERE BookTitle=?"
	//				util.Db.Exec(dStr, book.BookTitle)
	//				fmt.Println(book.BookTitle)
	//				flag = false
	//			}
	//			resp, err := http.Get(book.ImageURLL)
	//			pix, _ := ioutil.ReadAll(resp.Body)
	//			if err != nil || len(pix) < 100 {
	//				dStr := "DELETE FROM books WHERE ISBN=?"
	//				util.Db.Exec(dStr, book.ISBN)
	//				fmt.Println(book.BookTitle)
	//				flag = false
	//			}
	//			resp.Body.Close()
	//			if flag {
	//				fmt.Println(book.BookTitle)
	//			}
	//		}
	//	}
	//	fmt.Println("DONE!!!")
	//})

	//router.GET("/login", func(c *gin.Context) {

	//})

	//router.Run("localhost:9999")
}
