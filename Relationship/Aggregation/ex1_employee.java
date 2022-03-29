package main;

// Step 1: Define class Address
// ====> Kinh nghiệm cho thấy là, 
//       không nên sử dụng private access modifier cho các member trong class này
// Bởi vì, nếu ta muốn access từ class khác, nó được public hoặc không define
class Address {  
    String addressLine;
    String city;
    String state;

    public Address(String addressLine, String city, String state)    
    {    
        this.addressLine    = addressLine;    
        this.city           = city;    
        this.state          = state;    
    }    
}

class Employee  
{    
    private int id;  
    private String name;  
    private Address address;  

    public Employee(int id, String name, Address address)    
    {    
        this.id = id;    
        this.name = name;    
        this.address = address;    
    }    
    
    public void display()    
    {    
        System.out.println("ID: " + this.id);
        System.out.println("Name: " + this.name);
        System.out.println("AddressLine: " + this.address.addressLine); //==> Access member trong class Address
        System.out.println("City: " + this.address.city);
        System.out.println("State: " + this.address.state);
    }    
};   


public class Aggregation {
    public static void main(String []Args) {
        // Create a instance reference to object Address
        Address add = new Address("Hau Giang", "HCM", "70000");

        Employee emp = new Employee(1, "Hua Van Thong", add);
        emp.display();
    }
}  