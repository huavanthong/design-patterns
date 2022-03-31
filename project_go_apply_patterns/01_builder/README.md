# Reference
More details: [here](https://golangbyexample.com/builder-pattern-golang/)

# Introduction
We familiarized with Builder pattern inside a class. Refer: [ex1_nutrifacts](https://github.com/huavanthong/design-patterns/blob/master/patterns/01_buidler/ex1_nutrifacts.go)  
However, you need to think that the scope of Builder pattern affect to object. In the above example, we can see: 
* We create a object from nutritionFacts class, typically, it is a small scope.
* Then, we implement Builder pattern inside a class nutritionFacts. This Builder class will help you create a object following your options.
Right now, in this example, we will learn about a large scope of object. For a example, this is a House class.  
For the first words, I want to take a look about class House. This is a big class, thinking about this class, you can see, we can build this house will many factors. Therefore, if you want to develop a perfect software, you need to become expert in construction.
# Questions


# Analyze Requirement
```
        What things do we need to implement a house?
```
### Director
If you want to a director to control the construction of the house in your hand. Let's create a class to control it. 

### Hous requirement
A house is implemented following requirement:
* Window type
* Door type
* Floor
* Roof type
* Walll pannel
* ....

### Concrete
For each above item, we need to implement many type of concretes for this item:  
* Normal builder: We want to build a normal house with a simple supplies.
* Igloo builder: We want to build a wooden house.
* Brick builder: We want to build a brick house.
* Cement builder: We want to build a house out of cement.
* ...

## Builder pattern
From this analysis, we think about Builder pattern to achieve your target. This is a time we summary when we need to use this pattern.
### Purpose


# Getting Started
Right now, please investigate about Builder pattern example in this project.  

To build this project.
```
go build
```

To execute
```
>.\builder.exe
```