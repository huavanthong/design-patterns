# Introduction
Entity–Boundary–Interactor
==========================

EBI is a modern application architecture. Its goal is to separate business logic from presentation
logic. With an EBI architecture, by creating different delivery mechanisms, the same business logic
can be used on many platforms.

This repo is just the source for the website. 

.. class:: bold

Go to the official documentation_
---------------------------------

# Project structure 
```lua
your-project-name/
├──/cmd
│  ├── app
|  |  |--main.go
|--/internal
|  |--/entity
|  |  |--entity.go
|  |--/boundary
|  |  |--boundary.go
|  |--/interactor
|  |  |--interactor.go
|--/pkg
|  |--/database
|  |  |--database.go
|  |--/utils
|  |  |--utils.go
|--/config
|  |--config.go
|--/vendor
|--go.mod
|--go.sum
```

Sau đó, ta đã hỏi ChatGPT lại lần nữa. Và nó đưa ra project structure như sau:
```lua
myapp/
├── app/
│   ├── boundary/
│   │   ├── boundary.go
│   │   └── boundary_test.go
│   ├── entity/
│   │   ├── entity.go
│   │   └── entity_test.go
│   ├── interactor/
│   │   ├── interactor.go
│   │   └── interactor_test.go
│   └── main.go
├── config/
│   └── config.go
├── repository/
│   ├── repository.go
│   └── repository_test.go
├── go.mod
└── go.sum

```
Và ta đã hỏi tại sao lại không để  tất cả các code vào trong folder internal như câu trả lời trước???  
ChatGPT đã trả lời:
```
Thư mục internal/ thường được sử dụng để chứa code bên trong một module, nhằm hạn chế sự truy cập vào code bên trong module từ bên ngoài. Tuy nhiên, khi sử dụng kiến trúc EBI, chúng ta thường muốn tách biệt các lớp Boundary, Entity và Interactor để giảm sự phụ thuộc giữa chúng, giúp code trở nên dễ dàng bảo trì và mở rộng hơn. Việc đặt các file của từng lớp trong các thư mục riêng biệt giúp người phát triển dễ dàng tìm kiếm và thay đổi code của từng lớp mà không ảnh hưởng đến các lớp khác.

Ngoài ra, việc sử dụng thư mục internal/ cũng có thể gây ra một số vấn đề trong quá trình phát triển và quản lý module. Ví dụ, các dependency sẽ không thể truy cập vào code trong thư mục internal/, điều này có thể gây khó khăn trong việc quản lý dependency của project.

Vì vậy, khi sử dụng kiến trúc EBI, thường không cần sử dụng thư mục internal/. Thay vào đó, chúng ta có thể sử dụng cơ chế access modifier của Golang (public và private) để quản lý quyền truy cập vào các thành phần của module.
```

# Getting Started


# Ideas
Dưới đây là một số ý tưởng để phát triển đoạn code trên:

1. Thêm đa dạng các hình học khác vào hệ thống, ví dụ như tam giác, hình tròn, hình elip, v.v.
2. Tách biệt lớp Entity ra khỏi các gói Boundary và Interactor, để tăng tính module và dễ dàng mở rộng.
3. Tạo một giao diện đồ họa để người dùng có thể nhập kích thước các hình học và xem kết quả tính toán.
4. Thêm một lớp trung gian trước khi đưa dữ liệu vào lớp Boundary để kiểm tra tính hợp lệ của dữ liệu nhập vào trước khi xử lý.
5. Tạo một hệ thống lưu trữ để lưu trữ các kết quả tính toán trước đó và cho phép người dùng xem lại lịch sử tính toán.
6. Sử dụng các thư viện đồ họa để tạo ra biểu đồ và hình ảnh của các hình học được tính toán.
7. Thêm các tính năng phức tạp hơn, ví dụ như tính toán diện tích và chu vi cho các hình học 3D hoặc các hình dạng đặc biệt khác.
8. Tạo một hệ thống phân quyền để chỉ cho phép người dùng đăng nhập và sử dụng các tính năng của ứng dụng theo vai trò và quyền hạn của họ.
9. Cải thiện hiệu suất của hệ thống bằng cách sử dụng các kỹ thuật tối ưu hoá mã nguồn, tối ưu hoá cơ sở dữ liệu và sử dụng công nghệ cache.
10. Thêm các tính năng liên quan đến giao tiếp với hệ thống bên ngoài, ví dụ như tích hợp với API của bên thứ ba hoặc tạo một  RESTful API để cung cấp dịch vụ tính toán hình học cho ứng dụng khác.


