/*
Reference:
    https://www.meisternote.com/app/note/qG2friWrbBn5/open-closed-principle-demo

Experience:
    1. Chúng ta phải hiểu rằng, một repos nên xây dựng các hoạt động cơ bản CRUD mà thôi.
    2. Với các hoạt động khác, chẳng hạn SearchByName, SearchByView, SearchByLike ... ta cần phải tách rời nó ra.const
    3. Ngoài ra, việc tách rời ra, giúp ta giữ được các chức năng hoạt động ở repos cũ, mà kết hợp các được các chức năng
       hoạt động mới ở repos mới chẳng hạn như ElasticSearch. Giúp ta query nhanh hơn.
    4. Cuối cùng, luôn nghĩ đến nguyên tắc thiết kế SOLID.
*/
package main

import (
	"fmt"
)

type UserRepository interface {
	Create(user *User) error
	GetByID(id int) (*User, error)
	Update(user *User) error
	Delete(id int) error
}

type User struct {
	ID    int
	Name  string
	Email string
}

type MongoDBRepository struct {
	// implementation
}

func (r *MongoDBRepository) Create(user *User) error {
	// implementation
	fmt.Println("MongoDBRepository: Create user")
	return nil
}

func (r *MongoDBRepository) GetByID(id int) (*User, error) {
	// implementation
	fmt.Println("MongoDBRepository: Get user by ID")
	return &User{ID: id, Name: "John Doe", Email: "john.doe@example.com"}, nil
}

func (r *MongoDBRepository) Update(user *User) error {
	// implementation
	fmt.Println("MongoDBRepository: Update user")
	return nil
}

func (r *MongoDBRepository) Delete(id int) error {
	// implementation
	fmt.Println("MongoDBRepository: Delete user")
	return nil
}

// Exeperience 2: Tách các chức năng nâng cao ra
type UserSearchRepository interface {
	SearchByName(name string) ([]*User, error)
}

func (r *MongoDBRepository) SearchByName(name string) ([]*User, error) {
	// implementation
	fmt.Println("MongoDBRepository: Search user by name")
	return []*User{{ID: 1, Name: "John Doe", Email: "john.doe@example.com"}}, nil
}

// Exeperience 3: Tách ra cho việc sử dụng ElasticSearch
type ElasticSearchRepository struct {
	// implementation
}

func (r *ElasticSearchRepository) SearchByName(name string) ([]*User, error) {
	// implementation
	fmt.Println("ElasticSearchRepository: Search user by name")
	return []*User{{ID: 1, Name: "John Doe", Email: "john.doe@example.com"}}, nil
}

func main() {
	mongoDBRepo := &MongoDBRepository{}
	elasticSearchRepo := &ElasticSearchRepository{}

	user := &User{
		ID:    1,
		Name:  "John Doe",
		Email: "john.doe@example.com",
	}

	// Create user using MongoDB repository
	if err := mongoDBRepo.Create(user); err != nil {
		fmt.Println(err)
	}

	// Get user by ID using MongoDB repository
	if u, err := mongoDBRepo.GetByID(user.ID); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("User retrieved: %+v\n", *u)
	}

	// Update user using MongoDB repository
	if err := mongoDBRepo.Update(user); err != nil {
		fmt.Println(err)
	}

	// Delete user using MongoDB repository
	if err := mongoDBRepo.Delete(user.ID); err != nil {
		fmt.Println(err)
	}

	// Search user using MongoDB repository
	if users, err := mongoDBRepo.SearchByName("John Doe"); err != nil {
		fmt.Println(err)
	} else {
		for _, u := range users {
			fmt.Printf("User retrieved: %+v\n", *u)
		}
	}

	// Search user by name using ElasticSearch repository
	if users, err := elasticSearchRepo.SearchByName("John Doe"); err != nil {
		fmt.Println(err)
	} else {
		for _, u := range users {
			fmt.Printf("User retrieved: %+v\n", *u)
		}
	}
}
