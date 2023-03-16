# Introduction
This is a example to demonstrate that you can use composition to seperate a class into smaller class.  
It's getting benefits in designing on software which use Single Responsibility Principle.  
More details: 
    * Reference at [here](https://edwardthienhoang.wordpress.com/2013/11/09/single-responsibility-principle-srp/)
    * Analyze at [here](https://www.meisternote.com/app/note/opdi86PCfUSp/java-radio-composition)

# Questions
Đối với ví dụ này, chúng ta cần trả lời được các câu hỏi sau:
1. Trong UML Diagram, chúng ta cần phải biết tại sao các PowerManagement, VolumneMangement, PowerStateManagement được composite vào trong Radio, và ta hiểu nó được vẽ trong UML Diagram là như thế nào?
2. Trong UML Diagram, chúng ta cũng biết rằng Radio thì được liên kết với các class PowerManagement, VolumneMangement, PowerStateManagement vì vậy Radio có quyền thêm vào hoặc không thêm vào, vậy cách thiết kế để thêm các class đó vào là như thế nào?
3. Ngoài ra, ta cũng hiểu rằng, việc thiết kế các class trên có thể được public ra ngoài hoặc đưa vào bên trong bản thiết kế của Radio. Vậy ta cần phải hiểu được cách bỏ tất cả các object composite vào bên trong là như thế nào?
# Getting Started
