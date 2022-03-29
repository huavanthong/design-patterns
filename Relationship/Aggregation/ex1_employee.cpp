#include <iostream>  

using namespace std;  

class Address {  
    // Access Specifiers
    public:  
        string addressLine, city, state;    

    // Constructor
    Address(string addressLine, string city, string state)    
    {    
        this->addressLine = addressLine;    
        this->city = city;    
        this->state = state;    
    }    
};  

class Employee    
{   
    // Access Specifiers
    private:  
        Address* address;  //Employee HAS-A Address   
    public:  
        int id;    
        string name;  

    // Constructor
    Employee(int id, string name, Address* address)    
    {    
        this->id = id;    
        this->name = name;    
        this->address = address;    
    }    

    // Define method
    void display()    
    {    
        cout<<id <<" "<<name<< " "<<     
            address->addressLine<< " "<< address->city<< " "<<address->state<<endl;    
    }    
};   

int main(void) {  
    Address a1= Address("C-146, Sec-15","Noida","UP");    
    Employee e1 = Employee(101,"Nakul",&a1);    
            e1.display();   
   return 0;  
}  