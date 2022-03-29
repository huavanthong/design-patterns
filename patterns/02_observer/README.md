# Introduction
This tutorial help you understand about Observe pattern in Behavioral Patterns

# Question
* [Why we consider Youtube channels is a subject in ex1_youtube_sub.go?]()
* []


# Answer Question

# Experience
### Intent
* Define a one-to-many dependency between objects so that when one object changes state, all its dependents are notified and updated automatically.
* Encapsulate the core (or common or engine) components in a Subject abstraction, and the variable (or optional or user interface) components in an Observer hierarchy.
* The "View" part of Model-View-Controller.

## Go programming

## Java programming
Các bước để implement Observer Pattern như sau:
**Step 1:** Tạo các common function cho mấy chức năng khác implement nó.
**Step 2:** Tạo table constant để ta có thể refer đến methods mà ta muốn xử lý
**Step 3:** Tạo các structure để handle cho từng methods
**Step 4:** Implement tất cả methods trong từng handler của nó
**Step 5:** Implement switch case to select the optional method

-----------------------------------------------------------------------