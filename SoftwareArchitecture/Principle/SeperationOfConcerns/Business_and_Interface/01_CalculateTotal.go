package main

import "fmt"

/****************************************** Business Logic *********************************************/
type Product struct {
	ID    string
	Name  string
	Price float64
}

type ShoppingCart struct {
	items []Product
}

func (cart *ShoppingCart) addItem(item Product) {
	cart.items = append(cart.items, item)
}

/*
Introduction:
Cụ thể, đoạn mã này sử dụng vòng lặp for range để duyệt qua từng phần tử (item) và chỉ số (index) của mỗi phần tử trong slice cart.items.
Với mỗi phần tử, nó kiểm tra xem ID của phần tử đó có bằng với ID của mặt hàng cần xóa hay không.
Nếu có, nó sử dụng hàm append để xóa phần tử đó khỏi slice cart.items.

Cụ thể, đoạn mã sử dụng một kỹ thuật để xóa một phần tử khỏi slice trong Go như sau:

    * cart.items[:i] trả về một slice mới bao gồm tất cả các phần tử từ đầu slice (vị trí 0) đến trước phần tử cần xóa (vị trí i).
    * cart.items[i+1:] trả về một slice mới bao gồm tất cả các phần tử từ sau phần tử cần xóa (vị trí i+1) đến cuối slice.
    * append(cart.items[:i], cart.items[i+1:]...) kết hợp hai slice này lại và trả về slice mới, không bao gồm phần tử được xóa.

Sau khi xóa phần tử khỏi slice, đoạn mã sử dụng từ khóa break để thoát khỏi vòng lặp, vì nếu đã tìm thấy phần tử cần xóa,
không cần phải duyệt qua các phần tử khác nữa.

Fix bug:
	cannot use cart.items[i + 1] (variable of type Product) as []Product value in argument to appendcompilerIncompatibleAssign
```
	for i, v := range cart.items {
		if v.ID == item.ID {
			cart.items = append(cart.items[:i], cart.items[i+1:]...)
			break
		}
	}
````
*/
func (cart *ShoppingCart) removeItem(item Product) {

	for i, v := range cart.items {
		if v.ID == item.ID {
			cart.items = append(cart.items[:i], append([]Product{cart.items[i+1]}, cart.items[i+2:]...)...)
			break
		}
	}
}

func (cart *ShoppingCart) calculateTotal() float64 {

	var total float64

	for _, v := range cart.items {
		total += v.Price
	}

	return total
}

/****************************************** Interface User *********************************************/
type ShoppingCartView struct {
	cart *ShoppingCart
}

func (view *ShoppingCartView) displayItems() {

	fmt.Println("Your shopping cart contains:")
	for _, v := range view.cart.items {
		fmt.Printf("- %s ($%.2f)\n", v.Name, v.Price)
	}
}

func (view *ShoppingCartView) displayTotal() {
	fmt.Printf("Total: $%.2f\n", view.cart.calculateTotal())
}
