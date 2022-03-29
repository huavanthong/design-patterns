# Introduction
This tutorial help you understand about Factory Method pattern in Creational Patterns

# Question
* [How do you use a enum in Java?]()

# Answer Question
### Enum
```
enum PaymentType { 
	Cash,
	DebitCard;
}
```
More details: [here]()


# Experience
### Purpose for Builder Patterns
Câu hỏi đặt ra là khi nào chúng ta nên sử dụng các Factory Method Pattern trong design của chúng ta.
* Khi ta cần chọn các phương thức xử lý trong logic code của chúng ta.
* Ta phải tạo đối tượng trung gian để có thể refer đến các phương thức mà ta muốn xử lý.
## Go programming
Các step để implement một Factory Methods.
**Step 1:** Tạo các common function cho mấy class sử dụng interface
**Step 2:** Tạo table constant để ta có thể refer đến methods mà ta muốn xử lý
**Step 3:** Tạo các structure để handle cho từng methods
**Step 4:** Implement tất cả methods trong từng handler của nó
**Step 5:** Implement switch case to select the optional method
--------------------------------------------------------------------------------
**Demo Factory Methods Pattern in Go**
```go
type anotherClass struct {
	member1 string
    member2 string
}

// Step 1:
type Methods interface {
	methods1(amount float32) string
	methods2() string
    methods3() anotherClass
}

// Step 2:
type methodType int

const (
	Type1 PaymentType = iota
	Type2
    Type3
    ...
)

// Step 3:
type HandleMethod1 struct{}
type HandleMethod2 struct{}
type HandleMethod3 struct{
    next anotherClass
}

// Step 4:
func (c *HandleMethod1) methods1(amount float32) string { 
    return " "     
}

func (c *HandleMethod1) methods2() string { 
    return " "
}

func (c *HandleMethod1) methods3() anotherClass { 

    return c.next    
}

// Implement another Handler

// Step 5:
func getMethods(t methodType) Methods {
	switch t {
	case Type1:
		return new(HandleMethod1)
	case Type2:
		return new(HandleMethod2)
    case Type3:
		return new(HandleMethod3)
	default:
		return nil
	}
}
```
More details: [here]()

## Java programming
Các step để implement một Factory Methods.
**Step 1:** Sử dụng Interface để tạo các methods trong cung cấp cho các Handle còn lại.
**Step 2:** Implement các Interface trong nó.
--------------------------------------------------------------------------------
**Demo Factory Methods Pattern in Java**
```java
// Step 1:
interface Methods {
    String methods1(float32 amount);
    String methods2();
    anotherClass methods3();
}

class anotherClass {
    private String member1;
    private String member2;
}

// Step 2: 
class HandleMethod1 implements Methods {
    private anotherClass obj;

    @Override
    public String methods1(float32 amount)() {
        ...   
    }

    @Override
    public String methods2() {
        ...
    }

    @Override
    public anotherClass methods3() {
        ..
    }
}

class HandleMethod2 implements Methods {
    private anotherClass obj;

    @Override
    public String methods1(float32 amount)() {
        ...   
    }

    @Override
    public String methods2() {
        ...
    }

    @Override
    public anotherClass methods3() {
        ..
    }
}


```