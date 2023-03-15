/*
Reference:
    https://www.meisternote.com/app/note/qG2friWrbBn5/open-closed-principle-demo

Experience:
    1. Luôn define các phương thức đặc thù để thực hiện các hoạt động CRUD cơ bản nhất.
	2. Việc xây dựng các phương thức đó phải độc lập và không phụ thuộc vào các
	   => Trong cả hai trường hợp, BookRepository đều giữ nguyên, không cần thay đổi,
	   nhưng triển khai của nó được thay đổi để phù hợp với cơ sở dữ liệu cụ
*/

package main

import (
	"database/sql"

	"go.mongodb.org/mongo-driver/mongo"
)

type BookRepository interface {
	Create(book *Book) error
	Read(id string) (*Book, error)
	Update(book *Book) error
	Delete(id string) error
}

type Book struct {
	ID    int
	Title string
}

// Ví dụ, nếu bạn muốn sử dụng cơ sở dữ liệu MySQL, bạn có thể triển khai BookRepository như sau:
type MySQLBookRepository struct {
	db *sql.DB
}

func NewMySQLBookRepository(db *sql.DB) *MySQLBookRepository {
	return &MySQLBookRepository{db: db}
}

func (r *MySQLBookRepository) Create(book *Book) error {
	// implementation

	return nil
}

func (r *MySQLBookRepository) Read(id string) (*Book, error) {
	// implementation

	return nil, nil
}

func (r *MySQLBookRepository) Update(book *Book) error {
	// implementation

	return nil
}

func (r *MySQLBookRepository) Delete(id string) error {
	// implementation

	return nil
}

// Nếu bạn muốn sử dụng cơ sở dữ liệu MongoDB thay vì MySQL, bạn có thể triển khai BookRepository với MongoDB như sau:
type MongoDBBookRepository struct {
	collection *mongo.Collection
}

func NewMongoDBBookRepository(collection *mongo.Collection) *MongoDBBookRepository {
	return &MongoDBBookRepository{collection: collection}
}

func (r *MongoDBBookRepository) Create(book *Book) error {
	// implementation

	return nil
}

func (r *MongoDBBookRepository) Read(id string) (*Book, error) {
	// implementation

	return nil, nil
}

func (r *MongoDBBookRepository) Update(book *Book) error {
	// implementation

	return nil
}

func (r *MongoDBBookRepository) Delete(id string) error {
	// implementation

	return nil
}

func main() {

}
