# Overview
Để hiểu các kinh nghiệm coding trong việc lập trình trong dự án này. Chúng ta có thể Find all tất cả các kinh nghiệm qua keyword: experience
# Project Structure
This is a project structure for overview our architecture
```lua
├── common/
│   ├── common.go
├── entity/
│   ├── rectangle.go
│   └── circle.go
│
├── boundary/
│   ├── rectangle_input.go
│   ├── rectangle_output.go
│   ├── circle_input.go
│   ├── circle_output.go
│   ├── rectangle_boundary.go
│   └── circle_boundary.go
│
├── interactor/
│   ├── rectangle_interactor.go
│   └── circle_interactor.go
│
├── builder/
│   ├── shape_builder.go
│   ├── rectangle_builder.go
│   └── circle_builder.go
│
├── repository/
│   ├── rectangle_repository.go
│   └── circle_repository.go
│
├── go.mod
└── main.go
```
Project Structure là một kiến trúc phần mềm theo mô hình Clean Architecture. Nó được chia thành các lớp phân tách với nhau nhằm đảm bảo tính rõ ràng và dễ bảo trì.

* Lớp "common" bao gồm các file chung mà được sử dụng bởi các phần khác của project. Lớp "entity" định nghĩa các đối tượng cốt lõi của ứng dụng, bao gồm hình chữ nhật và hình tròn.

* Lớp "boundary" định nghĩa các lớp đầu vào và đầu ra cho các thực thể. Các lớp này được sử dụng để tách các yêu cầu vào và ra khỏi ứng dụng. Mỗi hình chữ nhật và hình tròn đều có một lớp đầu vào và đầu ra tương ứng.

* Lớp "interactor" chứa các file thực hiện nghiệp vụ và logic xử lý cho các hình. Lớp này bao gồm các tầng use cases được xử lý bởi các interactor khác nhau, ví dụ như tạo hình chữ nhật mới hoặc tính diện tích của hình tròn.

* Lớp "builder" định nghĩa các builder để tạo ra các đối tượng trong ứng dụng, bao gồm builder cho hình chữ nhật và hình tròn.

* Lớp "repository" chứa các file tương tác với cơ sở dữ liệu. Các file này được sử dụng để lưu trữ và truy xuất các đối tượng trong ứng dụng, bao gồm hình chữ nhật và hình tròn.  

Tổng quan về kiến trúc này là tách biệt rõ ràng giữa các lớp và phân chia các trách nhiệm giữa chúng. Kiến trúc này giúp cho code dễ bảo trì, dễ phát triển và có khả năng tái sử dụng cao.

# Application Purpose

# Run application
Để build application
```powershell
go build -o app
```

Để run application
```powershell
./app.exe
```
# Run test
Để run test tại mỗi layer:
```powershell
go test ./[Target Component]

Ex: 
go test ./repository
```

Để in tất cả các test case kèm theo trạng thái của chúng (pass/fail) và thời gian thực thi
```powershell
go test -v ./[Target Component]
```


Để in ra tỉ lệ code được bao phủ (code coverage) bởi các unit test
```powershell
go test -cover ./[Target Component]
```


Để in ra tỉ lệ code được bao phủ và lưu vào một file coverage.out
```powershell
go test -coverprofile=coverage.out ./[Target Component]
```


Để mở file coverage.out với trình duyệt web để xem các vùng code được bao phủ và không được bao phủ của ứng dụng của bạn.
```powershell
 go tool cover -html=coverage.out ./[Target Component]
```
