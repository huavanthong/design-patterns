package entity

/*
Experience 5: Kế thừa trong Golang được thực hiện thế nào? Việc kế thừa giúp ích được gì?

Nếu chúng ta muốn tạo những method dùng chung cho tất cả các hình học khác. Ta cần phải khai báo kiểu interface{}.
Việc này giúp tacó thể:
	+ khởi tạo các hình học khác nhau trong một kiểu Shape.
	+ các hình học muốn trở thành con của Shape cần phải implement tất cả method mà Shape define.

Refer: shape_test.go để hiểu rõ test case trong việc kế thừa.
*/

type Shape interface {
	GetArea() float64
	GetPerimeter() float64
	Draw() string
}
