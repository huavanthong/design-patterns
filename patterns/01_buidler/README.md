# Introduction
This tutorial help you understand about Builder pattern in Creational Patterns

# Question
* [Difference between object and instance?](#object-and-instance)
* [How do you create a object in Go/Java/Python using Buidler Pattern?](#methods-initialize-a-object)
* [As you know, Go don't support OOP, how do you create a class from Go?](#class-in-gojavapython)

# Answer Question
### Object and Instance
- Objects are the results of instantiating a class. 
- Instance is a memory block, which contains the reference to an object. 

### Methods initialize a object
Go programming
```go
func newClassBuilder() classBuilder {
	return &myclass{}
}
```

Java programming
```java
public myclass build() {
    return new myclass(this);
}
```
### Class in Go/Java/Python
In Go programming, we use struct to define a class
# Experience
### Purpose for Builder Patterns
Câu hỏi đặt ra là khi nào chúng ta nên sử dụng các Builder Pattern trong design của chúng ta.
* Khi ta cần khởi tạo một class với nhiều paramters
* Khi ta muốn quản lý việc khởi tạo một class với các parameters theo mục đích cụ thể.
## Go programming
Đối với Go programming, để tạo Builder Pattern, thì ta cần phải hiểu về keyword: "Interface".  
Interface trong Go, giúp cho ta có thể implement được các methods theo các bước sau:

**Step 1:** Tạo một structure với các members
**Step 2:** Tạo các structure buidler() cho myclass của ta
**Step 3:** Implement các methods trong classBuilder 
**Step 4:** Phải tạo một methods để initialize a object Builder
--------------------------------------------------------------------------------
**Demo Builder Pattern in Go**
```go
// Step 1:
type myclass struct {
    member1 int 
    member2 string
}

// Step 2:
type classBuilder interface {
    init_methods_1(a int, b int) classBuilder
    init_methods_2(c int) classBuilder
    ...
	Build() myclass
}

// Step 3:
func (p *nutritionFacts) init_methods_1(a int, b int) classBuilder {
    p.member1 = a
    p.member2 = b
    return p
}

func (p *nutritionFacts) Build() myclass {
    return p
}

// Step 4:
func newClassBuilder() classBuilder {
	return &myclass{}
}

```

## Java programming
Đối với Java programming, để tạo Builder Pattern, ta phải hiểu cách tạo một class Builder nằm trong class lớn.  
**Step 1:** Tạo một myclass chứa các member cần thiết.
**Step 2:** Tạo classBuilder nằm trong myclass của ta để tạo các method builder.
**Step 3:** Implement các methods builder
**Step 4:** Phải tạo một method để tạo new object myclass từ classBuilder
--------------------------------------------------------------------------------
**Demo Builder Pattern in Java**
```java
public myclass {
    private int member1;
    private int member2;
    ...

    public classBuilder{
        private int member1;
        private int member2;
        ...

        public classBuilder init_methods_1(int a, int b) {
            this.member1 = a;
            this.member2 = b;
        }

        public classBuilder init_methods_2(int c) {

        }

        public myclass build() {
            return new myclass(this);
        }
    }

    public myclass(classBuilder builder) {
        this.member1 = builder.member1;
        this.member2 = builder.member2;
        ...
    }

}
```

## Python programming

```python

```