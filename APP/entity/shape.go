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

/* Experience 6: SOLID principle
Nguyên lý Thay thế Liskov phát biểu như sau:
“Lớp D được gọi là kế thừa từ lớp B khi và chỉ khi với mọi hàm F thao tác trên các đối tượng của B,
cách cư xử (behavior) của F không đổi khi thay thế các đối tượng của B bằng các đối tượng của D”.

Vì vậy ở đây, ta thấy rằng, nếu muốn một entity nào được gọi là class con của class Shape này, thì
câc entiy con đó phải implement hết tất cả các phương thức đó.
====> điều này tuân theo nguyên tắc Liskov Subtitution Principle
*/
